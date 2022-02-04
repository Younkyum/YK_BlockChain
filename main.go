package main

import (
	"github.com/Younkyum/nomadcoin/blockchain"
	"github.com/Younkyum/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
