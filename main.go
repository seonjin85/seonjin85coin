package main

import "github.com/seonjin85/seonjin85coin/blockchain"

func main() {
	blockchain.Blockchain().AddBlock("first")
	blockchain.Blockchain().AddBlock("second")
	blockchain.Blockchain().AddBlock("third")
}
