package unittest

import (
	"chain-api-imgo/resource"
	"chain-api-imgo/services"
	"chain-api-imgo/utils"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

type CrossChain struct {
	Abi          string
	AbiBody      []byte
	Bin          string
	BinBody      []byte
	Chain        *sdk.ChainClient
	ContractName string
}

func (addr *CrossChain) Init() error {
	c, err := getChain()
	if err != nil {
		return err
	}
	addr.Chain = c
	addr.AbiBody, _ = resource.Get(addr.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	addr.BinBody, _ = resource.Get(addr.Bin) //"asserts/sol/MgtvNFTv2b1.bin"
	return nil
}

func (a *CrossChain) Run(args *Args) (err error) {
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

func (a *CrossChain) test1() {
	a.ccget("092601")
}

func (a *CrossChain) test() (err error) {
	rand.Seed(time.Now().UnixMilli())
	keyValues := []string{}
	for i := 0; i < 2; i++ {
		v := rand.Int63n(1000000)
		key := fmt.Sprintf("test_key%d", v)
		value := fmt.Sprintf("test_value_%d", v)
		err = a.ccset(key, value)
		if err != nil {
			return err
		}
		keyValues = append(keyValues, key+"="+value)
	}
	log.Println(strings.Join(keyValues, "\n"))
	return nil
}

func (a *CrossChain) createContract(contractNameHex string) (err error) {
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

func (a *CrossChain) ccset(key, value string) error {
	abiBytes, _ := resource.Get(a.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       "cc_set",
		Abi:          string(abiBytes),
		Sync:         true,
	}, key, value)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("Error Response code: %d", resp.Code)
	}
	return nil
}

func (a *CrossChain) ccget(key string) error {
	abiBytes, _ := resource.Get(a.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       "cc_get",
		Abi:          string(abiBytes),
		Sync:         true,
	}, key)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("Error Response code: %d", resp.Code)
	}

	r, e := utils.ReadOutput(string(a.AbiBody), "cc_get", resp.ContractResult.Result)
	log.Println(r, e)
	return nil
}
