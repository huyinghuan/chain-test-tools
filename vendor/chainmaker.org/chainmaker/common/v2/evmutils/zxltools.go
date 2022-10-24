/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package evmutils

import (
	"crypto/elliptic"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/tjfoc/gmsm/sm3"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"chainmaker.org/chainmaker/common/v2/crypto/asym/ecdsa"
	"chainmaker.org/chainmaker/common/v2/crypto/asym/rsa"
	"chainmaker.org/chainmaker/common/v2/crypto/asym/sm2"
	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
)

const (
	ZXAddressLength    = 22
	ZXAddrPrefixLength = 2
	ZXAddrSuffixLength = 20
	ZXAddrPrefix       = "ZX"
)

// rsaPublicKey reflects the ASN.1 structure of a PKCS#1 public key.
type rsaPublicKeyASN struct {
	N *big.Int
	E int
}

// ZXAddressFromPublicKey computes the address of the given public key object in Zhi Xin Lian format
func ZXAddressFromPublicKey(pk crypto.PublicKey) (string, error) {
	var pkBytes []byte
	var err error
	switch k := pk.(type) {
	case *sm2.PublicKey:
		pkBytes = elliptic.Marshal(k.K.Curve, k.K.X, k.K.Y)
	case *ecdsa.PublicKey:
		pkBytes = elliptic.Marshal(k.K.Curve, k.K.X, k.K.Y)
	case *rsa.PublicKey:
		pkBytes, err = asn1.Marshal(rsaPublicKeyASN{
			N: k.K.N,
			E: k.K.E,
		})
		if err != nil {
			return "", fmt.Errorf("fail to marshal RSA public key: %v", err)
		}
	default:
		return "", fmt.Errorf("unsupported public key type [%T]", k)
	}
	sm3Hash := sm3.New()
	_, err = sm3Hash.Write(pkBytes)
	if err != nil {
		return "", err
	}
	pkDgst := sm3Hash.Sum(nil)
	if len(pkDgst) <= ZXAddrSuffixLength {
		return "", fmt.Errorf("invalid public key")
	}
	addrBytes := pkDgst[:ZXAddrSuffixLength]
	addrHex := hex.EncodeToString(addrBytes)
	return ZXAddrPrefix + addrHex, nil
}

// ZXAddressFromPublicKeyDER computes the address in Zhi Xin Lian format from a public key DER
func ZXAddressFromPublicKeyDER(pkDER []byte) (string, error) {
	pk, err := asym.PublicKeyFromDER(pkDER)
	if err != nil {
		return "", fmt.Errorf("fail to resolve public key from DER format: %v", err)
	}
	return ZXAddressFromPublicKey(pk)
}

// ZXAddressFromPublicKeyPEM computes the address in Zhi Xin Lian format from a public key PEM
func ZXAddressFromPublicKeyPEM(pkPEM []byte) (string, error) {
	pemBlock, _ := pem.Decode(pkPEM)
	if pemBlock == nil {
		return "", fmt.Errorf("fail to resolve public key from PEM string")
	}

	return ZXAddressFromPublicKeyDER(pemBlock.Bytes)
}

// ZXAddressFromCertificatePEM computes the address in Zhi Xin Lian format from a certificate PEM
func ZXAddressFromCertificatePEM(certPEM []byte) (string, error) {
	pemBlock, _ := pem.Decode(certPEM)
	if pemBlock == nil {
		return "", fmt.Errorf("fail to resolve certificate from ")
	}

	cert, err := bcx509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("fail to resolve certificate from PEM format: %v", err)
	}

	return ZXAddressFromPublicKey(cert.PublicKey)
}

// ZXAddressFromCertificatePath computes the address in Zhi Xin Lian format from a certificate file path
func ZXAddressFromCertificatePath(certPath string) (string, error) {
	certContent, err := ioutil.ReadFile(certPath)
	if err != nil {
		return "", fmt.Errorf("fail to load certificate from file [%s]: %v", certPath, err)
	}

	return ZXAddressFromCertificatePEM(certContent)
}
