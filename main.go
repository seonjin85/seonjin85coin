package main

import (
	"github.com/seonjin85/seonjin85coin/blockchain"
	"github.com/seonjin85/seonjin85coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
