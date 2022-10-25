package services

import (
	"bytes"
	"chain-api-imgo/config"
	"chain-api-imgo/forms"
	"chain-api-imgo/resource"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"

	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// var ContractInfoVar = map[int64]models.ContractModel{}

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

// 给请求进行节点管理证书签名
func makeEndorsementEntryList(payload *common.Payload) (endorsementEntryList []*common.EndorsementEntry, err error) {
	crts := config.GetConfig().ChainAdminCrts
	for _, crt := range crts {
		adminKeyBody, err := resource.Get(crt.TlsKeyPath)
		if err != nil {
			return nil, err
		}
		adminCrtsBody, err := resource.Get(crt.TlsCrtPath)
		if err != nil {
			return nil, err
		}
		cert, err := sdkutils.ParseCert(adminCrtsBody)
		if err != nil {
			return nil, err
		}

		hashAlgo, err := bcx509.GetHashFromSignatureAlgorithm(cert.SignatureAlgorithm)
		if err != nil {
			return nil, err
		}
		var orgId string
		if len(cert.Subject.Organization) != 0 {
			orgId = cert.Subject.Organization[0]
		}
		entry, err := sdkutils.MakeEndorser(orgId, hashAlgo, accesscontrol.MemberType_CERT, adminKeyBody, adminCrtsBody, payload)
		if err != nil {
			return nil, err
		}
		endorsementEntryList = append(endorsementEntryList, entry)
	}
	return
}

type CreateContractParams struct {
	ContractName string
	SymboleName  string
	Version      string
	PublishCount uint64
}

func NewContract(client *sdk.ChainClient, abiBytes []byte, binBytes []byte, contractNameHex string, args ...interface{}) (*common.TxResponse, error) {
	myAbi, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, err
	}
	var params []interface{}
	params = append(params, args...)
	dataByte, err := myAbi.Pack("", params...)
	if err != nil {
		return nil, err
	}
	data := hex.EncodeToString(dataByte)
	pairs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}
	//contractNameHex := CalcContractName(contractName)
	resp, err := adminsCreateContract(client, contractNameHex, binBytes, "v1.0", pairs, forms.NFTExtQuery{Sync: true})
	return resp, err
}

