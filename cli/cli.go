package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	explorer "github.com/Younkyum/nomadcoin/explorer/templates"
	"github.com/Younkyum/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to YK Coin\n\n")
	fmt.Printf("Pleas use the following flags:\n\n")
	fmt.Printf("-port:   Set the PORT of the server\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
