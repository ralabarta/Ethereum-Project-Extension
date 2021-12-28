package ganache

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (G *Ganache) GetHeader(Height *big.Int) (header *types.Header, err error) {
	header, err = G.Cli.HeaderByNumber(context.Background(), Height)
	return
}

func (G *Ganache) GetBlock(Height *big.Int) (block *types.Block, err error) {
	if Height == nil {
		header, er := G.Cli.HeaderByNumber(context.Background(), Height)
		if er != nil {
			err = er
			return
		}
		Height = header.Number
	}

	block, err = G.Cli.BlockByNumber(context.Background(), Height)
	return
}
