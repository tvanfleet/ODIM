// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregator.proto

package aggregator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AggregatorRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	RequestBody          []byte   `protobuf:"bytes,2,opt,name=RequestBody,proto3" json:"RequestBody,omitempty"`
	URL                  string   `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatorRequest) Reset()         { *m = AggregatorRequest{} }
func (m *AggregatorRequest) String() string { return proto.CompactTextString(m) }
func (*AggregatorRequest) ProtoMessage()    {}
func (*AggregatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{0}
}

func (m *AggregatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatorRequest.Unmarshal(m, b)
}
func (m *AggregatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatorRequest.Marshal(b, m, deterministic)
}
func (m *AggregatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatorRequest.Merge(m, src)
}
func (m *AggregatorRequest) XXX_Size() int {
	return xxx_messageInfo_AggregatorRequest.Size(m)
}
func (m *AggregatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatorRequest proto.InternalMessageInfo

func (m *AggregatorRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *AggregatorRequest) GetRequestBody() []byte {
	if m != nil {
		return m.RequestBody
	}
	return nil
}

func (m *AggregatorRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type AggregatorResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage        string            `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AggregatorResponse) Reset()         { *m = AggregatorResponse{} }
func (m *AggregatorResponse) String() string { return proto.CompactTextString(m) }
func (*AggregatorResponse) ProtoMessage()    {}
func (*AggregatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{1}
}

