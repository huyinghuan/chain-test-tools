package utils

import (
	"chain-api-imgo/resource"
	"log"
	"testing"
)

func TestAddressConver(t *testing.T) {

	userCrtBytes, _ := resource.Get("asserts/test/user.crt")
	raw, address, _, err := ParseAddrAndSkiFromCrt(string(userCrtBytes))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(raw, address)

	addr, err := ConvertAddressStringToAddress(raw)
	if err != nil {
		t.Fatal(err)
	}
	aferRaw := ConvertAddressToAddressString(addr)
	log.Println(aferRaw, addr)
	if raw != aferRaw {
		t.Fatal("address convert error")
	}
}
