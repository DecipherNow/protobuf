// Code generated by protoc-gen-go.
// source: thrift.proto
// DO NOT EDIT!

/*
Package thrift is a generated protocol buffer package.

It is generated from these files:
	thrift.proto

It has these top-level messages:
*/
package thrift

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_FieldOrd = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50000,
	Name:          "thrift.field_ord",
	Tag:           "varint,50000,opt,name=field_ord,json=fieldOrd",
	Filename:      "thrift.proto",
}

var E_OneofOrd = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.OneofOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50000,
	Name:          "thrift.oneof_ord",
	Tag:           "varint,50000,opt,name=oneof_ord,json=oneofOrd",
	Filename:      "thrift.proto",
}

func init() {
	proto.RegisterExtension(E_FieldOrd)
	proto.RegisterExtension(E_OneofOrd)
}

func init() { proto.RegisterFile("thrift.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc9, 0x28, 0xca,
	0x4c, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x14, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0x95, 0x56, 0x36, 0x5c, 0x9c, 0x69, 0x99, 0xa9, 0x39, 0x29,
	0xf1, 0xf9, 0x45, 0x29, 0x42, 0xb2, 0x7a, 0x10, 0xf5, 0x7a, 0x30, 0xf5, 0x7a, 0x6e, 0x20, 0x39,
	0xff, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x89, 0x0b, 0x6d, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x1c, 0x60, 0x1d, 0xfe, 0x45, 0x29, 0x20, 0xdd, 0xf9, 0x79, 0xa9, 0xf9, 0x69, 0x38, 0x74, 0xfb,
	0x83, 0xe4, 0x30, 0x74, 0x83, 0x75, 0xf8, 0x17, 0xa5, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba,
	0xc4, 0xd2, 0xfe, 0xb4, 0x00, 0x00, 0x00,
}
