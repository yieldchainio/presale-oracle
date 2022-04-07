package oracle

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	presale "github.com/yieldchain/presale-oracle/contracts"
)

type Oracle struct {
	rpc         string
	address     common.Address
	client      *ethclient.Client
	contract    *presale.Presale
	contributed *big.Int
	channel     chan *big.Int
	auth        *bind.TransactOpts
}

func NewOracle(rpc string, contract string, chainId *big.Int, channel chan *big.Int, privateKey *ecdsa.PrivateKey) *Oracle {

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contract)
	oracle := Oracle{rpc: rpc, address: contractAddress, channel: channel, auth: auth}

	return &oracle
}

func (o *Oracle) Connect() {
	log.Println(o.rpc)
	client, err := ethclient.Dial(o.rpc)
	if err != nil {
		log.Fatal(err)
	}

	o.client = client

	ctr, err := presale.NewPresale(o.address, client)
	if err != nil {
		log.Fatal(err)
	}
	o.contract = ctr

}

func (o *Oracle) Listen() {
	ctx := context.Background()
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	callOpts := &bind.CallOpts{Context: ctx}

	channel := make(chan *presale.PresaleContribution)
	chainTotal, err := o.contract.Contributed(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	o.channel <- chainTotal

	sub, err := o.contract.WatchContribution(watchOpts, channel, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-channel:
			fmt.Println("Account ", event.Buyer, " contributed ", event.Amount, " in ", event.Token)
			o.channel <- event.Amount
		case err := <-sub.Err():
			log.Fatal(err)
		}
	}
}

func (o *Oracle) Stop() {
	ctx := context.Background()
	chainId, err := o.client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := o.client.PendingNonceAt(context.Background(), o.auth.From)
	if err != nil {
		log.Fatal(err)
	}

	o.auth.Nonce = big.NewInt(int64(nonce))

	/*gasPrice, err := o.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	o.auth.GasPrice = gasPrice*/

	callOpts := &bind.CallOpts{Context: ctx}
	isOpen, err := o.contract.IsOpen(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	if isOpen {
		tx, err := o.contract.SetSaleOpen(o.auth, false)
		if err != nil {
			log.Fatalf("Failed to stop presale at %d (%s): %v", chainId.Int64(), o.address.String(), err)
		}
		log.Println("Stopping contributions in TX ", tx.Hash())
	} else {
		log.Printf("Already closed on %d (%s)", chainId.Int64(), o.address)
	}
}
