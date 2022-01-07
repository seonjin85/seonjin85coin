package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash,omitempty"`
	Heigh    int    `json:"height"`
}

func createBlock(data string, prevHash string, heigh int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Heigh:    heigh,
	}

	//will be changed
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Heigh)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	return block
}
