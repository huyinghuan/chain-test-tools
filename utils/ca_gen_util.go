package utils

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strconv"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
)

func CheckKeyType(keyTypeStr string) (crypto.KeyType, error) {
	var (
		keyType crypto.KeyType
		ok      bool
	)
	if keyType, ok = crypto.Name2KeyTypeMap[keyTypeStr]; !ok {
		return keyType, fmt.Errorf("check key type failed: key type is unsupport")
	}
	return keyType, nil
}

func CheckHashType(hashTypeStr string) (crypto.HashType, error) {
	var (
		hashType crypto.HashType
		ok       bool
	)
	if hashType, ok = crypto.HashAlgoMap[hashTypeStr]; !ok {
		return hashType, fmt.Errorf("check hash type failed: hash type is unsupport")
	}
	return hashType, nil
}

//Convert extkeyusage to string
func ExtKeyUsageToString(extKeyUsage []x509.ExtKeyUsage) (string, error) {
	var extKeyUsageStr []string
	for _, v := range extKeyUsage {
		vStr := strconv.Itoa(int(v))
		extKeyUsageStr = append(extKeyUsageStr, vStr)
	}
	jsonBytes, err := json.Marshal(extKeyUsageStr)
	if err != nil {
		return "", fmt.Errorf("parse extKeyUsage to string faield: %s", err.Error())
	}
	return string(jsonBytes), nil
}

//ParseCertificate parse cert file to x.509 cert struct
func ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	var (
		cert *bcx509.Certificate
		err  error
	)
	block, rest := pem.Decode(certBytes)
	if block == nil {
		cert, err = bcx509.ParseCertificate(rest)
	} else {
		cert, err = bcx509.ParseCertificate(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse x509 cert failed: %s", err.Error())
	}
	return bcx509.ChainMakerCertToX509Cert(cert)
}

//ParseCsr parse csr file to x.509 cert request
func ParseCsr(csrBytes []byte) (*x509.CertificateRequest, error) {
	var (
		csrBC *bcx509.CertificateRequest
		err   error
	)
	block, rest := pem.Decode(csrBytes)
	if block == nil {
		csrBC, err = bcx509.ParseCertificateRequest(rest)
	} else {
		csrBC, err = bcx509.ParseCertificateRequest(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse certificate request failed: %s", err.Error())
	}
	return bcx509.ChainMakerCertCsrToX509CertCsr(csrBC)
}

//Convert privatekey byte to privatekey
func ParsePrivateKey(privateKeyBytes []byte) (crypto.PrivateKey, error) {
	var (
		privateKey crypto.PrivateKey
		err        error
	)
	block, rest := pem.Decode(privateKeyBytes)
	if block == nil {
		privateKey, err = asym.PrivateKeyFromDER(rest)
	} else {
		privateKey, err = asym.PrivateKeyFromDER(block.Bytes)
	}
	if err != nil {
		return nil, fmt.Errorf("parse private key failed: %s", err.Error())
	}
	return privateKey, nil
}
