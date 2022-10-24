// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: syscontract/cert_manage.proto

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

// methods of certificate management
type CertManageFunction int32

const (
	// add one certificate
	CertManageFunction_CERT_ADD CertManageFunction = 0
	// delete certificates
	CertManageFunction_CERTS_DELETE CertManageFunction = 1
	// query certificates
	CertManageFunction_CERTS_QUERY CertManageFunction = 2
	// freeze certificates
	CertManageFunction_CERTS_FREEZE CertManageFunction = 3
	// unfreeze certificates
	CertManageFunction_CERTS_UNFREEZE CertManageFunction = 4
	// revoke certificates
	CertManageFunction_CERTS_REVOKE CertManageFunction = 5
)

var CertManageFunction_name = map[int32]string{
	0: "CERT_ADD",
	1: "CERTS_DELETE",
	2: "CERTS_QUERY",
	3: "CERTS_FREEZE",
	4: "CERTS_UNFREEZE",
	5: "CERTS_REVOKE",
}

var CertManageFunction_value = map[string]int32{
	"CERT_ADD":       0,
	"CERTS_DELETE":   1,
	"CERTS_QUERY":    2,
	"CERTS_FREEZE":   3,
	"CERTS_UNFREEZE": 4,
	"CERTS_REVOKE":   5,
}

func (x CertManageFunction) String() string {
	return proto.EnumName(CertManageFunction_name, int32(x))
}

func (CertManageFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_66924a6b3678c62f, []int{0}
}

func init() {
	proto.RegisterEnum("syscontract.CertManageFunction", CertManageFunction_name, CertManageFunction_value)
}

func init() { proto.RegisterFile("syscontract/cert_manage.proto", fileDescriptor_66924a6b3678c62f) }

var fileDescriptor_66924a6b3678c62f = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xae, 0x2c, 0x4e,
	0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x4f, 0x4e, 0x2d, 0x2a, 0x89, 0xcf, 0x4d, 0xcc,
	0x4b, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x92, 0xd6, 0xaa, 0xe5,
	0x12, 0x72, 0x4e, 0x2d, 0x2a, 0xf1, 0x05, 0x2b, 0x70, 0x2b, 0xcd, 0x4b, 0x2e, 0xc9, 0xcc, 0xcf,
	0x13, 0xe2, 0xe1, 0xe2, 0x70, 0x76, 0x0d, 0x0a, 0x89, 0x77, 0x74, 0x71, 0x11, 0x60, 0x10, 0x12,
	0xe0, 0xe2, 0x01, 0xf1, 0x82, 0xe3, 0x5d, 0x5c, 0x7d, 0x5c, 0x43, 0x5c, 0x05, 0x18, 0x85, 0xf8,
	0xb9, 0xb8, 0x21, 0x22, 0x81, 0xa1, 0xae, 0x41, 0x91, 0x02, 0x4c, 0x08, 0x25, 0x6e, 0x41, 0xae,
	0xae, 0x51, 0xae, 0x02, 0xcc, 0x42, 0x42, 0x5c, 0x7c, 0x10, 0x91, 0x50, 0x3f, 0xa8, 0x18, 0x0b,
	0x42, 0x55, 0x90, 0x6b, 0x98, 0xbf, 0xb7, 0xab, 0x00, 0xab, 0x53, 0xfa, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xc9, 0xe6, 0x17, 0xa5, 0xeb, 0x25, 0x67, 0x24, 0x66, 0xe6,
	0xe5, 0x26, 0x66, 0xa7, 0x16, 0xe9, 0x15, 0x24, 0xe9, 0x21, 0xb9, 0x3b, 0x0a, 0x59, 0x2a, 0xbf,
	0x28, 0x5d, 0x1f, 0xc1, 0xd5, 0x2f, 0x48, 0xd2, 0x4d, 0xcf, 0xd7, 0x2f, 0x33, 0xd2, 0x47, 0x52,
	0x9f, 0xc4, 0x06, 0xf6, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x12, 0x0b, 0xd3, 0xef, 0x1c,
	0x01, 0x00, 0x00,
}