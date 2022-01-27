package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to YK Coin\n\n")
	fmt.Printf("Pleas use the following commads:\n\n")
	fmt.Printf("exploere: Start the HTML Explorer\n")
	fmt.Printf("rest:     Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
}
