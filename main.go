package main

import (
	"fmt"

	"github.com/seonjin85/seonjin85coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hata : %s\n", block.Hash)
		fmt.Printf("Prev Hata : %s\n", block.PrevHash)
	}
}
