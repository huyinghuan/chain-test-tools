package utils

import (
	"fmt"
	"math/big"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils"
)

func ConvertAddressStringToAddress(addressRaw string) (evmutils.Address, error) {
	if !strings.HasPrefix(addressRaw, "0x") {
		return evmutils.Address{}, fmt.Errorf("地址格式错误:%x", addressRaw)
	}
	n := new(big.Int)
	n.SetString(addressRaw[2:], 16)
	afterConver := evmutils.FromBigInt(n)
	address := evmutils.BigToAddress(afterConver)
	return address, nil
}
func ConvertAddressToAddressString(orginAddress evmutils.Address) string {
	n := new(big.Int)
	n.SetBytes(orginAddress[:])
	afterConver := evmutils.FromBigInt(n)
	return fmt.Sprintf("0x%x", afterConver.AsStringKey())
}

func ConvertAddressBytesToAddressString(orginAddress []byte) string {
	n := new(big.Int)
	n.SetBytes(orginAddress)
	afterConver := evmutils.FromBigInt(n)
	return fmt.Sprintf("0x%x", afterConver.AsStringKey())
}
