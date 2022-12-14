/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto/hash"
	commonPb "chainmaker.org/chainmaker/pb-go/v2/common"
	configPb "chainmaker.org/chainmaker/pb-go/v2/config"
	"chainmaker.org/chainmaker/pb-go/v2/consensus"
	"chainmaker.org/chainmaker/pb-go/v2/syscontract"
	"github.com/gogo/protobuf/proto"
)

// default timestamp is "2020-11-30 0:0:0"
const (
	defaultTimestamp           = int64(1606669261)
	errMsgMarshalChainConfFail = "proto marshal chain config failed, %s"
)

// CreateGenesis create genesis block (with read-write set) based on chain config
func CreateGenesis(cc *configPb.ChainConfig) (*commonPb.Block, []*commonPb.TxRWSet, error) {
	var (
		err      error
		tx       *commonPb.Transaction
		rwSet    *commonPb.TxRWSet
		txHash   []byte
		hashType = cc.Crypto.Hash
	)

	// generate config tx, read-write set, and hash
	if tx, err = genConfigTx(cc); err != nil {
		return nil, nil, fmt.Errorf("create genesis config tx failed, %s", err)
	}

	if rwSet, err = genConfigTxRWSet(cc); err != nil {
		return nil, nil, fmt.Errorf("create genesis config tx read-write set failed, %s", err)
	}

	if tx.Result.RwSetHash, err = CalcRWSetHash(cc.Crypto.Hash, rwSet); err != nil {
		return nil, nil, fmt.Errorf("calculate genesis config tx read-write set hash failed, %s", err)
	}

	if txHash, err = CalcTxHash(cc.Crypto.Hash, tx); err != nil {
		return nil, nil, fmt.Errorf("calculate tx hash failed, %s", err)
	}

	// generate genesis block
	genesisBlock := &commonPb.Block{
		Header: &commonPb.BlockHeader{
			ChainId:        cc.ChainId,
			BlockHeight:    0,
			BlockType:      commonPb.BlockType_CONFIG_BLOCK,
			PreBlockHash:   nil,
			BlockHash:      nil,
			PreConfHeight:  0,
			BlockVersion:   getBlockHeaderVersion(cc.Version),
			DagHash:        nil,
			RwSetRoot:      nil,
			TxRoot:         nil,
			BlockTimestamp: defaultTimestamp,
			Proposer:       nil,
			ConsensusArgs:  nil,
			TxCount:        1,
			Signature:      nil,
		},
		Dag: &commonPb.DAG{
			Vertexes: []*commonPb.DAG_Neighbor{
				{
					Neighbors: nil,
				},
			},
		},
		Txs: []*commonPb.Transaction{tx},
	}

	if genesisBlock.Header.TxRoot, err = hash.GetMerkleRoot(hashType, [][]byte{txHash}); err != nil {
		return nil, nil, fmt.Errorf("calculate genesis block tx root failed, %s", err)
	}

	if genesisBlock.Header.RwSetRoot, err = CalcRWSetRoot(hashType, genesisBlock.Txs); err != nil {
		return nil, nil, fmt.Errorf("calculate genesis block rwset root failed, %s", err)
	}

	if genesisBlock.Header.DagHash, err = CalcDagHash(hashType, genesisBlock.Dag); err != nil {
		return nil, nil, fmt.Errorf("calculate genesis block DAG hash failed, %s", err)
	}

	if genesisBlock.Header.BlockHash, err = CalcBlockHash(hashType, genesisBlock); err != nil {
		return nil, nil, fmt.Errorf("calculate genesis block hash failed, %s", err)
	}

	return genesisBlock, []*commonPb.TxRWSet{rwSet}, nil
}
func getBlockHeaderVersion(cfgVersion string) uint32 {
	if version, ok := specialVersionMapping[cfgVersion]; ok {
		return version
	}
	if cfgVersion > "v2.2.0" {
		version := string(cfgVersion[1]) + string(cfgVersion[3]) + string(cfgVersion[5])
		if strings.HasSuffix(cfgVersion, ".0") {
			//??????????????????????????????????????????xxx1
			version += "1"
		} else {
			//??????v2.2.0_alpha?????????v2.3.1???????????????
			version += "0"
		}

		v, err := strconv.Atoi(version)
		if err != nil {
			panic(err)
		}
		return uint32(v)
	}
	return 20
}

//?????????????????????????????????
var specialVersionMapping = map[string]uint32{
	"v2.2.0_alpha": 220,
	"v2.2.0":       2201,
}

