package ganache

import "github.com/ethereum/go-ethereum/ethclient"

var G *Ganache

type Ganache struct {
	Cli *ethclient.Client
}

func Init(Host string) (err error) {
	Cli, err := NewCli(Host)
	if err != nil {
		return
	}
	G = &Ganache{
		Cli: Cli,
	}
	return
}

func New() *Ganache {
	return G
}
