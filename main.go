package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Wecome to seonjin85 coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:    Start the HTML Explorer\n")
	fmt.Printf("rest:        Start the REST API (recommended)\n")
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