func genConfigTx(cc *configPb.ChainConfig) (*commonPb.Transaction, error) {
	var (
		err     error
		ccBytes []byte
		//payloadBytes []byte
	)

	if cc.Version == "v2.2.0_alpha" {
		cc.Block.TxParameterSize = 10
	}

	if ccBytes, err = proto.Marshal(cc); err != nil {
		return nil, fmt.Errorf(errMsgMarshalChainConfFail, err.Error())
	}

	payload := &commonPb.Payload{
		ChainId:      cc.ChainId,
		ContractName: syscontract.SystemContract_CHAIN_CONFIG.String(),
		Method:       "Genesis",
		Parameters:   make([]*commonPb.KeyValuePair, 0),
		Sequence:     cc.Sequence,
		TxType:       commonPb.TxType_INVOKE_CONTRACT,
		TxId:         GetTxIdWithSeed(defaultTimestamp),
		Timestamp:    defaultTimestamp,
	}
	payload.Parameters = append(payload.Parameters, &commonPb.KeyValuePair{
		Key:   syscontract.SystemContract_CHAIN_CONFIG.String(),
		Value: []byte(cc.String()),
	})

	//if payloadBytes, err = proto.Marshal(payload); err != nil {
	//	return nil, fmt.Errorf(errMsgMarshalChainConfFail, err.Error())
	//}

	tx := &commonPb.Transaction{
		Payload: payload,
		Result: &commonPb.Result{
			Code: commonPb.TxStatusCode_SUCCESS,
			ContractResult: &commonPb.ContractResult{
				Code: uint32(0),

				Result: ccBytes,
			},
			RwSetHash: nil,
		},
	}

	return tx, nil
}

func genConfigTxRWSet(cc *configPb.ChainConfig) (*commonPb.TxRWSet, error) {
	var (
		err         error
		ccBytes     []byte
		erc20Config *ERC20Config
		stakeConfig *StakeConfig
	)
	if cc.Consensus.Type == consensus.ConsensusType_DPOS {
		if erc20Config, stakeConfig, err = getDPosConfig(cc); err != nil {
			return nil, err
		}
	}

	if ccBytes, err = proto.Marshal(cc); err != nil {
		return nil, fmt.Errorf(errMsgMarshalChainConfFail, err.Error())
	}
	rwSets, err := totalTxRWSet(cc, ccBytes, erc20Config, stakeConfig)
	if err != nil {
		return nil, err
	}
	set := &commonPb.TxRWSet{
		TxId:     GetTxIdWithSeed(defaultTimestamp),
		TxReads:  nil,
		TxWrites: rwSets,
	}
	return set, nil
}

func totalTxRWSet(cc *configPb.ChainConfig, chainConfigBytes []byte,
	erc20Config *ERC20Config, stakeConfig *StakeConfig) (
	[]*commonPb.TxWrite, error) {
	txWrites := make([]*commonPb.TxWrite, 0)
	txWrites = append(txWrites, &commonPb.TxWrite{
		Key:          []byte(syscontract.SystemContract_CHAIN_CONFIG.String()),
		Value:        chainConfigBytes,
		ContractName: syscontract.SystemContract_CHAIN_CONFIG.String(),
	})
	if erc20Config != nil {
		erc20ConfigTxWrites := erc20Config.toTxWrites()
		txWrites = append(txWrites, erc20ConfigTxWrites...)
	}
	if stakeConfig != nil {
		stakeConfigTxWrites, err := stakeConfig.toTxWrites()
		if err != nil {
			return nil, err
		}
		txWrites = append(txWrites, stakeConfigTxWrites...)
	}
	//????????????????????????Contract????????????
	syscontractKeys := []int{}
	for k := range syscontract.SystemContract_name {
		syscontractKeys = append(syscontractKeys, int(k))
	}
	sort.Ints(syscontractKeys)
	for _, k := range syscontractKeys {
		name := syscontract.SystemContract_name[int32(k)]
		if (name == syscontract.SystemContract_T.String() ||
			name == syscontract.SystemContract_ACCOUNT_MANAGER.String()) &&
			cc.Version < "v2.2.0" {
			continue
		}

		nameWrite, addrWrite := initSysContractTxWrite(name, cc)
		txWrites = append(txWrites, nameWrite)

		if cc.Version >= "v2.3.0" {
			txWrites = append(txWrites, addrWrite)
		}
	}
	return txWrites, nil
}

func initSysContractTxWrite(name string, cc *configPb.ChainConfig) (*commonPb.TxWrite, *commonPb.TxWrite) {
	contract := &commonPb.Contract{
		Name:        name,
		Version:     "v1",
		RuntimeType: commonPb.RuntimeType_NATIVE,
		Status:      commonPb.ContractStatus_NORMAL,
		Creator:     nil,
	}

	if cc.Version >= "v2.2.3" {
		addr, _ := NameToAddrStr(name, cc.Vm.AddrType, getBlockHeaderVersion(cc.Version))
		contract.Address = addr
	}

	data, _ := contract.Marshal()
	nameWrite := &commonPb.TxWrite{
		Key:          GetContractDbKey(name),
		Value:        data,
		ContractName: syscontract.SystemContract_CONTRACT_MANAGE.String(),
	}

	if cc.Version < "v2.3.0" {
		return nameWrite, nil
	}

	addrWrite := &commonPb.TxWrite{
		Key:          GetContractDbKey(contract.Address),
		Value:        data,
		ContractName: syscontract.SystemContract_CONTRACT_MANAGE.String(),
	}

	return nameWrite, addrWrite
}
