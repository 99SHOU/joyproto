// Code generated by protoc-gen-go. DO NOT EDIT.
// source: framework_msg.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NodeRegisterReq struct {
	NodeType NODE_TYPE `protobuf:"varint,1,opt,name=NodeType,enum=pb.NODE_TYPE" json:"NodeType,omitempty"`
	NodeId   uint32    `protobuf:"varint,2,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *NodeRegisterReq) Reset()                    { *m = NodeRegisterReq{} }
func (m *NodeRegisterReq) String() string            { return proto.CompactTextString(m) }
func (*NodeRegisterReq) ProtoMessage()               {}
func (*NodeRegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NodeRegisterReq) GetNodeType() NODE_TYPE {
	if m != nil {
		return m.NodeType
	}
	return NODE_TYPE_MACHINE
}

func (m *NodeRegisterReq) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type NodeRegisterAck struct {
	NodeType NODE_TYPE `protobuf:"varint,1,opt,name=NodeType,enum=pb.NODE_TYPE" json:"NodeType,omitempty"`
	NodeId   uint32    `protobuf:"varint,2,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *NodeRegisterAck) Reset()                    { *m = NodeRegisterAck{} }
func (m *NodeRegisterAck) String() string            { return proto.CompactTextString(m) }
func (*NodeRegisterAck) ProtoMessage()               {}
func (*NodeRegisterAck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *NodeRegisterAck) GetNodeType() NODE_TYPE {
	if m != nil {
		return m.NodeType
	}
	return NODE_TYPE_MACHINE
}

func (m *NodeRegisterAck) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeRegisterReq)(nil), "pb.NodeRegisterReq")
	proto.RegisterType((*NodeRegisterAck)(nil), "pb.NodeRegisterAck")
}

func init() { proto.RegisterFile("framework_msg.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0x8e, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0x85, 0x88, 0x28, 0x85, 0x70,
	0xf1, 0xfb, 0xe5, 0xa7, 0xa4, 0x06, 0xa5, 0xa6, 0x67, 0x16, 0x97, 0xa4, 0x16, 0x05, 0xa5, 0x16,
	0x0a, 0x69, 0x72, 0x71, 0x80, 0x84, 0x42, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8,
	0x8c, 0x78, 0xf5, 0x0a, 0x92, 0xf4, 0xfc, 0xfc, 0x5d, 0x5c, 0xe3, 0x43, 0x22, 0x03, 0x5c, 0x83,
	0xe0, 0xd2, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0xb6, 0x67, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f,
	0x10, 0x94, 0x87, 0x6e, 0xaa, 0x63, 0x72, 0x36, 0x15, 0x4c, 0x4d, 0x62, 0x03, 0x3b, 0xd9, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x77, 0x3d, 0xf3, 0xdb, 0x00, 0x00, 0x00,
}