func (m *AggregatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatorResponse.Unmarshal(m, b)
}
func (m *AggregatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatorResponse.Marshal(b, m, deterministic)
}
func (m *AggregatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatorResponse.Merge(m, src)
}
func (m *AggregatorResponse) XXX_Size() int {
	return xxx_messageInfo_AggregatorResponse.Size(m)
}
func (m *AggregatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatorResponse proto.InternalMessageInfo

func (m *AggregatorResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *AggregatorResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *AggregatorResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AggregatorResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type RediscoverSystemInventoryRequest struct {
	SystemID             string   `protobuf:"bytes,1,opt,name=SystemID,proto3" json:"SystemID,omitempty"`
	SystemURL            string   `protobuf:"bytes,2,opt,name=SystemURL,proto3" json:"SystemURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RediscoverSystemInventoryRequest) Reset()         { *m = RediscoverSystemInventoryRequest{} }
func (m *RediscoverSystemInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*RediscoverSystemInventoryRequest) ProtoMessage()    {}
func (*RediscoverSystemInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{2}
}

func (m *RediscoverSystemInventoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RediscoverSystemInventoryRequest.Unmarshal(m, b)
}
func (m *RediscoverSystemInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RediscoverSystemInventoryRequest.Marshal(b, m, deterministic)
}
func (m *RediscoverSystemInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RediscoverSystemInventoryRequest.Merge(m, src)
}
func (m *RediscoverSystemInventoryRequest) XXX_Size() int {
	return xxx_messageInfo_RediscoverSystemInventoryRequest.Size(m)
}
func (m *RediscoverSystemInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RediscoverSystemInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RediscoverSystemInventoryRequest proto.InternalMessageInfo

func (m *RediscoverSystemInventoryRequest) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *RediscoverSystemInventoryRequest) GetSystemURL() string {
	if m != nil {
		return m.SystemURL
	}
	return ""
}

type RediscoverSystemInventoryResponse struct {
	TaskURL              string   `protobuf:"bytes,1,opt,name=TaskURL,proto3" json:"TaskURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RediscoverSystemInventoryResponse) Reset()         { *m = RediscoverSystemInventoryResponse{} }
func (m *RediscoverSystemInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*RediscoverSystemInventoryResponse) ProtoMessage()    {}
func (*RediscoverSystemInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{3}
}

func (m *RediscoverSystemInventoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RediscoverSystemInventoryResponse.Unmarshal(m, b)
}
func (m *RediscoverSystemInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RediscoverSystemInventoryResponse.Marshal(b, m, deterministic)
}
func (m *RediscoverSystemInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RediscoverSystemInventoryResponse.Merge(m, src)
}
func (m *RediscoverSystemInventoryResponse) XXX_Size() int {
	return xxx_messageInfo_RediscoverSystemInventoryResponse.Size(m)
}
func (m *RediscoverSystemInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RediscoverSystemInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RediscoverSystemInventoryResponse proto.InternalMessageInfo

func (m *RediscoverSystemInventoryResponse) GetTaskURL() string {
	if m != nil {
		return m.TaskURL
	}
	return ""
}

type UpdateSystemStateRequest struct {
	SystemUUID           string   `protobuf:"bytes,1,opt,name=SystemUUID,proto3" json:"SystemUUID,omitempty"`
	SystemID             string   `protobuf:"bytes,2,opt,name=SystemID,proto3" json:"SystemID,omitempty"`
	SystemURI            string   `protobuf:"bytes,3,opt,name=SystemURI,proto3" json:"SystemURI,omitempty"`
	UpdateKey            string   `protobuf:"bytes,4,opt,name=UpdateKey,proto3" json:"UpdateKey,omitempty"`
	UpdateVal            string   `protobuf:"bytes,5,opt,name=UpdateVal,proto3" json:"UpdateVal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSystemStateRequest) Reset()         { *m = UpdateSystemStateRequest{} }
func (m *UpdateSystemStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSystemStateRequest) ProtoMessage()    {}
func (*UpdateSystemStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{4}
}

func (m *UpdateSystemStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSystemStateRequest.Unmarshal(m, b)
}
func (m *UpdateSystemStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSystemStateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSystemStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSystemStateRequest.Merge(m, src)
}
func (m *UpdateSystemStateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSystemStateRequest.Size(m)
}
func (m *UpdateSystemStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSystemStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSystemStateRequest proto.InternalMessageInfo

func (m *UpdateSystemStateRequest) GetSystemUUID() string {
	if m != nil {
		return m.SystemUUID
	}
	return ""
}

func (m *UpdateSystemStateRequest) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *UpdateSystemStateRequest) GetSystemURI() string {
	if m != nil {
		return m.SystemURI
	}
	return ""
}

func (m *UpdateSystemStateRequest) GetUpdateKey() string {
	if m != nil {
		return m.UpdateKey
	}
	return ""
}

func (m *UpdateSystemStateRequest) GetUpdateVal() string {
	if m != nil {
		return m.UpdateVal
	}
	return ""
}

type UpdateSystemStateResponse struct {
	TaskURL              string   `protobuf:"bytes,1,opt,name=TaskURL,proto3" json:"TaskURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSystemStateResponse) Reset()         { *m = UpdateSystemStateResponse{} }
func (m *UpdateSystemStateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateSystemStateResponse) ProtoMessage()    {}
func (*UpdateSystemStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60785b04c84bec7e, []int{5}
}

func (m *UpdateSystemStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSystemStateResponse.Unmarshal(m, b)
}
func (m *UpdateSystemStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSystemStateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateSystemStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSystemStateResponse.Merge(m, src)
}
func (m *UpdateSystemStateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateSystemStateResponse.Size(m)
}
func (m *UpdateSystemStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSystemStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSystemStateResponse proto.InternalMessageInfo

func (m *UpdateSystemStateResponse) GetTaskURL() string {
	if m != nil {
		return m.TaskURL
	}
	return ""
}

func init() {
	proto.RegisterType((*AggregatorRequest)(nil), "AggregatorRequest")
	proto.RegisterType((*AggregatorResponse)(nil), "AggregatorResponse")
	proto.RegisterMapType((map[string]string)(nil), "AggregatorResponse.HeaderEntry")
	proto.RegisterType((*RediscoverSystemInventoryRequest)(nil), "RediscoverSystemInventoryRequest")
	proto.RegisterType((*RediscoverSystemInventoryResponse)(nil), "RediscoverSystemInventoryResponse")
	proto.RegisterType((*UpdateSystemStateRequest)(nil), "UpdateSystemStateRequest")
	proto.RegisterType((*UpdateSystemStateResponse)(nil), "UpdateSystemStateResponse")
}

func init() { proto.RegisterFile("aggregator.proto", fileDescriptor_60785b04c84bec7e) }

var fileDescriptor_60785b04c84bec7e = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0x8d, 0x21, 0xe4, 0xfb, 0x18, 0xa8, 0x4a, 0x36, 0xb4, 0x35, 0x6e, 0x95, 0x12, 0xab, 0xaa,
	0xb8, 0xf2, 0x05, 0x55, 0xd5, 0xa6, 0x6a, 0xa4, 0xf2, 0xd7, 0x24, 0x6a, 0xa2, 0x48, 0x36, 0xe4,
	0xaa, 0x37, 0x1b, 0x3c, 0x21, 0x08, 0xe3, 0xa5, 0xbb, 0x0b, 0x12, 0x6f, 0xd5, 0x67, 0xea, 0x6b,
	0xf4, 0xa6, 0xf2, 0x1f, 0x98, 0x00, 0xa5, 0x86, 0xbb, 0xdd, 0xb3, 0x3b, 0x67, 0xce, 0x9c, 0x19,
	0xaf, 0x0c, 0x05, 0xda, 0xeb, 0x71, 0xec, 0x51, 0xc9, 0xb8, 0x31, 0xe2, 0x4c, 0x32, 0x7d, 0x00,
	0x87, 0xb5, 0x19, 0x66, 0xe2, 0x8f, 0x31, 0x0a, 0x49, 0x74, 0xc8, 0x5b, 0x28, 0x44, 0x9f, 0xb9,
	0x6d, 0x36, 0x40, 0x57, 0x55, 0xca, 0x4a, 0x25, 0x6b, 0x2e, 0x60, 0xa4, 0x0c, 0xb9, 0xf0, 0x7a,
	0x9d, 0xd9, 0x53, 0x35, 0x55, 0x56, 0x2a, 0x79, 0x33, 0x0e, 0x91, 0x02, 0xa4, 0x3b, 0xe6, 0x95,
	0x9a, 0xf6, 0x83, 0xbd, 0xa5, 0xfe, 0x4b, 0x01, 0x12, 0xcf, 0x26, 0x46, 0xcc, 0x15, 0x48, 0x8e,
	0x01, 0x84, 0xa4, 0x72, 0x2c, 0x1a, 0xcc, 0x46, 0x3f, 0x59, 0xc6, 0x8c, 0x21, 0xe4, 0x0d, 0x3c,
	0x09, 0x76, 0xd7, 0x28, 0x04, 0xed, 0xa1, 0x9f, 0x2c, 0x6b, 0x2e, 0x82, 0xe4, 0x03, 0x1c, 0x3c,
	0x20, 0xb5, 0x91, 0xab, 0xe9, 0x72, 0xba, 0x92, 0xab, 0xbe, 0x36, 0x96, 0x53, 0x19, 0x17, 0xfe,
	0x8d, 0x96, 0x2b, 0xf9, 0xd4, 0x0c, 0xaf, 0x13, 0x02, 0xfb, 0x77, 0x5e, 0x09, 0xfb, 0x7e, 0x09,
	0xfe, 0x5a, 0x3b, 0x85, 0x5c, 0xec, 0xaa, 0x57, 0xca, 0x00, 0xa7, 0xa1, 0x0f, 0xde, 0x92, 0x14,
	0x21, 0x33, 0xa1, 0xce, 0x38, 0xd2, 0x12, 0x6c, 0x3e, 0xa5, 0x3e, 0x2a, 0xfa, 0x77, 0x28, 0x9b,
	0x68, 0xf7, 0x45, 0x97, 0x4d, 0x90, 0x5b, 0x53, 0x21, 0x71, 0x78, 0xe9, 0x4e, 0xd0, 0x95, 0x8c,
	0x4f, 0x23, 0x83, 0x35, 0xf8, 0x3f, 0x3c, 0x69, 0x86, 0xa4, 0xb3, 0x3d, 0x79, 0x05, 0xd9, 0x60,
	0xed, 0x99, 0x17, 0xb0, 0xcf, 0x01, 0xfd, 0x0c, 0x4e, 0xfe, 0xc2, 0x1e, 0x1a, 0xaa, 0xc2, 0x7f,
	0x6d, 0x2a, 0x06, 0x1e, 0x41, 0xc0, 0x1e, 0x6d, 0xf5, 0x9f, 0x0a, 0xa8, 0x9d, 0x91, 0x4d, 0x25,
	0x06, 0xb1, 0x96, 0xa4, 0x12, 0x23, 0x55, 0xc7, 0x00, 0x61, 0xa2, 0xce, 0x4c, 0x57, 0x0c, 0x59,
	0x50, 0x9d, 0x5a, 0xaf, 0xfa, 0x32, 0x6c, 0xf9, 0x1c, 0xf0, 0x4e, 0x83, 0xac, 0xdf, 0x30, 0xf0,
	0x39, 0x6b, 0xce, 0x81, 0xf9, 0xe9, 0x2d, 0x75, 0xd4, 0x4c, 0xfc, 0xf4, 0x96, 0x3a, 0xfa, 0x7b,
	0x28, 0xad, 0x50, 0xbc, 0xa9, 0xd2, 0xea, 0x6f, 0x00, 0x98, 0x0f, 0x00, 0xa9, 0xc3, 0xb3, 0x73,
	0x94, 0x11, 0xd0, 0x67, 0xae, 0x85, 0x7c, 0xd2, 0xef, 0x22, 0x21, 0xc6, 0xd2, 0xfc, 0x6b, 0x47,
	0x2b, 0x46, 0x47, 0xdf, 0x23, 0x55, 0xc8, 0x98, 0x28, 0x50, 0x26, 0x89, 0xf9, 0x02, 0x47, 0x16,
	0xca, 0x26, 0xde, 0xd3, 0xb1, 0x23, 0xeb, 0x8c, 0xc9, 0x1b, 0xee, 0xcf, 0xdc, 0xbf, 0x33, 0xd8,
	0x50, 0x5a, 0xdb, 0x71, 0x72, 0x62, 0x6c, 0x9a, 0x35, 0x4d, 0x37, 0x36, 0x0e, 0x8c, 0xbe, 0x47,
	0xae, 0xe0, 0x70, 0xc9, 0x65, 0x52, 0x32, 0xd6, 0xcd, 0x8a, 0xa6, 0x19, 0x6b, 0x9b, 0xa2, 0xef,
	0x91, 0x1a, 0x14, 0x6b, 0xb6, 0x1d, 0x77, 0x9b, 0x8d, 0x79, 0x32, 0xb3, 0x9b, 0xf0, 0xc2, 0x6b,
	0x98, 0xe3, 0xec, 0xc4, 0x52, 0x83, 0xe2, 0xa3, 0xb6, 0x6f, 0x23, 0x24, 0x28, 0x75, 0x57, 0x96,
	0x26, 0x3a, 0xb8, 0x23, 0xcb, 0x67, 0x78, 0xda, 0xe0, 0x18, 0xd3, 0x92, 0x28, 0xfa, 0x0c, 0x0a,
	0x8b, 0x96, 0xa2, 0x48, 0x12, 0x7e, 0x0a, 0xf9, 0x98, 0x97, 0x49, 0x75, 0x2f, 0x56, 0x9f, 0x28,
	0xba, 0x01, 0xcf, 0x6b, 0xb6, 0xdd, 0x72, 0x70, 0x88, 0xae, 0x14, 0x6d, 0xb6, 0x15, 0xc9, 0x05,
	0xbc, 0x34, 0x71, 0xc8, 0x26, 0x18, 0xf1, 0x7c, 0xe5, 0x6c, 0xb8, 0x15, 0x53, 0x0b, 0x54, 0xff,
	0x19, 0x88, 0x88, 0x6e, 0xee, 0xb7, 0xa2, 0xb1, 0xe0, 0xed, 0x8a, 0x97, 0x61, 0x47, 0xd2, 0xd9,
	0x57, 0xd3, 0x60, 0xae, 0x8b, 0x5d, 0x6f, 0xca, 0xae, 0x51, 0x3e, 0x30, 0x5b, 0x24, 0x7c, 0xb4,
	0xce, 0x51, 0x3e, 0xa6, 0x48, 0xc0, 0x70, 0x77, 0xe0, 0xff, 0x5d, 0xbc, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x78, 0x2b, 0x7e, 0x3d, 0x71, 0x08, 0x00, 0x00,
}