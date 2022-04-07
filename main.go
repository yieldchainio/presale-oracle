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

	chainChannels := make([]chan *big.Int, len(config.Networks))
	for i, _ := range config.Networks {
		chainChannels[i] = make(chan *big.Int)
	}
	total := big.NewInt(0)

	for i, network := range config.Networks {
		go func() {
			contractAddress := common.HexToAddress(network.Contract)
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

			chainChannels[i] <- chainTotal

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
					chainChannels[i] <- event.Amount
				case err := <-sub.Err():
					log.Fatal(err)
				}
			}

		}()
	}

	aggregate := make(chan *big.Int)
	for _, ch := range chainChannels {
		go func(c chan *big.Int) {
			for msg := range c {
				aggregate <- msg
			}
		}(ch)
	}

	for {
		select {
		case newContrib := <-aggregate:
			total = total.Add(total, newContrib)
			fmt.Println("Total: ", total)
		}
	}

}
