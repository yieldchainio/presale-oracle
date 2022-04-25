package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	config "github.com/yieldchain/presale-oracle/config"
	"github.com/yieldchain/presale-oracle/oracle"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You need to provide config file path as the only param.")
	}

	config := config.NewConfig()

	config.Load(args[0])
	fmt.Println(config)

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	//Setup an oracle per configured network
	chainChannels := make([]chan *big.Int, len(config.Networks))
	errorChannels := make([]chan bool, len(config.Networks))
	oracles := make([]*oracle.Oracle, len(config.Networks))
	for i, network := range config.Networks {
		chainChannels[i] = make(chan *big.Int)
		errorChannels[i] = make(chan bool)
		o := oracle.NewOracle(
			network.Rpc,
			network.Contract,
			big.NewInt(int64(network.ChainId)),
			network.Name, chainChannels[i],
			errorChannels[i],
			privateKey,
		)
		oracles[i] = o

		o.Connect()
		go o.Listen()
		go o.GetHardcap()
	}

	//Setup the  contribution channels
	aggregate := make(chan *big.Int)
	network := make(chan int)
	for i, ch := range chainChannels {
		go func(c chan *big.Int, net int) {
			for msg := range c {
				aggregate <- msg
				network <- net
			}
		}(ch, i)
	}

	//Setup error channels
	errorAgg := make(chan bool)
	errorNet := make(chan int)
	for i, ch := range errorChannels {
		go func(c chan bool, net int) {
			for msg := range c {
				errorAgg <- msg
				errorNet <- net
			}
		}(ch, i)
	}

	for {
		select {
		case newContrib := <-aggregate:
			total := big.NewInt(0)
			for _, oracle := range oracles {
				total = total.Add(total, oracle.Total)
			}

			x := <-network
			fmt.Println("Adding ", newContrib, " from oracle ", oracles[x].Name)
			fmt.Println("New total: ", total)
			hardCap := oracles[x].HardCap
			fmt.Println("HardCap on ", oracles[x].Name, ": ", total, "/", hardCap)
			if total.Cmp(hardCap) > 0 {
				for _, oracle := range oracles {
					go oracle.Stop()
				}
			}
		case listenerError := <-errorAgg:
			log.Println("Error handling ", listenerError)
			id := <-errorNet
			time.Sleep(1 * time.Second)
			log.Println("Listener ", oracles[id].Name, " failed, restarting")
			go oracles[id].Listen()
		}
	}

}
