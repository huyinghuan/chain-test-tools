package unittest

import (
	"chain-api-imgo/resource"
	"chain-api-imgo/services"
	"chain-api-imgo/utils"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

func ProfileTest(args *Args) {
	var wg sync.WaitGroup
	wg.Add(args.N)
	needWait := args.times() > 1000
	for i := 0; i < args.N; i++ {
		go func() {
			for j := 0; j < args.Loop; j++ {
				SendCreateContract(args.Tx)
			}
			if needWait && args.S > 0 {
				time.Sleep(time.Duration(args.S) * time.Second)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func SendCreateContract(tx int) {
	txSuccessCount := 0
	txFailCount := 0
	defer func() {
		log.Printf("完成交易 sucess:%d fail:%d total:%d \n", txSuccessCount, txFailCount, txSuccessCount+txFailCount)
	}()
	var err error
	var crtBody []byte
	if crtBody, err = resource.Get("asserts/localtest/crypto-config/mangochain1.mgtv.com/user/client1/client1.sign.crt"); err != nil {
		log.Println(err)
		return
	}
	chainClient, err := getChain()
	abiBytes, _ := resource.Get("asserts/sol/mgtv.abi")
	binBytes, _ := resource.Get("asserts/sol/mgtv.bin")
	planPublishCount := new(big.Int).SetUint64(0)

	rand.Seed(time.Now().Unix())
	contractName := "profile_" + utils.RandomStr(24)
	contractNameHex := services.CalcContractName(contractName)
	ressponse, err := services.NewContract(chainClient, abiBytes, binBytes, contractNameHex, contractName, "ZP", planPublishCount)
	if err != nil {
		log.Println(err)
		return
	}
	if ressponse.Code != 0 {
		log.Println(ressponse.Message)
		return
	}
	for i := 1; i <= tx; i++ {
		_, addr, _, _ := utils.ParseAddrAndSkiFromCrtBytes(crtBody)
		token := new(big.Int).SetUint64(uint64(i))
		//resp, err := Invoke(client, contract.TokenAbi, contract.TokenName, "mint", address, token)
		resp, err := services.InvokeWithExtParams(chainClient, services.InvokeParams{
			ContractName: contractName,
			Method:       "mint",
			Abi:          string(abiBytes),
			Sync:         false,
		}, addr, token)
		if err != nil {
			log.Println(err)
			return
		}
		if resp.Code != 0 {
			txFailCount++
			log.Println(resp.Message)
			return
		}
		txSuccessCount++
	}
}
