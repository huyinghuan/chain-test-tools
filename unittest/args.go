package unittest

import (
	"chain-api-imgo/config"
	"chain-api-imgo/resource"
	"chain-api-imgo/services"
	"fmt"

	"log"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

type Args struct {
	Env  string
	N    int
	Tx   int
	Loop int
	S    int
}

func (a *Args) times() int {
	return a.N * a.Loop * a.Tx
}

func getChain() (chainClient *sdk.ChainClient, err error) {
	var keyBody, crtBody, signKeyBody, signCrtBody []byte
	conf := config.GetConfig()
	if keyBody, err = resource.Get(conf.Client.Key); err != nil {
		log.Println(err)
		return
	}
	if crtBody, err = resource.Get(conf.Client.Crt); err != nil {
		return
	}
	if signKeyBody, err = resource.Get(conf.Client.SignKey); err != nil {
		log.Println(err)
		return
	}
	if signCrtBody, err = resource.Get(conf.Client.SignCrt); err != nil {
		return
	}
	chainClient, _, err = services.GetKeepAliveChainClient(true, keyBody, crtBody, signKeyBody, signCrtBody)
	return
}

func simpleShow(fn func()) {
	fmt.Println("\n\n\n----------------")
	fn()
	fmt.Println("----------------")
}
