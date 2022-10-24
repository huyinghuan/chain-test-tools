package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ReadOutput 读取合约函数返回结果
func ReadOutput(abiRaw string, methodName string, data []byte) ([]interface{}, error) {
	myAbi, err := abi.JSON(strings.NewReader(abiRaw))
	if err != nil {
		return nil, err
	}
	method, ok := myAbi.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("合约方法不存在:%s", method)
	}
	return method.Outputs.Unpack(data)
}

func ReadOutputWithABI(myAbi abi.ABI, methodName string, data []byte) ([]interface{}, error) {
	method, ok := myAbi.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("合约方法不存在:%s", method)
	}
	return method.Outputs.Unpack(data)
}

func ReadInput(abiBytes string, data string) (methodName string, args map[string]interface{}, err error) {
	abiObj, _ := abi.JSON(strings.NewReader(abiBytes))
	decodedSig, err := hex.DecodeString(data[0:8])
	if err != nil {
		return "", nil, err
	}
	// 每个函数的签名
	// for key, v := range abiObj.Methods {
	// 	log.Println(key, hex.EncodeToString(v.ID))
	// }
	result := map[string]interface{}{}
	// 构造函数
	if hex.EncodeToString(decodedSig) == "00000000" {
		decodedData, e := hex.DecodeString(data)
		if e != nil {
			return "", nil, e
		}
		if err := abiObj.Constructor.Inputs.UnpackIntoMap(result, decodedData); err != nil {
			return "", nil, err
		} else {
			return "constructor", result, nil
		}
	}
	method, err := abiObj.MethodById(decodedSig)
	if err != nil {
		return
	}
	decodedData, err := hex.DecodeString(data[8:])
	if err != nil {
		return
	}
	err = method.Inputs.UnpackIntoMap(result, decodedData)
	return method.Name, result, err
}
