package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	config "github.com/yieldchain/presale-oracle/config"
	presale "github.com/yieldchain/presale-oracle/contracts"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You need to provide config file path as the only param.")
	}

	config := config.NewConfig()

	config.Load(args[0])
	fmt.Println(config)

	fmt.Println("Testing")
	client, err := ethclient.Dial(config.Networks[0].Rpc)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	contributed := make(chan *big.Int)
	total := big.NewInt(0)

	go func() {
		contractAddress := common.HexToAddress(config.Networks[0].Contract)
		contract, err := presale.NewPresale(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		var start uint64 = 18209890
		watchOpts := &bind.WatchOpts{Context: ctx, Start: &start}

		callOpts := &bind.CallOpts{Context: ctx}

		channel := make(chan *presale.PresaleContribution)
		chainTotal, err := contract.Contributed(callOpts)
		if err != nil {
			log.Fatal(err)
		}

		contributed <- chainTotal

		sub, err := contract.WatchContribution(watchOpts, channel, nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Subscribed")
		defer sub.Unsubscribe()

		for {
			select {
			case event := <-channel:
				fmt.Println("Account ", event.Buyer, " contributed ", event.Amount, " in ", event.Token)
				contributed <- event.Amount
			case err := <-sub.Err():
				log.Fatal(err)
			}
		}

	}()

	for {
		select {
		case newContrib := <-contributed:
			total = total.Add(total, newContrib)
			fmt.Println(total)
		}
	}

}
