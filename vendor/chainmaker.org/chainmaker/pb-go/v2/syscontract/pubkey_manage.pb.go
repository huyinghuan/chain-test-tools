// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: syscontract/pubkey_manage.proto

package syscontract

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// methods of pubkey management
type PubkeyManageFunction int32

const (
	// add one pubkey
	PubkeyManageFunction_PUBKEY_ADD PubkeyManageFunction = 0
	// delete pubkeys
	PubkeyManageFunction_PUBKEY_DELETE PubkeyManageFunction = 1
	// query pubkeys
	PubkeyManageFunction_PUBKEY_QUERY PubkeyManageFunction = 2
)

var PubkeyManageFunction_name = map[int32]string{
	0: "PUBKEY_ADD",
	1: "PUBKEY_DELETE",
	2: "PUBKEY_QUERY",
}

var PubkeyManageFunction_value = map[string]int32{
	"PUBKEY_ADD":    0,
	"PUBKEY_DELETE": 1,
	"PUBKEY_QUERY":  2,
}

func (x PubkeyManageFunction) String() string {
	return proto.EnumName(PubkeyManageFunction_name, int32(x))
}

func (PubkeyManageFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_deeb8bb0200c1cb3, []int{0}
}

func init() {
	proto.RegisterEnum("syscontract.PubkeyManageFunction", PubkeyManageFunction_name, PubkeyManageFunction_value)
}

func init() { proto.RegisterFile("syscontract/pubkey_manage.proto", fileDescriptor_deeb8bb0200c1cb3) }

var fileDescriptor_deeb8bb0200c1cb3 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xae, 0x2c, 0x4e,
	0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2f, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x8c, 0xcf,
	0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x52, 0xa0,
	0xe5, 0xcd, 0x25, 0x12, 0x00, 0x56, 0xe3, 0x0b, 0x56, 0xe2, 0x56, 0x9a, 0x97, 0x5c, 0x92, 0x99,
	0x9f, 0x27, 0xc4, 0xc7, 0xc5, 0x15, 0x10, 0xea, 0xe4, 0xed, 0x1a, 0x19, 0xef, 0xe8, 0xe2, 0x22,
	0xc0, 0x20, 0x24, 0xc8, 0xc5, 0x0b, 0xe5, 0xbb, 0xb8, 0xfa, 0xb8, 0x86, 0xb8, 0x0a, 0x30, 0x0a,
	0x09, 0x70, 0xf1, 0x40, 0x85, 0x02, 0x43, 0x5d, 0x83, 0x22, 0x05, 0x98, 0x9c, 0xd2, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x36, 0xbf, 0x28, 0x5d, 0x2f, 0x39, 0x23,
	0x31, 0x33, 0x2f, 0x37, 0x31, 0x3b, 0xb5, 0x48, 0xaf, 0x20, 0x49, 0x0f, 0xc9, 0x15, 0x51, 0xc8,
	0x52, 0xf9, 0x45, 0xe9, 0xfa, 0x08, 0xae, 0x7e, 0x41, 0x92, 0x6e, 0x7a, 0xbe, 0x7e, 0x99, 0x91,
	0x3e, 0x92, 0xfa, 0x24, 0x36, 0xb0, 0x4f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x88,
	0x75, 0x14, 0xec, 0x00, 0x00, 0x00,
}
