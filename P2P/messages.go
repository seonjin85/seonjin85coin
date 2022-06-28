package p2p

import (
	"encoding/json"

	"github.com/seonjin85/seonjin85coin/blockchain"
	"github.com/seonjin85/seonjin85coin/utils"
)

type MessageKind int

const (
	MessageNewstBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func (m *Message) addPayload(p interface{}) {
	//  marshal : json -> v , unmarshal : v -> json
	b, err := json.Marshal(p)
	utils.HandleErr(err)
	m.Payload = b
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}
	m.addPayload(payload)
	mJson, err := json.Marshal(m)
	utils.HandleErr(err)
	return mJson
}

func sendNewstBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewstBlock, b)
	p.inbox <- m
}
