package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Wecome to seonjin85 coin\n")
	fmt.Printf("Please use the following commands:\n")
	fmt.Printf("explorer:    Start the HTML Explorer\n")
	fmt.Printf("rest:        Start the REST API (recommended)\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}
