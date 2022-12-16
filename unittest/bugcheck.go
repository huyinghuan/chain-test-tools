package unittest

import (
	"chain-api-imgo/resource"
	"chain-api-imgo/services"
	"chain-api-imgo/utils"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type BugCheckt struct {
	Abi          string
	AbiBody      []byte
	Bin          string
	BinBody      []byte
	Chain        *sdk.ChainClient
	ContractName string
}

func (addr *BugCheckt) Init() error {
	c, err := getChain()
	if err != nil {
		return err
	}
	addr.Chain = c
	addr.AbiBody, _ = resource.Get(addr.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	addr.BinBody, _ = resource.Get(addr.Bin) //"asserts/sol/MgtvNFTv2b1.bin"
	return nil
}

func (a *BugCheckt) Run(args *Args) (err error) {
	// return a.doTx(args)
	//return a.doQueryTx("9a18030f1fb346b99f781d252936def4450b54a5c8f4406fbca3d9bfc4745635") // old
	return a.doQueryTx("ce27b7a5a8b74dbfb5d1022411352c953bd0b7b6915b4980b91e6fadabfc900b") // after upgrade
}

func (a *BugCheckt) doQueryTx(txId string) error {
	chain, err := getChain()
	if err != nil {
		return err
	}
	tx, err := chain.GetTxByTxId(txId)
	if err != nil {
		return err
	}
	log.Println(tx)
	return nil
}

// 执行交易
func (a *BugCheckt) doTx(args *Args) (err error) {
	contractName := a.ContractName
	contractNameHex := ""
	if a.ContractName == "" {
		contractName = "profile_" + utils.RandomStr(24)
		contractNameHex = services.CalcContractName(contractName)
		err = a.createContract(contractNameHex)
		a.ContractName = contractName
	}
	if err != nil {
		return err
	}
	log.Println(a.ContractName, contractNameHex)
	a.test1()
	return
}

func (a *BugCheckt) test1() {
	// a.int2str(92601)
	// a.int2str(92601)
	//a.call("baseURI")
	a.call("charCodeAt", "0")
	a.call("int2str", new(big.Int).SetUint64(22))
	a.call("tokenURI", new(big.Int).SetUint64(22))
}

func (a *BugCheckt) createContract(contractNameHex string) (err error) {
	txSuccessCount := 0
	txFailCount := 0
	defer func() {
		log.Printf("完成交易 sucess:%d fail:%d total:%d \n", txSuccessCount, txFailCount, txSuccessCount+txFailCount)
	}()
	rand.Seed(time.Now().Unix())
	ressponse, err := services.NewContract(a.Chain, a.AbiBody, a.BinBody, contractNameHex)
	if err != nil {
		return
	}
	if ressponse.Code != 0 {
		err = fmt.Errorf("Error Response code: %d", ressponse.Code)
		return
	}
	return
}

func (a *BugCheckt) call(method string, args ...interface{}) error {
	abiBytes, _ := resource.Get(a.Abi)
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       method,
		Abi:          string(abiBytes),
		Sync:         true,
	}, args...)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("Error Response code: %d", resp.Code)
	}
	myAbi, err := abi.JSON(strings.NewReader(string(a.AbiBody)))
	if err != nil {
		return err
	}
	simpleShow(func() {
		result, _ := utils.ReadOutputWithABI(myAbi, method, resp.ContractResult.Result)
		fmt.Println("metod", result, resp.TxId)
		// log.Println(resp.ContractResult.Result)
	})
	return nil

}
