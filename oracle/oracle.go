package oracle

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	presale "github.com/yieldchain/presale-oracle/contracts"
)

type Oracle struct {
	rpc          string
	Name         string
	address      common.Address
	client       *ethclient.Client
	contract     *presale.Presale
	contributed  *big.Int
	channel      chan *big.Int
	errorChannel chan bool
	auth         *bind.TransactOpts
	Total        *big.Int
	HardCap      *big.Int
}

func NewOracle(rpc string, contract string, chainId *big.Int, chainName string, channel chan *big.Int, errorChannel chan bool, privateKey *ecdsa.PrivateKey) *Oracle {

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(chainName, ": ", err)
	}

	contractAddress := common.HexToAddress(contract)
	oracle := Oracle{rpc: rpc, address: contractAddress, channel: channel, errorChannel: errorChannel, auth: auth, Name: chainName, Total: big.NewInt(0)}

	return &oracle
}

func (o *Oracle) Connect() {
	client, err := ethclient.Dial(o.rpc)
	if err != nil {
		log.Fatal("Error on ", o.Name, ": ", err)
	}

	o.client = client

	ctr, err := presale.NewPresale(o.address, client)
	if err != nil {
		log.Fatal("Error on ", o.Name, ": ", err)
	}
	o.contract = ctr

}

func (o *Oracle) Listen() {
	log.Println("Starting ", o.Name, "listener...")

	ctx := context.Background()

	var start uint64
	header, err := o.client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Println(o.Name, ": ", err)
	} else {
		start = header.Number.Uint64()
	}

	watchOpts := &bind.WatchOpts{Context: ctx, Start: &start}

	callOpts := &bind.CallOpts{Context: ctx}

	channel := make(chan *presale.PresaleContribution)
	chainTotal, err := o.contract.Contributed(callOpts)
	if err != nil {
		log.Println("Error on ", o.Name, ": ", err)
		o.errorChannel <- true
		return
	}

	o.Total = o.Total.Add(o.Total, chainTotal)
	o.channel <- chainTotal

	sub, err := o.contract.WatchContribution(watchOpts, channel, nil, nil)
	if err != nil {
		log.Println("Error on ", o.Name, ": ", err)
		o.errorChannel <- true
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-channel:
			fmt.Println("Account ", event.Buyer, " contributed ", event.Amount, " in ", event.Token)
			o.Total = o.Total.Add(o.Total, event.Amount)
			o.channel <- event.Amount
		case err := <-sub.Err():
			log.Printf("Error in %s: %v", o.Name, err)
			o.errorChannel <- (err != nil)
			log.Println("Error sent")
			o.Total.SetUint64(0)
			return
		}
	}
}

func (o *Oracle) Stop() {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		nonce, err := o.client.PendingNonceAt(context.Background(), o.auth.From)
		if err != nil {
			log.Println(err)
			continue
		}

		o.auth.Nonce = big.NewInt(int64(nonce))

		gasPrice, err := o.client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		o.auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(i+2)))

		callOpts := &bind.CallOpts{Context: ctx}
		isOpen, err := o.contract.IsOpen(callOpts)
		if err != nil {
			log.Println(err)
			continue
		}

		if isOpen {
			tx, err := o.contract.SetSaleOpen(o.auth, false)
			if err != nil {
				log.Printf("Failed to stop presale at %s (%s), nonce %d: %v", o.Name, o.address.String(), nonce, err)
				log.Println("Retry to stop ", o.Name)
				continue
			}
			log.Println("Stopping contributions in TX ", tx.Hash().Hex())
			success := false
			time.Sleep(10 * time.Second)
			for j := 0; j < 6; j++ {
				receipt, err := o.client.TransactionReceipt(ctx, common.BytesToHash([]byte("0x8dbca7f86c08278f0a7702114a3fc4e90830be0576f05950dc11a2af10457933"))) //tx.Hash())
				if err != nil {
					log.Println("TX ", tx.Hash(), ": ", err)
					continue
				} else if receipt.Status == 1 {
					success = true
					break
				}
				time.Sleep(5 * time.Second)
			}
			if success {
				isOpen, err := o.contract.IsOpen(callOpts)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("Status of ", o.Name, ": ", isOpen)
				break
			}
		} else {
			log.Printf("Already closed on %s (%s)", o.Name, o.address)
			break
		}

	}
}

func (o *Oracle) GetHardcap() {
	ctx := context.Background()

	callOpts := &bind.CallOpts{Context: ctx}
	hardCap, err := o.contract.HARDCAP(callOpts)
	if err != nil {
		log.Println("Error on ", o.Name, ": ", err)
		o.errorChannel <- true
		return
	}

	o.HardCap = hardCap

}
