package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraUrl = "https://mainnet.infura.io/v3/3e1ba7b154e840e585a87c4e1822918b"
var ganacheURl = "http://127.0.0.1:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURl)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xa90a89b5917d8907Fe962943D6Dca29b243469e7")
	bal, _ := client.BalanceAt(context.Background(), address, nil)
	fmt.Println("Balance: ", bal)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number())
}
