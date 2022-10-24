package unittest

import (
	"chain-api-imgo/resource"
	"chain-api-imgo/services"

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
	var keyBody, crtBody []byte
	if keyBody, err = resource.Get("asserts/localtest/crypto-config/mangochain1.mgtv.com/user/client1/client1.sign.key"); err != nil {
		log.Println(err)
		return
	}
	if crtBody, err = resource.Get("asserts/localtest/crypto-config/mangochain1.mgtv.com/user/client1/client1.sign.crt"); err != nil {
		return
	}
	chainClient, _, err = services.GetKeepAliveChainClient(true, keyBody, crtBody)
	return
}

func simpleShow(fn func()) {
	log.Println("\n\n\n----------------")
	fn()
	log.Println("\n\n\n----------------")
}