/**
私有创建合约
管理员给新合约签名部署
所有合约使用 solidity 语言创建
*/
func adminsCreateContract(client *sdk.ChainClient, contractName string, contractContent []byte, version string, kvs []*common.KeyValuePair, extParams forms.NFTExtQuery) (*common.TxResponse, error) {
	payload, err := client.CreateContractCreatePayload(contractName,
		version,
		string(contractContent),
		common.RuntimeType_EVM,
		kvs)
	if err != nil {
		return nil, err
	}
	endorsementEntrys, err := makeEndorsementEntryList(payload)
	if err != nil {
		return nil, err
	}
	timeout := int64(10)
	if extParams.Timeout > 0 {
		timeout = extParams.Timeout
	}
	resp, err := client.SendContractManageRequest(payload, endorsementEntrys, timeout, extParams.Sync)
	if err != nil {
		return nil, err
	}
	err = CheckProposalRequestResp(resp, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CalcContractName(contractName string) string {
	return hex.EncodeToString(evmutils.Keccak256([]byte(contractName)))[24:]
}

/**
获取用户的链接
OrgId, ChainId
*/
func GetUserChainClient(userKeyBytes, userCertBytes, userSignKeyBytes, userSignCertBytes []byte, args ...string) (*sdk.ChainClient, error) {
	conf := config.GetConfig().ChainNode
	orgId := conf.OrgId
	chainId := conf.ChainId

	if len(args) == 2 && args[0] != "" && args[1] != "" {
		orgId = args[0]
		chainId = args[1]
	}
	nodeOptions := make([]sdk.ChainClientOption, 0)
	nodeOptions = append(nodeOptions, sdk.WithChainClientOrgId(orgId))
	nodeOptions = append(nodeOptions, sdk.WithChainClientChainId(chainId))
	nodeOptions = append(nodeOptions, sdk.WithUserKeyBytes(userKeyBytes))
	nodeOptions = append(nodeOptions, sdk.WithUserCrtBytes(userCertBytes))
	nodeOptions = append(nodeOptions, sdk.WithUserSignKeyBytes(userSignKeyBytes))
	nodeOptions = append(nodeOptions, sdk.WithUserSignCrtBytes(userSignCertBytes))
	nodeOptions = append(nodeOptions, sdk.WithAuthType("permissionedwithcert"))
	// var caCerts []string
	// for _, caPath := range conf.CaPaths {
	// 	caCertsBytes, err := resource.Get(caPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	caCerts = append(caCerts, string(caCertsBytes))
	// }
	for _, remote := range conf.Remotes {
		caCertsBytes, err := resource.Get(remote.Ca)
		if err != nil {
			return nil, err
		}
		node := sdk.NewNodeConfig(
			// 节点地址，格式：127.0.0.1:12301
			sdk.WithNodeAddr(remote.Address),
			// 节点连接数
			sdk.WithNodeConnCnt(conf.ConnCnt),
			// 节点是否启用TLS认证
			sdk.WithNodeUseTLS(conf.Tls),
			// 根证书路径，支持多个
			//sdk.WithNodeCAPaths(nodeConfig.GetCaPaths()),
			sdk.WithNodeCACerts([]string{string(caCertsBytes)}),
			// TLS Hostname
			sdk.WithNodeTLSHostName(conf.TlsHost),
		)
		log.Println("connet node:", remote)
		nodeOptions = append(nodeOptions, sdk.AddChainClientNodeConfig(node))
	}
	//nodeOptions = append(nodeOptions, sdk.WithChainClientLogger(ChainClientLogInstance))
	chainClient, err := sdk.NewChainClient(nodeOptions...)
	if err != nil {
		return nil, err
	}
	return chainClient, nil
}

// func LoadContract(ctx context.Context) {
// 	log := ctx.Value(ContextLogKey).(*goutils.ServerContext)
// 	_ = log

// 	list, err := models.GetContractList()
// 	if err != nil || len(list) < 1 {
// 		log.Info("没有需要重新reload  contract的数据")
// 		return
// 	}

// 	for _, v := range list {
// 		if v.State != 1 {
// 			continue
// 		}
// 		v.TokenAbiByte = []byte(v.TokenAbi)
// 		ContractInfoVar[v.Id] = v
// 	}
// 	log.Info("load")
// }
//
type InvokeParams struct {
	Abi          string
	ContractName string
	Method       string
	Sync         bool
	Timeout      int64
}

func InvokeWithExtParams(client *sdk.ChainClient, params InvokeParams, args ...interface{}) (*common.TxResponse, error) {
	// 计算合约名
	contractName := CalcContractName(params.ContractName)
	// 计算ab参数
	abiObj, _ := abi.JSON(strings.NewReader(params.Abi))
	calldata, err := abiObj.Pack(params.Method, args...) // 失败
	if err != nil {
		return nil, err
	}
	data := hex.EncodeToString(calldata)
	methodName := data[0:8]
	var kvs []*common.KeyValuePair = []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}
	timeout := int64(10)
	if params.Timeout > 0 {
		timeout = params.Timeout
	}
	return client.InvokeContract(contractName,
		methodName,
		"",
		kvs, timeout, params.Sync)
}
func Invoke(client *sdk.ChainClient, abiBytes string, contractName, method string, args ...interface{}) (*common.TxResponse, error) {
	return InvokeWithExtParams(client, InvokeParams{
		Abi:          abiBytes,
		ContractName: contractName,
		Method:       method,
		Sync:         true,
		Timeout:      10,
	}, args...)

}

func InvokeWithABI(client *sdk.ChainClient, abiObj abi.ABI, contractName, method string, args ...interface{}) (*common.TxResponse, error) {
	// 计算合约名
	contractName = CalcContractName(contractName)
	// 计算ab参数
	calldata, err := abiObj.Pack(method, args...) // 失败
	if err != nil {
		return nil, err
	}
	data := hex.EncodeToString(calldata)
	methodName := data[0:8]
	var kvs []*common.KeyValuePair = []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}
	return client.InvokeContract(contractName,
		methodName,
		"",
		kvs, 10, true)

}
