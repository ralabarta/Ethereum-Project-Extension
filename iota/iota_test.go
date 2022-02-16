package iota

import (
	"encoding/json"
	"fmt"
	"github.com/loveandpeople-DAG/goClient/api"
	"github.com/loveandpeople-DAG/goClient/trinary"
	"testing"
)

const (
	TestLink     = "http://27.102.115.79:14265"
	User         = ""
	Pwd          = ""
	CreationSeed = ""
	CAddr        = ""
	TestSeed     = ""
	TestAddr     = ""
)

func InitTest() {
	_, err := InitLink(TestLink, User, Pwd)
	if err != nil {
		fmt.Println("InitLink err ", err.Error())
	}
}

/* run TestLink_GetNewAddress */
func TestLink_GetNewAddress(t *testing.T) {
	InitTest()
	add, err := Lin.GetNewAddress(TestSeed, 0)
	if err != nil {
		fmt.Println("GetNewAddress err ", err.Error())
	} else {
		for k, v := range add {
			fmt.Println(k, " => ", v)
		}
	}
}

func TestIotaGetNodeInfo(t *testing.T) {
	InitTest()
	info, err := Lin.API.GetNodeInfo()
	if err != nil {
		fmt.Println("GetNodeInfo err ", err.Error())
	} else {
		fmt.Printf("INFO IS \n %+v \n", info)
		//fmt.Println("Info Time", info.Time)
		//fmt.Println("Info NodeAlias", info.NodeAlias)
		//fmt.Println("Info AutopeeringID", info.AutopeeringID)
		//fmt.Println("Info Uptime", info.Uptime)
		//fmt.Println("Info AppName", info.AppName)
		//fmt.Println("Info IsHealthy", info.IsHealthy)
		//fmt.Println("Info Duration", info.Duration)
		//fmt.Println("Info TransactionsToRequest", info.TransactionsToRequest)
		//fmt.Println("Info Tips", info.Tips)
		//fmt.Println("Info Neighbors", info.Neighbors)
		//fmt.Println("Info LatestMilestone", info.LatestMilestone)
		//fmt.Println("Info LatestMilestoneIndex", info.LatestMilestoneIndex)
		//fmt.Println("Info LatestSolidSubtangleMilestone", info.LatestSolidSubtangleMilestone)
		//fmt.Println("Info LatestSolidSubtangleMilestoneIndex", info.LatestSolidSubtangleMilestoneIndex)
	}
}

func TestIotaGetNeighbors(t *testing.T) {
	InitTest()
	Neighbors, err := Lin.API.GetNeighbors()
	if err != nil {
		fmt.Println("GetNeighbors err ", err.Error())
	} else {
		fmt.Printf("Neighbors %+v \n", Neighbors)
	}
}

func TestIotaGetBalances(t *testing.T) {
	InitTest()
	TestIotaGetInput(t)
	addrs := []string{
		"OLCHRMEFITACDJWFTSIQWPOUWZIPCTEFUPEXTMDELUCAXENBDWIINONPFXHRTIUMVAU99VNERAVVQYYXCWXPXRGAYD",
	}
	Balances, err := Lin.API.GetBalances(addrs)
	if err != nil {
		fmt.Println("GetBalances err ", err.Error())
	} else {
		fmt.Printf("Balances %+v \n", Balances)
	}
}

