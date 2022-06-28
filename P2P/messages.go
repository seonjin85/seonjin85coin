package p2p

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

