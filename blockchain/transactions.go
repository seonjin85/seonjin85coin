package blockchain

import (
	"errors"
	"time"

	"github.com/seonjin85/seonjin85coin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	TXs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getID() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Onwer  string
	Amount int
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}

	txOuts := []*TxOut{
		{address, minerReward},
	}

	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	return &tx
}

func makeTx(from, to string , amount, int) (*TX, error){
  if Blockchain().BalanceByAddress(from) < amount{
	  return nil, errors.New("not enough money")
  }
}

func (m *mempool) AddTx(to string, amount int) error {
	tx , err := makeTx("seonjin85",to, amount)

	if err != nil{
		return err
	}
	m.TXs = append(m.TXs, tx)
	return nil
}
