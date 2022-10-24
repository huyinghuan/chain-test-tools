package config

import (
	"fmt"
)

type RtCode struct {
	Code int         `json:"code"`           // 返回码
	Msg  string      `json:"msg"`            // msg
	Data interface{} `json:"data,omitempty"` // 返回的具体数据
}

func (r *RtCode) Error() string {
	return r.Msg
}

const (
	RATE_TYPE_SEND_BARRAGE = "send_live_barrage"
)

const (
	HTTP_CODE_SUCCESS = 200
)

// ----------------------------------------------------
// 					返回码列表
// -----------------------------------------------------
var (
	CODE_OK                           = RtCode{Code: HTTP_CODE_SUCCESS, Msg: "操作成功"}
	CODE_FORBID                       = RtCode{Code: 403, Msg: "无授权"}
	CODE_SYS_ERROR                    = RtCode{Code: -1, Msg: "系统错误"}
	CODE_AUTH_ERROR                   = RtCode{Code: 401, Msg: "签名异常"}
	CODE_NOTFOUND_ERROR               = RtCode{Code: 404, Msg: "页面不存在"}
	CODE_PathParams_ERROR             = RtCode{Code: 406, Msg: "页面不存在"}
	CODE_SYSTEM_500_ERROR             = RtCode{Code: 500, Msg: "页面"}
	CODE_DB_ERROR                     = RtCode{Code: 501, Msg: "系统错误"}
	CODE_Address_ERROR                = RtCode{Code: 2000, Msg: "地址错误"}
	CODE_PARAMS_ERROR                 = RtCode{Code: 2001, Msg: "参数错误"}
	CODE_CLIENT_ERROR                 = RtCode{Code: 2002, Msg: "拿取client失败"}
	CODE_CONTRACT_ERROR               = RtCode{Code: 2003, Msg: "创建合约失败"}
	CODE_USER_REGISTER_ERROR          = RtCode{Code: 2004, Msg: "创建用户失败"}
	CODE_USER_REGISTER_P_ERROR        = RtCode{Code: 2004, Msg: "创建用户失败,类型错误"}
	CODE_USER_CERT_RENEW_ERROR        = RtCode{Code: 2005, Msg: "延长证书时长失败"}
	CODE_USER_CERT_RENEW_P_ERROR      = RtCode{Code: 2006, Msg: "延长证书时长失败"}
	CODE_MINT_ERROR                   = RtCode{Code: 2007, Msg: "创建(mint)NFT失败"}
	CODE_Transfer_ERROR               = RtCode{Code: 2008, Msg: "转移NFT失败"}
	CODE_Balance_ERROR                = RtCode{Code: 2009, Msg: "查询账户余额失败"}
	CODE_Owner_ERROR                  = RtCode{Code: 2010, Msg: "查询归属失败"}
	CODE_TotalSupply_ERROR            = RtCode{Code: 2011, Msg: "发行总量查询失败"}
	CODE_PublicCount_ERROR            = RtCode{Code: 2012, Msg: "已发行总量查询失败"}
	CODE_TokenURI_ERROR               = RtCode{Code: 2013, Msg: "获取TokenURI失败"}
	CODE_BalanceDetails_ERROR         = RtCode{Code: 2014, Msg: "获取详细资产列表失败"}
	CODE_TemplateNoFound_ERROR        = RtCode{Code: 2015, Msg: "合约模版找不到"}
	CODE_QueryTransaction_ERROR       = RtCode{Code: 2016, Msg: "查询交易失败"}
	CODE_UseCrt_ERROR                 = RtCode{Code: 2017, Msg: "用户证书错误"}
	CODE_ContractParams_ERROR         = RtCode{Code: 2018, Msg: "合约参数错误"}
	CODE_QueryTransactionParams_ERROR = RtCode{Code: 2019, Msg: "查询交易参数错误"}
	CODE_PublicCertBody_ERROR         = RtCode{Code: 2020, Msg: "证书内容错误"}
	CODE_PublicCert_TO_NodeId_ERROR   = RtCode{Code: 2021, Msg: "根据证书内容计算nodeid失败"}
	CODE_ChainNode_Cert_TLS_ERROR     = RtCode{Code: 2022, Msg: "生成节点tls证书失败"}
	CODE_ChainNode_Cert_Sign_ERROR    = RtCode{Code: 2023, Msg: "生成节点sign证书失败"}
	CODE_ChainNode_Cert_Type_ERROR    = RtCode{Code: 2023, Msg: "节点类型参数失败"}
	CODE_Burn_ERROR                   = RtCode{Code: 2024, Msg: "销毁NFT失败"}
)

const SEQ_KEY = "seq"

type SC map[string]interface{}

func A2S(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
