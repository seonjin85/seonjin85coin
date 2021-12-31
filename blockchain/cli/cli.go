package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/seonjin85/seonjin85coin/explorer"
	"github.com/seonjin85/seonjin85coin/rest"
)

func usage() {
	fmt.Printf("Wecome to seonjin85 coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000 :    Start the HTML Explorer\n")
	fmt.Printf("-mode=rest :    Choose between 'html' and 'rest'\n")
	os.Exit(0)
}
func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "choose between 'html' and 'rest'")

	flag.Parse()
	switch *mode {
	case "rest":
		//start rest api
		rest.Start(*port)
	case "html":
		//start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
}
