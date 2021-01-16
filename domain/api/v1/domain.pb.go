// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/domain.proto

package v1

import (
	domain "domain"
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

type GetRequest struct {
	PortId               string   `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type GetResponse struct {
	Port                 *domain.Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetPort() *domain.Port {
	if m != nil {
		return m.Port
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Ports                []*domain.Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPorts() []*domain.Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type CreateRequest struct {
	Port                 *domain.Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{4}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetPort() *domain.Port {
	if m != nil {
		return m.Port
	}
	return nil
}

type CreateResponse struct {
	Port                 *domain.Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{5}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetPort() *domain.Port {
	if m != nil {
		return m.Port
	}
	return nil
}

type SaveRequest struct {
	Port                 *domain.Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{6}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetPort() *domain.Port {
	if m != nil {
		return m.Port
	}
	return nil
}

type SaveResponse struct {
	Port                 *domain.Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{7}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetPort() *domain.Port {
	if m != nil {
		return m.Port
	}
	return nil
}

type DeleteRequest struct {
	PortId               string   `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be358047ab4d9a95, []int{9}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetRequest)(nil), "api.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "api.v1.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "api.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "api.v1.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "api.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "api.v1.CreateResponse")
	proto.RegisterType((*SaveRequest)(nil), "api.v1.SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "api.v1.SaveResponse")
	proto.RegisterType((*DeleteRequest)(nil), "api.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "api.v1.DeleteResponse")
}

func init() { proto.RegisterFile("api/domain.proto", fileDescriptor_be358047ab4d9a95) }

var fileDescriptor_be358047ab4d9a95 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2c,
	0xc8, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x86, 0x88, 0xa2, 0x48, 0x2a, 0xa9, 0x72, 0x71, 0xb9, 0xa7,
	0x96, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x73, 0xb1, 0x17, 0xe4, 0x17, 0x95,
	0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x4a,
	0xfa, 0x5c, 0xdc, 0x60, 0x65, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x0a, 0x5c, 0x2c, 0x20,
	0x09, 0xb0, 0x22, 0x6e, 0x23, 0x1e, 0x3d, 0xa8, 0x91, 0x01, 0xf9, 0x45, 0x25, 0x41, 0x60, 0x19,
	0x25, 0x5e, 0x2e, 0x6e, 0x9f, 0xcc, 0x62, 0x98, 0xc1, 0x4a, 0x46, 0x5c, 0x3c, 0x10, 0x2e, 0xd4,
	0x00, 0x25, 0x2e, 0x56, 0x90, 0xb2, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0c, 0x13, 0x20, 0x52, 0x4a,
	0x86, 0x5c, 0xbc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x30, 0xd7, 0x11, 0xb6, 0xd5, 0x88, 0x8b,
	0x0f, 0xa6, 0x85, 0x68, 0x97, 0xea, 0x73, 0x71, 0x07, 0x27, 0x96, 0x91, 0x60, 0x89, 0x01, 0x17,
	0x0f, 0x44, 0x03, 0xd1, 0x56, 0x68, 0x70, 0xf1, 0xba, 0xa4, 0xe6, 0xa4, 0x22, 0x7c, 0x82, 0x33,
	0x9c, 0x05, 0xb8, 0xf8, 0x60, 0x2a, 0x21, 0xa6, 0x3b, 0x71, 0x44, 0x81, 0xe2, 0x4f, 0xbf, 0xcc,
	0x30, 0x89, 0x0d, 0x1c, 0x63, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x30, 0x01, 0xcd,
	0xe2, 0x01, 0x00, 0x00,
}
