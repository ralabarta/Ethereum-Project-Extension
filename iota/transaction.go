package iota

import (
	"github.com/loveandpeople-DAG/goClient/api"
	"github.com/loveandpeople-DAG/goClient/bundle"
	"github.com/loveandpeople-DAG/goClient/converter"
	"github.com/loveandpeople-DAG/goClient/transaction"
	"github.com/loveandpeople-DAG/goClient/trinary"
)

type Transaction interface {
	GetMsg() string
	GetTo() string
	GetValue() uint64
}

func (L *Link) getTrytes(data string) (Trytes trinary.Trytes, err error) {
	Trytes, err = converter.ASCIIToTrytes(data)
	return
}

func (L *Link) CreatTransaction(tr Transaction, Seed trinary.Trytes) (THash *trinary.Hash, err error) {
	Message, err := L.getTrytes(tr.GetMsg())
	transfers := bundle.Transfers{
		{
			Address: tr.GetTo(),
			Value:   tr.GetValue(),
			Message: Message,
		},
	}
	trytes, err := L.API.PrepareTransfers(Seed, transfers, api.PrepareTransfersOptions{})
	if err != nil {
		return
	}

	myBundle, err := L.API.SendTrytes(trytes, depth, minimumWeightMagnitude)
	if err != nil {
		return
	}
	tailTransactionHash := bundle.TailTransactionHash(myBundle)
	THash = &tailTransactionHash
	return
}

func (L *Link) SendMessage(message Transaction, Seed trinary.Trytes) (THash *trinary.Hash, err error) {
	Message, err := L.getTrytes(message.GetMsg())
	if err != nil {
		return
	}
	transfers := bundle.Transfers{
		{
			Address: DefaultAddr,
			Value:   0,
			Message: Message,
		},
	}
	trytes, err := L.API.PrepareTransfers(Seed, transfers, api.PrepareTransfersOptions{})
	if err != nil {
		return
	}

	myBundle, err := L.API.SendTrytes(trytes, depth, minimumWeightMagnitude)
	if err != nil {
		return
	}
	tailTransactionHash := bundle.TailTransactionHash(myBundle)
	THash = &tailTransactionHash
	return
}

func (L *Link) GetMessage(THash *trinary.Hash) (jsonMsg string, err error) {
	Bundle, err := L.API.GetBundle(*THash)
	if err != nil {
		return
	}
	jsonMsg, err = transaction.ExtractJSON(Bundle)
	return
}
