//go:build linux
// +build linux

/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sym

import (
	"encoding/hex"
	"errors"
	"fmt"

	"chainmaker.org/chainmaker/common/v2/crypto/engine"
	"chainmaker.org/chainmaker/common/v2/opencrypto"
	gsm4 "chainmaker.org/chainmaker/common/v2/opencrypto/gmssl/sm4"
	tsm4 "chainmaker.org/chainmaker/common/v2/opencrypto/tencentsm/sm4"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/sym/aes"
	"chainmaker.org/chainmaker/common/v2/crypto/sym/sm4"
)

var (
	errAESKeyLength = errors.New("aes key len must be 128bit，192bit，256bit")
	errSM4KeyLength = errors.New("sm4 key len must be 128bit")
)

func GenerateSymKeyStr(keyType crypto.KeyType, keyHex string) (crypto.SymmetricKey, error) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}

	return GenerateSymKey(keyType, key)
}

func GenerateSymKey(keyType crypto.KeyType, key []byte) (crypto.SymmetricKey, error) {
	bits := len(key) * 8

	switch keyType {
	case crypto.AES:
		if bits != int(crypto.BITS_SIZE_128) &&
			bits != int(crypto.BITS_SIZE_192) && bits != int(crypto.BITS_SIZE_256) {
			return nil, errAESKeyLength
		}
		return &aes.AESKey{Key: key}, nil
	case crypto.SM4:
		if bits != int(crypto.BITS_SIZE_128) {
			return nil, errSM4KeyLength
		}
		switch engine.CryptoEngine {
		case opencrypto.GmSSL:
			return gsm4.SM4Key{Key: key}, nil
		case opencrypto.TencentSM:
			return tsm4.SM4Key{Key: key}, nil
		}
		return &sm4.SM4Key{Key: key}, nil
	default:
		return nil, fmt.Errorf("unsupport symmetric algorithm")
	}
}
