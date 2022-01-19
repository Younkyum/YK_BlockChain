package main

import (
	explorer "github.com/Younkyum/nomadcoin/explorer/templates"
	"github.com/Younkyum/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
