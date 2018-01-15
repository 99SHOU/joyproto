// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EventDefine.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	EventDefine.proto
	MsgDefine.proto

It has these top-level messages:
	LoginReq
	LoginRespon
	ConnectToGateReq
	ConnectToGateRespon
	RoleInfo
	RoleListReq
	RoleListRespon
	CreateRoleReq
	CreateRoleRespon
	EnterGameReq
	EnterGameRespon
	LeaveGameReq
	LeaveGameRespon
	CreateRoomReq
	CreateRoomRespon
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// couple msg start at 1; odd:Req, even:Respon
// single msg start at 10000
type EGameMsgID int32

const (
	EGameMsgID_EGMI_UNKNOW EGameMsgID = 0
	// couple msg start
	EGameMsgID_EGMI_LOGIN_REQ              EGameMsgID = 1
	EGameMsgID_EGMI_LOGIN_RESPON           EGameMsgID = 2
	EGameMsgID_EGMI_CONNECT_TO_GATE_REQ    EGameMsgID = 3
	EGameMsgID_EGMI_CONNECT_TO_GATE_RESPON EGameMsgID = 4
	// single msg start
	EGameMsgID_EGMI_SINGLE_MSG_START EGameMsgID = 10000
)

var EGameMsgID_name = map[int32]string{
	0:     "EGMI_UNKNOW",
	1:     "EGMI_LOGIN_REQ",
	2:     "EGMI_LOGIN_RESPON",
	3:     "EGMI_CONNECT_TO_GATE_REQ",
	4:     "EGMI_CONNECT_TO_GATE_RESPON",
	10000: "EGMI_SINGLE_MSG_START",
}
var EGameMsgID_value = map[string]int32{
	"EGMI_UNKNOW":                 0,
	"EGMI_LOGIN_REQ":              1,
	"EGMI_LOGIN_RESPON":           2,
	"EGMI_CONNECT_TO_GATE_REQ":    3,
	"EGMI_CONNECT_TO_GATE_RESPON": 4,
	"EGMI_SINGLE_MSG_START":       10000,
}

func (x EGameMsgID) String() string {
	return proto.EnumName(EGameMsgID_name, int32(x))
}
func (EGameMsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SERVER_TYPE int32

const (
	SERVER_TYPE_START   SERVER_TYPE = 0
	SERVER_TYPE_MACHINE SERVER_TYPE = 1
	SERVER_TYPE_LOGIN   SERVER_TYPE = 2
	SERVER_TYPE_GATE    SERVER_TYPE = 3
	SERVER_TYPE_CENTER  SERVER_TYPE = 4
	SERVER_TYPE_LOGIC   SERVER_TYPE = 5
	SERVER_TYPE_END     SERVER_TYPE = 6
)

var SERVER_TYPE_name = map[int32]string{
	0: "START",
	1: "MACHINE",
	2: "LOGIN",
	3: "GATE",
	4: "CENTER",
	5: "LOGIC",
	6: "END",
}
var SERVER_TYPE_value = map[string]int32{
	"START":   0,
	"MACHINE": 1,
	"LOGIN":   2,
	"GATE":    3,
	"CENTER":  4,
	"LOGIC":   5,
	"END":     6,
}

func (x SERVER_TYPE) String() string {
	return proto.EnumName(SERVER_TYPE_name, int32(x))
}
func (SERVER_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterEnum("pb.EGameMsgID", EGameMsgID_name, EGameMsgID_value)
	proto.RegisterEnum("pb.SERVER_TYPE", SERVER_TYPE_name, SERVER_TYPE_value)
}

func init() { proto.RegisterFile("EventDefine.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4a, 0x03, 0x41,
	0x10, 0x85, 0x33, 0x3f, 0x99, 0x68, 0x0d, 0x68, 0xa5, 0x20, 0x20, 0x2a, 0xb8, 0xcf, 0xc2, 0x8d,
	0x27, 0x18, 0x7a, 0x8a, 0xb6, 0x31, 0x53, 0x1d, 0xbb, 0x5b, 0x45, 0x37, 0x8d, 0x81, 0x51, 0x5c,
	0x98, 0x04, 0x0d, 0x9e, 0xc3, 0x3b, 0x78, 0x51, 0x99, 0x16, 0x17, 0x2e, 0xdc, 0x7e, 0xef, 0x7d,
	0x45, 0xf1, 0x60, 0xca, 0x1f, 0xfd, 0x7a, 0xd7, 0xf6, 0x4f, 0x2f, 0xeb, 0xfe, 0x7c, 0xfb, 0xb6,
	0xd9, 0x6d, 0x28, 0xdf, 0xae, 0xe6, 0x5f, 0x19, 0x00, 0xeb, 0xc7, 0xd7, 0xbe, 0x7b, 0x7f, 0x36,
	0x2d, 0x1d, 0x42, 0xcd, 0xba, 0x33, 0xf1, 0x46, 0xae, 0xc4, 0xde, 0xe1, 0x88, 0x08, 0x0e, 0x12,
	0x58, 0x58, 0x6d, 0x24, 0x3a, 0xbe, 0xc6, 0x8c, 0x66, 0x30, 0xfd, 0xc3, 0xfc, 0xd2, 0x0a, 0xe6,
	0x74, 0x0a, 0x47, 0x09, 0x2b, 0x2b, 0xc2, 0x2a, 0xc4, 0x60, 0xa3, 0x6e, 0x02, 0x27, 0xa9, 0xa0,
	0x33, 0x38, 0xf9, 0x27, 0x4d, 0x7a, 0x49, 0xc7, 0x30, 0x4b, 0x05, 0x6f, 0x44, 0x2f, 0x38, 0x76,
	0x5e, 0x47, 0x1f, 0x1a, 0x17, 0xf0, 0x53, 0xe6, 0x0f, 0x50, 0x7b, 0x76, 0xb7, 0xec, 0x62, 0xb8,
	0x5f, 0x32, 0xed, 0xc3, 0xf8, 0x27, 0x1a, 0x51, 0x0d, 0x93, 0xae, 0x51, 0x97, 0x46, 0x18, 0xb3,
	0x81, 0xa7, 0x9f, 0x30, 0xa7, 0x3d, 0x28, 0x87, 0xf3, 0x58, 0x10, 0x40, 0xa5, 0x58, 0x02, 0x3b,
	0x2c, 0x7f, 0x0b, 0x0a, 0xc7, 0x34, 0x81, 0x82, 0xa5, 0xc5, 0x6a, 0x55, 0xa5, 0x31, 0x2e, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xba, 0x57, 0x70, 0x25, 0x21, 0x01, 0x00, 0x00,
}