package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Networks   []Network `json:"networks"`
	PrivateKey string    `json:"privateKey"`
}

type Network struct {
	Rpc      string `json:"rpc"`
	Contract string `json:"contract"`
	ChainId  uint   `json:"chainId"`
	Name     string `json:"name"`
}

func NewConfig() *Config {
	config := Config{}

	return &config
}

func (c *Config) Load(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, c)
}

func (c *Config) String() string {
	result := ""
	result += fmt.Sprintf("Networks:\n")
	for _, network := range c.Networks {
		result += fmt.Sprintf("Chain %d\n", network.ChainId)
		result += fmt.Sprintf("\tRpc: %s\n", network.Rpc)
		result += fmt.Sprintf("\tContract: %s\n", network.Contract)
		result += fmt.Sprintf("---------------------- \n")
	}

	return result
}
