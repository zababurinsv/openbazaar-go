// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	EncryptedMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_MESSAGE            Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_ACK          Message_MessageType = 5
	Message_ORDER_CONFIRMATION Message_MessageType = 6
	Message_RATING             Message_MessageType = 7
	Message_DISPUTE_OPEN       Message_MessageType = 8
	Message_DISPUTE_CLOSE      Message_MessageType = 9
	Message_REFUND             Message_MessageType = 10
)

var Message_MessageType_name = map[int32]string{
	0:  "PING",
	1:  "MESSAGE",
	2:  "FOLLOW",
	3:  "UNFOLLOW",
	4:  "ORDER",
	5:  "ORDER_ACK",
	6:  "ORDER_CONFIRMATION",
	7:  "RATING",
	8:  "DISPUTE_OPEN",
	9:  "DISPUTE_CLOSE",
	10: "REFUND",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"MESSAGE":            1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_ACK":          5,
	"ORDER_CONFIRMATION": 6,
	"RATING":             7,
	"DISPUTE_OPEN":       8,
	"DISPUTE_CLOSE":      9,
	"REFUND":             10,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Message struct {
	MessageType Message_MessageType  `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EncryptedMessage struct {
	EphemPubKey []byte `protobuf:"bytes,1,opt,name=ephemPubKey,proto3" json:"ephemPubKey,omitempty"`
	Ciphertext  []byte `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (m *EncryptedMessage) Reset()                    { *m = EncryptedMessage{} }
func (m *EncryptedMessage) String() string            { return proto.CompactTextString(m) }
func (*EncryptedMessage) ProtoMessage()               {}
func (*EncryptedMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*EncryptedMessage)(nil), "EncryptedMessage")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}

var fileDescriptor2 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x50, 0x4f, 0x6b, 0xfa, 0x40,
	0x14, 0xfc, 0xc5, 0x9f, 0x7f, 0x5f, 0x62, 0xd9, 0x3e, 0x4a, 0xb1, 0x3d, 0x14, 0xf1, 0xd4, 0x53,
	0x84, 0x16, 0x7a, 0x0f, 0xba, 0x4a, 0x50, 0x77, 0xc3, 0x26, 0xd2, 0xa3, 0x44, 0xdd, 0xda, 0x82,
	0x9a, 0xa0, 0x11, 0xba, 0x1f, 0xa7, 0x1f, 0xa0, 0xdf, 0xb1, 0xeb, 0x26, 0x81, 0x9c, 0xf6, 0xcd,
	0x9b, 0x99, 0x37, 0xcb, 0x40, 0xf7, 0x20, 0xcf, 0xe7, 0x78, 0x27, 0xdd, 0xf4, 0x94, 0x64, 0xc9,
	0xe3, 0xc3, 0x2e, 0x49, 0x76, 0x7b, 0x39, 0x34, 0x68, 0x7d, 0xf9, 0x18, 0xc6, 0x47, 0x95, 0x53,
	0x83, 0x9f, 0x1a, 0xb4, 0x16, 0xb9, 0x18, 0xdf, 0xc0, 0x2e, 0x7c, 0x91, 0x4a, 0x65, 0xcf, 0xea,
	0x5b, 0xcf, 0x37, 0x2f, 0x77, 0x6e, 0x41, 0x97, 0xef, 0x95, 0x13, 0x55, 0x21, 0xba, 0xd0, 0x4a,
	0x63, 0xb5, 0x4f, 0xe2, 0x6d, 0xaf, 0xa6, 0x3d, 0xb6, 0xf6, 0xe4, 0x81, 0x6e, 0x19, 0xe8, 0x7a,
	0x47, 0x25, 0x4a, 0xd1, 0xe0, 0xd7, 0x02, 0xbb, 0x72, 0x0c, 0xdb, 0x50, 0x0f, 0x7c, 0x36, 0x25,
	0xff, 0xd0, 0xd6, 0x9f, 0xa1, 0x61, 0xe8, 0x4d, 0x29, 0xb1, 0x10, 0xa0, 0x39, 0xe1, 0xf3, 0x39,
	0x7f, 0x27, 0x35, 0x74, 0xa0, 0xbd, 0x64, 0x05, 0xfa, 0x8f, 0x1d, 0x68, 0x70, 0x31, 0xa6, 0x82,
	0xd4, 0xb1, 0x0b, 0x1d, 0x33, 0xae, 0xbc, 0xd1, 0x8c, 0x34, 0xf0, 0x1e, 0x30, 0x87, 0x23, 0xce,
	0x26, 0xbe, 0x58, 0x78, 0x91, 0xcf, 0x19, 0x69, 0x5e, 0x6f, 0x09, 0x3d, 0xeb, 0x90, 0x16, 0x12,
	0x70, 0xc6, 0x7e, 0x18, 0x2c, 0x23, 0xba, 0xe2, 0x01, 0x65, 0xa4, 0x8d, 0xb7, 0xd0, 0x2d, 0x37,
	0xa3, 0x39, 0x0f, 0x29, 0xe9, 0x18, 0x03, 0x9d, 0x2c, 0xd9, 0x98, 0xc0, 0x20, 0x02, 0x42, 0x8f,
	0x9b, 0x93, 0x4a, 0x33, 0xb9, 0x2d, 0xbb, 0xea, 0x83, 0x2d, 0xd3, 0x4f, 0x79, 0x08, 0x2e, 0xeb,
	0x99, 0x54, 0xa6, 0x2b, 0x47, 0x54, 0x57, 0xf8, 0x04, 0xb0, 0xf9, 0xd2, 0xf8, 0x94, 0xc9, 0xef,
	0xcc, 0x14, 0xe3, 0x88, 0xca, 0x66, 0xdd, 0x34, 0xe5, 0xbc, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0xbc, 0x80, 0x2d, 0xac, 0x01, 0x00, 0x00,
}