func TestIotaGetTransactions(t *testing.T) {
	InitTest()
	Transfers, err := Lin.API.GetTransfers(CreationSeed, api.GetTransfersOptions{Start: 3})
	if err != nil {
		fmt.Println("GetBalances err ", err.Error())
	} else {
		fmt.Println("GetBalances success ")
		fmt.Printf("%+v \n", Transfers)
		for k, v := range Transfers {
			fmt.Println("No.", k)
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			for vk, vv := range v {
				fmt.Println("vk No.", vk)
				fmt.Println("Bundle", vv.Bundle)
				fmt.Println("Address", vv.Address)
				fmt.Println("Value", vv.Value)
				fmt.Println("Hash", vv.Hash)
				fmt.Println("Confirmed", vv.Confirmed)
				fmt.Println("---------------------------------")
			}
		}
	}

	//Transactions, err := Lin.API.FindTransactionObjects(api.FindTransactionsQuery{
	//	Addresses: []trinary.Hash{
	//		"OLCHRMEFITACDJWFTSIQWPOUWZIPCTEFUPEXTMDELUCAXENBDWIINONPFXHRTIUMVAU99VNERAVVQYYXC",
	//	},
	//})
	//if err != nil {
	//	fmt.Println("GetBalances err ", err.Error())
	//} else {
	//	for k, v := range Transactions{
	//		fmt.Println("No.", k)
	//		fmt.Println("Address", v.Address)
	//		fmt.Println("Value", v.Value)
	//		fmt.Println("Hash", v.Hash)
	//		fmt.Println("Confirmed", v.Confirmed)
	//		fmt.Println("---------------------------------")
	//	}
	//}

	//TxHash := "VTCPXLFHNPABBYHTPICNFKDHHYJZHYGNCXUAJIFRCILHLZQJGSZTKXYPDSFGODQVA9LJPNVZHJFFIFRTC"
	//TraverseBundle, err := Lin.API.TraverseBundle(TxHash, bundle.Bundle{})
	//if err != nil {
	//	fmt.Println("GetBalances err ", err.Error())
	//} else {
	//	for k, v := range TraverseBundle{
	//		fmt.Println("No.", k)
	//		fmt.Println("Address", v.Address)
	//		fmt.Println("Value", v.Value)
	//		fmt.Println("Hash", v.Hash)
	//		fmt.Println("Confirmed", v.Confirmed)
	//		fmt.Println("---------------------------------")
	//	}
	//}

	//Hash := "DRQJXFIRMKJFQUUMHLETZMKKVTSJZRQRGSOBTHLAOKQBRVQQMPDAIATDIAGZQVGHHPJKYFASUVVQFX999"
	//Tx, err := Lin.API.GetTrytes(Hash)
	//if err != nil {
	//	fmt.Println("GetBalances err ", err.Error())
	//} else {
	//	for k, v := range Tx{
	//		fmt.Println("No.", k,"Tx", v)
	//	}
	//}

	//Addr := "VTCPXLFHNPABBYHTPICNFKDHHYJZHYGNCXUAJIFRCILHLZQJGSZTKXYPDSFGODQVA9LJPNVZHJFFIFRTC"
	//Tx, err := Lin.API.FindTransactions(api.FindTransactionsQuery{
	//	Addresses: trinary.Hashes{Addr},
	//})
	//if err != nil {
	//	fmt.Println("FindTransactions err ", err.Error())
	//} else {
	//	for k, v := range Tx{
	//		fmt.Println("No.", k,"Tx", v)
	//	}
	//}

}

func TestIotaGetInput(t *testing.T) {
	InitTest()
	opts := api.PrepareTransfersOptions{}
	var totalOutput uint64
	totalOutput = 1000
	Inputs, err := Lin.API.GetInputs(CreationSeed, api.GetInputsOptions{Security: opts.Security, Threshold: &totalOutput})
	if err != nil {
		fmt.Println("GetInputs err ", err.Error())
	} else {
		for k, v := range Inputs.Inputs {
			fmt.Println("No.", k)
			fmt.Println("Addr", v.Address)
			fmt.Println("KeyIndex", v.KeyIndex)
			fmt.Println("Security", v.Security)
			fmt.Println("Balance", v.Balance)
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		}
	}

}

type TestTr struct {
	To     trinary.Trytes
	Amount uint64
	Msg    TestMsg
}

type TestMsg struct {
	Msg string `json:"msg"`
}

func (T *TestTr) GetMsg() string {
	str, _ := json.Marshal(T.Msg)
	return string(str)
}

func (T *TestTr) GetTo() string {
	return T.To
}

func (T *TestTr) GetValue() uint64 {
	return T.Amount
}

func TestLink_Transaction(t *testing.T) {
	InitTest()
	Tr := &TestTr{
		To:     TestAddr,
		Amount: 1000000,
		Msg: TestMsg{
			Msg: "Test TX",
		},
	}

	THash, err := Lin.CreatTransaction(Tr, CreationSeed)
	if err != nil {
		panic("SendMessage " + err.Error())
	} else {
		fmt.Println("TxHash :", *THash)
	}

	//tmpHash := "DRQJXFIRMKJFQUUMHLETZMKKVTSJZRQRGSOBTHLAOKQBRVQQMPDAIATDIAGZQVGHHPJKYFASUVVQFX999"
	//THash := &tmpHash
	//TestTr, err := Lin.API.GetBundle(*THash)
	//if err != nil {
	//	panic("GetMessage " + err.Error())
	//}
	//for k, v := range TestTr {
	//	fmt.Println("k", k)
	//	fmt.Println("Tr Hash", v.Hash)
	//	fmt.Println("Tr Value", v.Value)
	//	fmt.Println("Tr Confirmed", v.Confirmed)
	//}

	//if Message != NMsg {
	//	panic("Message comparison failed!")
	//} else {
	//	fmt.Println("Message comparison success!")
	//}
}
