package iota

import (
	"github.com/loveandpeople-DAG/goClient/api"
	"github.com/loveandpeople-DAG/goClient/trinary"
)

const securityLevel = 2

func (L *Link) GetNewAddress(Seed string, index uint64) (addresses trinary.Hashes, err error) {
	addresses, err = L.API.GetNewAddress(Seed, api.GetNewAddressOptions{
		Security:  securityLevel,
		Index:     index,
		ReturnAll: true,
	})
	return
}
