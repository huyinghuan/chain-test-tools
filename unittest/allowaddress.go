package unittest

import (
	"bufio"
	"chain-api-imgo/resource"
	"chain-api-imgo/services"
	"chain-api-imgo/utils"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type AllowAddress struct {
	Abi          string
	AbiBody      []byte
	Bin          string
	BinBody      []byte
	AddressFile  string
	StartLine    int
	EndLine      int
	Chain        *sdk.ChainClient
	ContractName string
}

func (addr *AllowAddress) Init() error {
	c, err := getChain()
	if err != nil {
		return err
	}
	addr.Chain = c
	addr.AbiBody, _ = resource.Get(addr.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	addr.BinBody, _ = resource.Get(addr.Bin) //"asserts/sol/MgtvNFTv2b1.bin"
	return nil
}

func (a *AllowAddress) Run(args *Args) (err error) {
	//addressFile := ""
	contractName := a.ContractName
	if a.ContractName == "" {
		contractName, err = a.createContract()
		a.ContractName = contractName
	}
	if err != nil {
		return err
	}
	log.Println(a.ContractName)
	return a.fillAddress()
	//return a.getAllowAddressCount()
	//return a.ensureAddressInAllowList("0x8C41B466D082e67E1fB8ea74178D373678a8a314")
}

func (a *AllowAddress) createContract() (contractName string, err error) {
	txSuccessCount := 0
	txFailCount := 0
	defer func() {
		log.Printf("完成交易 sucess:%d fail:%d total:%d \n", txSuccessCount, txFailCount, txSuccessCount+txFailCount)
	}()

	planPublishCount := new(big.Int).SetUint64(0)
	rand.Seed(time.Now().Unix())
	contractName = "profile_" + utils.RandomStr(24)
	contractNameHex := services.CalcContractName(contractName)
	ressponse, err := services.NewContract(a.Chain, a.AbiBody, a.BinBody, contractNameHex, contractName, "ZP", planPublishCount)
	if err != nil {
		return
	}
	if ressponse.Code != 0 {
		err = fmt.Errorf("Error Response code: %d", ressponse.Code)
		return
	}
	return
}

func (a *AllowAddress) fillAddress() error {
	f, err := os.OpenFile(a.AddressFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	index := 0
	for sc.Scan() {
		if index < a.StartLine {
			index++
			continue
		}
		if index >= a.EndLine {
			break
		}
		index++
		a.addAddress(sc.Text())
	}
	return nil
}

func (a AllowAddress) addAddress(address string) error {

	abiBytes, _ := resource.Get(a.Abi) //"asserts/sol/MgtvNFTv2b1.abi"
	addr, _ := utils.ConvertAddressStringToAddress(address)
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       "addAllowAddress",
		Abi:          string(abiBytes),
		Sync:         false,
	}, addr)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("Error Response code: %d", resp.Code)
	}
	return nil
}

func (a AllowAddress) getAllowAddressCount() error {
	abiBytes, _ := resource.Get(a.Abi)
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       "getAllowAddressCount",
		Abi:          string(abiBytes),
		Sync:         true,
	})
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
	count, err := utils.ReadOutputWithABI(myAbi, "getAllowAddressCount", resp.ContractResult.Result)
	log.Println(count)
	return err
}

func (a *AllowAddress) ensureAddressInAllowList(beCheckedAddressRaw string) error {

	addr, _ := utils.ConvertAddressStringToAddress(beCheckedAddressRaw)
	start := time.Now()
	resp, err := services.InvokeWithExtParams(a.Chain, services.InvokeParams{
		ContractName: a.ContractName,
		Method:       "isAllowAddress",
		Abi:          string(a.AbiBody),
		Sync:         true,
	}, addr)
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
	count, err := utils.ReadOutputWithABI(myAbi, "isAllowAddress", resp.ContractResult.Result)
	log.Println(count, time.Since(start))
	return err
}
