package config

import "chainmaker.org/chainmaker/common/v2/crypto"

var Name2HashTypeMap = map[string]crypto.HashType{
	"SM3":      crypto.HASH_TYPE_SM3,
	"SHA256":   crypto.HASH_TYPE_SHA256,
	"SHA3_256": crypto.HASH_TYPE_SHA3_256,
}

type CaType int

const (
	//TLS catype of tls
	TLS CaType = iota + 1
	//SIGN catype of sign
	SIGN
	//SINGLE_ROOT catype of single_root
	SINGLE_ROOT
	//DOUBLE_ROOT catype of double_root
	DOUBLE_ROOT
)

//CaType2NameMap Ca type to string name
var CaType2NameMap = map[CaType]string{
	TLS:         "tls",
	SIGN:        "sign",
	SINGLE_ROOT: "single_root",
	DOUBLE_ROOT: "double_root",
}

//Name2CaTypeMap string name to ca type
var Name2CaTypeMap = map[string]CaType{
	"tls":         TLS,
	"sign":        SIGN,
	"single_root": SINGLE_ROOT,
	"double_root": DOUBLE_ROOT,
}
