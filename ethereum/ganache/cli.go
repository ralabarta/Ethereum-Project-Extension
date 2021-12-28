package ganache

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewCli(Host string) (Cli *ethclient.Client, err error) {
	Cli, err = ethclient.Dial(Host)
	return
}
