package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParseCreatedTimeFromCrt(content string) (int64, error) {
	p, _ := pem.Decode([]byte(content))
	if p == nil {
		return 0, fmt.Errorf("decode error")
	}
	cert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return 0, err
	}
	return cert.NotBefore.Unix(), nil
}
