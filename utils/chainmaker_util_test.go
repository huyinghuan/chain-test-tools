package utils

import (
	"chain-api-imgo/resource"
	"log"
	"math/big"
	"testing"

	"chainmaker.org/chainmaker/common/v2/evmutils"
)

func TestParseAddrAndSkiFromCrt(t *testing.T) {
	userCrtBytes, _ := resource.Get("asserts/test/mgmgtv.crt")
	raw, address, ski, err := ParseAddrAndSkiFromCrt(string(userCrtBytes))
	if err != nil {
		t.Fatal(err)
	}

	log.Println(raw, address, ski)
	n := new(big.Int)
	n.SetString(raw[2:], 16)
	log.Println(n)
	afterConver := evmutils.FromBigInt(n)
	log.Println(evmutils.BigToAddress(afterConver))
	//log.Println(v ...interface{})
	// addressInt := new(evmutils.Int)
	// addressInt.Set(n)
	// // addressInt.SetBytes([]byte(raw[2:]))
	// log.Println(evmutils.BigToAddress(addressInt))
}
