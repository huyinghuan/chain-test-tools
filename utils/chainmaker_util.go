package utils

import (
	"chain-api-imgo/resource"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

func CheckProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return errors.New(resp.ContractResult.Message)
	}

	return nil
}

// CreateChainClientWithSDKConf create a chain client with sdk config file path
func CreateChainClientWithSDKConf(sdkConfPath string) (*sdk.ChainClient, error) {
	cc, err := sdk.NewChainClient(
		sdk.WithConfPath(sdkConfPath),
	)
	if err != nil {
		return nil, err
	}

	// Enable certificate compression
	if cc.GetAuthType() == sdk.PermissionedWithCert {
		err = cc.EnableCertHash()
	}
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func CalcContractName(contractName string) string {
	return hex.EncodeToString(evmutils.Keccak256([]byte(contractName)))[24:]
}

func makeEndorsementEntry(keyFile string, certFile string, payload *common.Payload) (endorsementEntry *common.EndorsementEntry, err error) {
	adminKeyBody, err := resource.Get(keyFile)
	if err != nil {
		return
	}

	adminCrtsBody, err := resource.Get(certFile)
	if err != nil {
		return
	}

	cert, err := sdkutils.ParseCert(adminCrtsBody)
	if err != nil {
		return
	}

	hashAlgo, err := bcx509.GetHashFromSignatureAlgorithm(cert.SignatureAlgorithm)
	if err != nil {
		return
	}
	var orgId string
	if len(cert.Subject.Organization) != 0 {
		orgId = cert.Subject.Organization[0]
	}

	return sdkutils.MakeEndorser(orgId, hashAlgo, accesscontrol.MemberType_CERT, adminKeyBody, adminCrtsBody, payload)
}

/**
字符串处理
*/
func ParseAddrAndSkiFromCrt(crtStr string) (addressRaw string, addresss evmutils.Address, ski string, err error) {
	crtBytes := []byte(crtStr)
	return ParseAddrAndSkiFromCrtBytes(crtBytes)
}
func ParseAddrAndSkiFromCrtBytes(crtBytes []byte) (addressRaw string, addresss evmutils.Address, ski string, err error) {
	blockCrt, _ := pem.Decode(crtBytes)
	crt, err := bcx509.ParseCertificate(blockCrt.Bytes)
	if err != nil {
		return
	}
	ski = hex.EncodeToString(crt.SubjectKeyId)
	addrInt, err := evmutils.MakeAddressFromHex(ski)
	if err != nil {
		return
	}
	return fmt.Sprintf("0x%x", addrInt.AsStringKey()), evmutils.BigToAddress(addrInt), ski, nil
}
