package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/seonjin85/seonjin85coin/blockchain"
	"github.com/seonjin85/seonjin85coin/utils"
)

type MessageKind int

const (
	MessageNewstBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
	MessageNewBlockNotify
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewstBlock(p *peer) {
	fmt.Printf("sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewstBlock, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func notifyNewBlock(b *blockchain.Block, p *peer) {
	m := makeMessage(MessageNewBlockNotify, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewstBlock:
		fmt.Printf("Received the newest block from %s\n", p.key)
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)

		if payload.Height >= b.Height {
			//request all the blocks from 4000
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requestAllBlocks(p)
		} else {
			//send 4000 our block
			sendNewstBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks. \n", p.key)
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Received all the blocks from %s\n", p.key)
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().Replace(payload)
	case MessageNewBlockNotify:
		var payload *blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().AddPeerBlock(payload)
	}
}
