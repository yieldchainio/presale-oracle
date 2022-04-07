package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

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

	chainChannels := make([]chan *big.Int, len(config.Networks))
	for i, _ := range config.Networks {
		chainChannels[i] = make(chan *big.Int)
	}
	total := big.NewInt(0)

	oracles := make([]*oracle.Oracle, len(config.Networks))

	for i, network := range config.Networks {
		o := oracle.NewOracle(network.Rpc, network.Contract, big.NewInt(int64(network.ChainId)), chainChannels[i], privateKey)
		oracles[i] = o

		o.Connect()
		go o.Listen()
	}

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

	for {
		select {
		case newContrib := <-aggregate:
			total = total.Add(total, newContrib)
			x := <-network
			fmt.Println("Adding ", newContrib, " from oracle ", x)
			fmt.Println("Total: ", total)
			y := big.NewInt(760000000000)
			yMul := big.NewInt(10000000000)
			y = y.Mul(y, yMul)
			if total.Cmp(y) > 0 {
				for _, oracle := range oracles {
					oracle.Stop()
				}
			}

		}
	}

}
