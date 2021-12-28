package ganache

import (
	"fmt"
	"math/big"
	"testing"
)

var Host = "HTTP://127.0.0.1:7545"
var Gan *Ganache

func InitG() {
	err := Init(Host)
	if err != nil {

	}
	Gan = New()
}

func TestGanache_GetBalance(t *testing.T) {
	InitG()
	balance, err := Gan.GetBalance("0x2A3eBaD58f30501607171271421CaCb2C3aCB037", nil)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	fmt.Println("balance", balance)
}

func TestGanache_NewAccount(t *testing.T) {
	InitG()
	Gan.NewAccount()
}

func TestGanache_GetBlock(t *testing.T) {
	InitG()
	blockNumber := big.NewInt(0)
	block, err := Gan.GetBlock(blockNumber)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	//fmt.Printf("header : %+v \n", header)
	if block != nil {
		fmt.Printf("block %+v \n", block)
	}
}
