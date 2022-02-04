package main

import (
	"github.com/Younkyum/nomadcoin/cli"
	"github.com/Younkyum/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
