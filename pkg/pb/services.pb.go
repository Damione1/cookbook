// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: services.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1, 0x21, 0x0a, 0x10, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01,
	0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x8c, 0x01, 0x92, 0x41, 0x70, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x4c, 0x6f, 0x67, 0x20, 0x69, 0x6e, 0x20, 0x61,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x4d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x74, 0x68, 0x65, 0x69, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0xa0, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x4e, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x4c, 0x6f, 0x67, 0x20,
	0x6f, 0x75, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2c, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x2d,
	0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01,
	0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x12, 0xc5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x92, 0x41, 0x63, 0x0a, 0x0e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x36, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x61,
	0x6e, 0x20, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x62, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0xb3, 0x01, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x92, 0x41, 0x59, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x3b, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0xaa, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d,
	0x92, 0x41, 0x50, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x6f, 0x66, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xaa, 0x01,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x76, 0x92, 0x41, 0x5a, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18,
	0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x35, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x69, 0x72, 0x20, 0x49, 0x44, 0x2e, 0x62,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xc8, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x81, 0x01, 0x92, 0x41, 0x65, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x41, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x61, 0x20,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73,
	0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0xb0, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x67, 0x92, 0x41, 0x4a, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x27, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0xb1, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74, 0x92, 0x41, 0x57, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x3b, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x70, 0x6f, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xaa, 0x01, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x92, 0x41, 0x50, 0x0a,
	0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x70, 0x6f, 0x73, 0x74, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x2a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61,
	0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xa8, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74,
	0x92, 0x41, 0x58, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x47, 0x65, 0x74, 0x20,
	0x61, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x33, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61,
	0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x20, 0x62,
	0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9b, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x61, 0x92, 0x41, 0x46, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x2b, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x61, 0x6c, 0x6c, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x9b, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5e, 0x92, 0x41, 0x3f, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x27, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x70,
	0x6f, 0x73, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0xbf, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7c, 0x92, 0x41, 0x5d, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x73, 0x12, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x1a, 0x3d, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0xb8, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x92, 0x41, 0x56, 0x0a, 0x07, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x12, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x2c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61,
	0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x1a, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xb4, 0x01,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x92, 0x41, 0x5c, 0x0a, 0x07, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x35, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x20, 0x62,
	0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0xa9, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x73, 0x12, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x1a, 0x2d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20,
	0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0xa9, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x92, 0x41, 0x45, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x73, 0x12, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x1a, 0x29, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x61, 0x20, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x20, 0x66, 0x72,
	0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xdc, 0x01, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8c, 0x01, 0x92,
	0x41, 0x69, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x41, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20,
	0x6e, 0x65, 0x77, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x74,
	0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xd6, 0x01, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x86, 0x01, 0x92, 0x41,
	0x63, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x30, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x1a, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0xd4, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x92, 0x41,
	0x6b, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f,
	0x47, 0x65, 0x74, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x27, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x39, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x20,
	0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xc5, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x92, 0x41, 0x58, 0x0a, 0x0b, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x31, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0xc6, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x77, 0x92, 0x41, 0x52, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x6e,
	0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x2d, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x2a, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x95, 0x04, 0x92,
	0x41, 0xee, 0x03, 0x12, 0x5b, 0x0a, 0x08, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x48, 0x0a, 0x0e, 0x44, 0x61, 0x6d, 0x69, 0x65, 0x6e, 0x20, 0x47, 0x6f, 0x65, 0x68, 0x72, 0x69,
	0x67, 0x12, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x65, 0x31, 0x1a, 0x19,
	0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x40, 0x64, 0x61, 0x6d, 0x69, 0x65, 0x6e, 0x67,
	0x6f, 0x65, 0x68, 0x72, 0x69, 0x67, 0x2e, 0x63, 0x61, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30,
	0x5a, 0xa0, 0x01, 0x0a, 0x9d, 0x01, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x92,
	0x01, 0x08, 0x02, 0x12, 0x7d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20, 0x27, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x20, 0x7b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d,
	0x27, 0x2e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x00, 0x6a, 0x2e, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x6a, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6a, 0x2a, 0x0a, 0x07, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6a, 0x32, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6a, 0x26, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x1d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5a, 0x21, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2d, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: pb.LoginRequest
	(*LogoutRequest)(nil),            // 1: pb.LogoutRequest
	(*RefreshTokenRequest)(nil),      // 2: pb.RefreshTokenRequest
	(*CreateUserRequest)(nil),        // 3: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),        // 4: pb.UpdateUserRequest
	(*GetUserRequest)(nil),           // 5: pb.GetUserRequest
	(*ResetPasswordRequest)(nil),     // 6: pb.ResetPasswordRequest
	(*ChangePasswordRequest)(nil),    // 7: pb.ChangePasswordRequest
	(*CreatePostRequest)(nil),        // 8: pb.CreatePostRequest
	(*UpdatePostRequest)(nil),        // 9: pb.UpdatePostRequest
	(*GetPostRequest)(nil),           // 10: pb.GetPostRequest
	(*ListPostsRequest)(nil),         // 11: pb.ListPostsRequest
	(*DeletePostRequest)(nil),        // 12: pb.DeletePostRequest
	(*CreateRecipeRequest)(nil),      // 13: pb.CreateRecipeRequest
	(*UpdateRecipeRequest)(nil),      // 14: pb.UpdateRecipeRequest
	(*GetRecipeRequest)(nil),         // 15: pb.GetRecipeRequest
	(*ListRecipesRequest)(nil),       // 16: pb.ListRecipesRequest
	(*DeleteRecipeRequest)(nil),      // 17: pb.DeleteRecipeRequest
	(*CreateIngredientRequest)(nil),  // 18: pb.CreateIngredientRequest
	(*UpdateIngredientRequest)(nil),  // 19: pb.UpdateIngredientRequest
	(*GetIngredientRequest)(nil),     // 20: pb.GetIngredientRequest
	(*ListIngredientsRequest)(nil),   // 21: pb.ListIngredientsRequest
	(*DeleteIngredientRequest)(nil),  // 22: pb.DeleteIngredientRequest
	(*LoginResponse)(nil),            // 23: pb.LoginResponse
	(*LogoutResponse)(nil),           // 24: pb.LogoutResponse
	(*RefreshTokenResponse)(nil),     // 25: pb.RefreshTokenResponse
	(*CreateUserResponse)(nil),       // 26: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),       // 27: pb.UpdateUserResponse
	(*GetUserResponse)(nil),          // 28: pb.GetUserResponse
	(*ResetPasswordResponse)(nil),    // 29: pb.ResetPasswordResponse
	(*ChangePasswordResponse)(nil),   // 30: pb.ChangePasswordResponse
	(*CreatePostResponse)(nil),       // 31: pb.CreatePostResponse
	(*UpdatePostResponse)(nil),       // 32: pb.UpdatePostResponse
	(*GetPostResponse)(nil),          // 33: pb.GetPostResponse
	(*ListPostsResponse)(nil),        // 34: pb.ListPostsResponse
	(*DeletePostResponse)(nil),       // 35: pb.DeletePostResponse
	(*CreateRecipeResponse)(nil),     // 36: pb.CreateRecipeResponse
	(*UpdateRecipeResponse)(nil),     // 37: pb.UpdateRecipeResponse
	(*GetRecipeResponse)(nil),        // 38: pb.GetRecipeResponse
	(*ListRecipesResponse)(nil),      // 39: pb.ListRecipesResponse
	(*DeleteRecipeResponse)(nil),     // 40: pb.DeleteRecipeResponse
	(*CreateIngredientResponse)(nil), // 41: pb.CreateIngredientResponse
	(*UpdateIngredientResponse)(nil), // 42: pb.UpdateIngredientResponse
	(*GetIngredientResponse)(nil),    // 43: pb.GetIngredientResponse
	(*ListIngredientsResponse)(nil),  // 44: pb.ListIngredientsResponse
	(*DeleteIngredientResponse)(nil), // 45: pb.DeleteIngredientResponse
}
var file_services_proto_depIdxs = []int32{
	0,  // 0: pb.PortfolioService.LoginUser:input_type -> pb.LoginRequest
	1,  // 1: pb.PortfolioService.LogoutUser:input_type -> pb.LogoutRequest
	2,  // 2: pb.PortfolioService.RefreshToken:input_type -> pb.RefreshTokenRequest
	3,  // 3: pb.PortfolioService.CreateUser:input_type -> pb.CreateUserRequest
	4,  // 4: pb.PortfolioService.UpdateUser:input_type -> pb.UpdateUserRequest
	5,  // 5: pb.PortfolioService.GetUser:input_type -> pb.GetUserRequest
	6,  // 6: pb.PortfolioService.ResetPassword:input_type -> pb.ResetPasswordRequest
	7,  // 7: pb.PortfolioService.ChangePassword:input_type -> pb.ChangePasswordRequest
	8,  // 8: pb.PortfolioService.CreatePost:input_type -> pb.CreatePostRequest
	9,  // 9: pb.PortfolioService.UpdatePost:input_type -> pb.UpdatePostRequest
	10, // 10: pb.PortfolioService.GetPost:input_type -> pb.GetPostRequest
	11, // 11: pb.PortfolioService.ListPosts:input_type -> pb.ListPostsRequest
	12, // 12: pb.PortfolioService.DeletePost:input_type -> pb.DeletePostRequest
	13, // 13: pb.PortfolioService.CreateRecipe:input_type -> pb.CreateRecipeRequest
	14, // 14: pb.PortfolioService.UpdateRecipe:input_type -> pb.UpdateRecipeRequest
	15, // 15: pb.PortfolioService.GetRecipe:input_type -> pb.GetRecipeRequest
	16, // 16: pb.PortfolioService.ListRecipes:input_type -> pb.ListRecipesRequest
	17, // 17: pb.PortfolioService.DeleteRecipe:input_type -> pb.DeleteRecipeRequest
	18, // 18: pb.PortfolioService.CreateIngredient:input_type -> pb.CreateIngredientRequest
	19, // 19: pb.PortfolioService.UpdateIngredient:input_type -> pb.UpdateIngredientRequest
	20, // 20: pb.PortfolioService.GetIngredient:input_type -> pb.GetIngredientRequest
	21, // 21: pb.PortfolioService.ListIngredients:input_type -> pb.ListIngredientsRequest
	22, // 22: pb.PortfolioService.DeleteIngredient:input_type -> pb.DeleteIngredientRequest
	23, // 23: pb.PortfolioService.LoginUser:output_type -> pb.LoginResponse
	24, // 24: pb.PortfolioService.LogoutUser:output_type -> pb.LogoutResponse
	25, // 25: pb.PortfolioService.RefreshToken:output_type -> pb.RefreshTokenResponse
	26, // 26: pb.PortfolioService.CreateUser:output_type -> pb.CreateUserResponse
	27, // 27: pb.PortfolioService.UpdateUser:output_type -> pb.UpdateUserResponse
	28, // 28: pb.PortfolioService.GetUser:output_type -> pb.GetUserResponse
	29, // 29: pb.PortfolioService.ResetPassword:output_type -> pb.ResetPasswordResponse
	30, // 30: pb.PortfolioService.ChangePassword:output_type -> pb.ChangePasswordResponse
	31, // 31: pb.PortfolioService.CreatePost:output_type -> pb.CreatePostResponse
	32, // 32: pb.PortfolioService.UpdatePost:output_type -> pb.UpdatePostResponse
	33, // 33: pb.PortfolioService.GetPost:output_type -> pb.GetPostResponse
	34, // 34: pb.PortfolioService.ListPosts:output_type -> pb.ListPostsResponse
	35, // 35: pb.PortfolioService.DeletePost:output_type -> pb.DeletePostResponse
	36, // 36: pb.PortfolioService.CreateRecipe:output_type -> pb.CreateRecipeResponse
	37, // 37: pb.PortfolioService.UpdateRecipe:output_type -> pb.UpdateRecipeResponse
	38, // 38: pb.PortfolioService.GetRecipe:output_type -> pb.GetRecipeResponse
	39, // 39: pb.PortfolioService.ListRecipes:output_type -> pb.ListRecipesResponse
	40, // 40: pb.PortfolioService.DeleteRecipe:output_type -> pb.DeleteRecipeResponse
	41, // 41: pb.PortfolioService.CreateIngredient:output_type -> pb.CreateIngredientResponse
	42, // 42: pb.PortfolioService.UpdateIngredient:output_type -> pb.UpdateIngredientResponse
	43, // 43: pb.PortfolioService.GetIngredient:output_type -> pb.GetIngredientResponse
	44, // 44: pb.PortfolioService.ListIngredients:output_type -> pb.ListIngredientsResponse
	45, // 45: pb.PortfolioService.DeleteIngredient:output_type -> pb.DeleteIngredientResponse
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	file_user_proto_init()
	file_post_proto_init()
	file_recipe_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
