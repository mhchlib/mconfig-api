// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/server.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type WatchConfigStreamRequest struct {
	AppKey               string            `protobuf:"bytes,1,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	ConfigKeys           []string          `protobuf:"bytes,2,rep,name=config_keys,json=configKeys,proto3" json:"config_keys,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WatchConfigStreamRequest) Reset()         { *m = WatchConfigStreamRequest{} }
func (m *WatchConfigStreamRequest) String() string { return proto.CompactTextString(m) }
func (*WatchConfigStreamRequest) ProtoMessage()    {}
func (*WatchConfigStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{0}
}

func (m *WatchConfigStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchConfigStreamRequest.Unmarshal(m, b)
}
func (m *WatchConfigStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchConfigStreamRequest.Marshal(b, m, deterministic)
}
func (m *WatchConfigStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchConfigStreamRequest.Merge(m, src)
}
func (m *WatchConfigStreamRequest) XXX_Size() int {
	return xxx_messageInfo_WatchConfigStreamRequest.Size(m)
}
func (m *WatchConfigStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchConfigStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchConfigStreamRequest proto.InternalMessageInfo

func (m *WatchConfigStreamRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *WatchConfigStreamRequest) GetConfigKeys() []string {
	if m != nil {
		return m.ConfigKeys
	}
	return nil
}

func (m *WatchConfigStreamRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ConfigVal struct {
	ConfigKey            string   `protobuf:"bytes,1,opt,name=config_key,json=configKey,proto3" json:"config_key,omitempty"`
	Val                  string   `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigVal) Reset()         { *m = ConfigVal{} }
func (m *ConfigVal) String() string { return proto.CompactTextString(m) }
func (*ConfigVal) ProtoMessage()    {}
func (*ConfigVal) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{1}
}

func (m *ConfigVal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigVal.Unmarshal(m, b)
}
func (m *ConfigVal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigVal.Marshal(b, m, deterministic)
}
func (m *ConfigVal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigVal.Merge(m, src)
}
func (m *ConfigVal) XXX_Size() int {
	return xxx_messageInfo_ConfigVal.Size(m)
}
func (m *ConfigVal) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigVal.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigVal proto.InternalMessageInfo

func (m *ConfigVal) GetConfigKey() string {
	if m != nil {
		return m.ConfigKey
	}
	return ""
}

func (m *ConfigVal) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type WatchConfigStreamResponse struct {
	Configs              []*ConfigVal `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WatchConfigStreamResponse) Reset()         { *m = WatchConfigStreamResponse{} }
func (m *WatchConfigStreamResponse) String() string { return proto.CompactTextString(m) }
func (*WatchConfigStreamResponse) ProtoMessage()    {}
func (*WatchConfigStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{2}
}

func (m *WatchConfigStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchConfigStreamResponse.Unmarshal(m, b)
}
func (m *WatchConfigStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchConfigStreamResponse.Marshal(b, m, deterministic)
}
func (m *WatchConfigStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchConfigStreamResponse.Merge(m, src)
}
func (m *WatchConfigStreamResponse) XXX_Size() int {
	return xxx_messageInfo_WatchConfigStreamResponse.Size(m)
}
func (m *WatchConfigStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchConfigStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchConfigStreamResponse proto.InternalMessageInfo

func (m *WatchConfigStreamResponse) GetConfigs() []*ConfigVal {
	if m != nil {
		return m.Configs
	}
	return nil
}

type GetNodeStoreDataResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeStoreDataResponse) Reset()         { *m = GetNodeStoreDataResponse{} }
func (m *GetNodeStoreDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeStoreDataResponse) ProtoMessage()    {}
func (*GetNodeStoreDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{3}
}

func (m *GetNodeStoreDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeStoreDataResponse.Unmarshal(m, b)
}
func (m *GetNodeStoreDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeStoreDataResponse.Marshal(b, m, deterministic)
}
func (m *GetNodeStoreDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeStoreDataResponse.Merge(m, src)
}
func (m *GetNodeStoreDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetNodeStoreDataResponse.Size(m)
}
func (m *GetNodeStoreDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeStoreDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeStoreDataResponse proto.InternalMessageInfo

func (m *GetNodeStoreDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetNodeStoreDataRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeStoreDataRequest) Reset()         { *m = GetNodeStoreDataRequest{} }
func (m *GetNodeStoreDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeStoreDataRequest) ProtoMessage()    {}
func (*GetNodeStoreDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{4}
}

func (m *GetNodeStoreDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeStoreDataRequest.Unmarshal(m, b)
}
func (m *GetNodeStoreDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeStoreDataRequest.Marshal(b, m, deterministic)
}
func (m *GetNodeStoreDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeStoreDataRequest.Merge(m, src)
}
func (m *GetNodeStoreDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeStoreDataRequest.Size(m)
}
func (m *GetNodeStoreDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeStoreDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeStoreDataRequest proto.InternalMessageInfo

type UpdateConfigRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Env                  string   `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	Config               string   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	Val                  string   `protobuf:"bytes,5,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigRequest) Reset()         { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()    {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{5}
}

func (m *UpdateConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigRequest.Unmarshal(m, b)
}
func (m *UpdateConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigRequest.Merge(m, src)
}
func (m *UpdateConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigRequest.Size(m)
}
func (m *UpdateConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigRequest proto.InternalMessageInfo

func (m *UpdateConfigRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *UpdateConfigRequest) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *UpdateConfigRequest) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *UpdateConfigRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *UpdateConfigRequest) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type UpdateConfiResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfiResponse) Reset()         { *m = UpdateConfiResponse{} }
func (m *UpdateConfiResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConfiResponse) ProtoMessage()    {}
func (*UpdateConfiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{6}
}

func (m *UpdateConfiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfiResponse.Unmarshal(m, b)
}
func (m *UpdateConfiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfiResponse.Marshal(b, m, deterministic)
}
func (m *UpdateConfiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfiResponse.Merge(m, src)
}
func (m *UpdateConfiResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConfiResponse.Size(m)
}
func (m *UpdateConfiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfiResponse proto.InternalMessageInfo

type GetNodeDetailResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeDetailResponse) Reset()         { *m = GetNodeDetailResponse{} }
func (m *GetNodeDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeDetailResponse) ProtoMessage()    {}
func (*GetNodeDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{7}
}

func (m *GetNodeDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeDetailResponse.Unmarshal(m, b)
}
func (m *GetNodeDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetNodeDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeDetailResponse.Merge(m, src)
}
func (m *GetNodeDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetNodeDetailResponse.Size(m)
}
func (m *GetNodeDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeDetailResponse proto.InternalMessageInfo

func (m *GetNodeDetailResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeletConfigRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Env                  string   `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	Config               string   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletConfigRequest) Reset()         { *m = DeletConfigRequest{} }
func (m *DeletConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DeletConfigRequest) ProtoMessage()    {}
func (*DeletConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{8}
}

func (m *DeletConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletConfigRequest.Unmarshal(m, b)
}
func (m *DeletConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletConfigRequest.Marshal(b, m, deterministic)
}
func (m *DeletConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletConfigRequest.Merge(m, src)
}
func (m *DeletConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DeletConfigRequest.Size(m)
}
func (m *DeletConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletConfigRequest proto.InternalMessageInfo

func (m *DeletConfigRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *DeletConfigRequest) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *DeletConfigRequest) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type DeletFilterRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Env                  string   `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletFilterRequest) Reset()         { *m = DeletFilterRequest{} }
func (m *DeletFilterRequest) String() string { return proto.CompactTextString(m) }
func (*DeletFilterRequest) ProtoMessage()    {}
func (*DeletFilterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c2b3ce2e73d961, []int{9}
}

func (m *DeletFilterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletFilterRequest.Unmarshal(m, b)
}
func (m *DeletFilterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletFilterRequest.Marshal(b, m, deterministic)
}
func (m *DeletFilterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletFilterRequest.Merge(m, src)
}
func (m *DeletFilterRequest) XXX_Size() int {
	return xxx_messageInfo_DeletFilterRequest.Size(m)
}
func (m *DeletFilterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletFilterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletFilterRequest proto.InternalMessageInfo

func (m *DeletFilterRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *DeletFilterRequest) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func init() {
	proto.RegisterType((*WatchConfigStreamRequest)(nil), "api.sdk.WatchConfigStreamRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.sdk.WatchConfigStreamRequest.MetadataEntry")
	proto.RegisterType((*ConfigVal)(nil), "api.sdk.ConfigVal")
	proto.RegisterType((*WatchConfigStreamResponse)(nil), "api.sdk.WatchConfigStreamResponse")
	proto.RegisterType((*GetNodeStoreDataResponse)(nil), "api.sdk.GetNodeStoreDataResponse")
	proto.RegisterType((*GetNodeStoreDataRequest)(nil), "api.sdk.GetNodeStoreDataRequest")
	proto.RegisterType((*UpdateConfigRequest)(nil), "api.sdk.UpdateConfigRequest")
	proto.RegisterType((*UpdateConfiResponse)(nil), "api.sdk.UpdateConfiResponse")
	proto.RegisterType((*GetNodeDetailResponse)(nil), "api.sdk.GetNodeDetailResponse")
	proto.RegisterType((*DeletConfigRequest)(nil), "api.sdk.DeletConfigRequest")
	proto.RegisterType((*DeletFilterRequest)(nil), "api.sdk.DeletFilterRequest")
}

func init() { proto.RegisterFile("proto/v1/server.proto", fileDescriptor_84c2b3ce2e73d961) }

var fileDescriptor_84c2b3ce2e73d961 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x50,
	0x0c, 0x5e, 0x96, 0xad, 0xa5, 0xee, 0x26, 0x15, 0x43, 0xb7, 0x2c, 0xe3, 0xa7, 0xcb, 0x55, 0x25,
	0xc6, 0x09, 0x8c, 0x9b, 0x09, 0xb8, 0x82, 0x0e, 0x34, 0x95, 0x21, 0x94, 0x09, 0x10, 0xdc, 0xc0,
	0x69, 0xeb, 0xb6, 0xd1, 0x92, 0xe6, 0x90, 0x9c, 0x46, 0xaa, 0x78, 0x57, 0x5e, 0x80, 0x97, 0x40,
	0xc9, 0x49, 0xd2, 0xff, 0x22, 0x24, 0x2e, 0x22, 0xc5, 0x3e, 0xf6, 0xe7, 0xcf, 0xfe, 0x2c, 0x43,
	0x5d, 0x84, 0x81, 0x0c, 0xec, 0xf8, 0xa9, 0x1d, 0x51, 0x18, 0x53, 0xc8, 0x52, 0x1b, 0xcb, 0x5c,
	0xb8, 0x2c, 0xea, 0xdd, 0x98, 0xc7, 0x83, 0x20, 0x18, 0x78, 0x64, 0xa7, 0xee, 0xce, 0xb8, 0x6f,
	0x93, 0x2f, 0xe4, 0x44, 0x45, 0x59, 0xbf, 0x34, 0x30, 0x3e, 0x73, 0xd9, 0x1d, 0xbe, 0x0e, 0x46,
	0x7d, 0x77, 0x70, 0x2d, 0x43, 0xe2, 0xbe, 0x43, 0x3f, 0xc6, 0x14, 0x49, 0x3c, 0x84, 0x32, 0x17,
	0xe2, 0xdb, 0x0d, 0x4d, 0x0c, 0xad, 0xa1, 0x35, 0x2b, 0x4e, 0x89, 0x0b, 0xd1, 0xa6, 0x09, 0x3e,
	0x84, 0x6a, 0x37, 0x8d, 0x4f, 0xde, 0x22, 0x63, 0xbb, 0xa1, 0x37, 0x2b, 0x0e, 0x28, 0x57, 0x9b,
	0x26, 0x11, 0xb6, 0xe1, 0x96, 0x4f, 0x92, 0xf7, 0xb8, 0xe4, 0x86, 0xde, 0xd0, 0x9b, 0xd5, 0x33,
	0x9b, 0x65, 0x7c, 0xd8, 0xba, 0x72, 0xec, 0x2a, 0xcb, 0xb8, 0x18, 0xc9, 0x70, 0xe2, 0x14, 0x00,
	0xe6, 0x0b, 0xd8, 0x9f, 0x7b, 0xc2, 0x1a, 0xe8, 0x53, 0x4e, 0xc9, 0x2f, 0xde, 0x85, 0xdd, 0x98,
	0x7b, 0x63, 0x32, 0xb6, 0x53, 0x9f, 0x32, 0x9e, 0x6f, 0x9f, 0x6b, 0xd6, 0x4b, 0xa8, 0xa8, 0x5a,
	0x9f, 0xb8, 0x87, 0xf7, 0x01, 0xa6, 0xbc, 0xb3, 0xfc, 0x4a, 0x41, 0x3b, 0xc1, 0x8d, 0xb9, 0x67,
	0xe8, 0x0a, 0x37, 0xe6, 0x9e, 0x75, 0x09, 0x47, 0x2b, 0xe8, 0x46, 0x22, 0x18, 0x45, 0x84, 0xa7,
	0x50, 0x56, 0xb9, 0x91, 0xa1, 0xa5, 0x3d, 0x62, 0xd1, 0x63, 0x51, 0xd2, 0xc9, 0x43, 0x2c, 0x06,
	0xc6, 0x5b, 0x92, 0xef, 0x83, 0x1e, 0x5d, 0xcb, 0x20, 0xa4, 0x16, 0x97, 0xbc, 0x40, 0x42, 0xd8,
	0x49, 0x47, 0x95, 0x30, 0xda, 0x73, 0xd2, 0x7f, 0xeb, 0x08, 0x0e, 0x97, 0xe3, 0xd3, 0x41, 0x59,
	0x3f, 0xe1, 0xce, 0x47, 0xd1, 0xe3, 0x92, 0x54, 0x99, 0x5c, 0xae, 0x1a, 0xe8, 0x5c, 0x88, 0x7c,
	0x2c, 0x5c, 0x88, 0xc4, 0x43, 0xa3, 0x38, 0x1b, 0x4a, 0xf2, 0x8b, 0x07, 0x50, 0x52, 0x84, 0xb2,
	0x2e, 0x33, 0x2b, 0xf1, 0xf7, 0x5d, 0x4f, 0x52, 0x68, 0xec, 0x28, 0xbf, 0xb2, 0xf2, 0x91, 0xec,
	0x4e, 0x47, 0x52, 0x9f, 0x2b, 0x9e, 0xb7, 0x60, 0x3d, 0x82, 0x7a, 0x46, 0xb7, 0x45, 0x92, 0xbb,
	0xde, 0xc6, 0xde, 0x3e, 0x00, 0xb6, 0xc8, 0x23, 0xf9, 0xdf, 0xf8, 0x5b, 0xe7, 0x19, 0xe2, 0x9b,
	0x94, 0xf6, 0x3f, 0x20, 0x9e, 0xfd, 0xd6, 0xa1, 0x7c, 0xa5, 0x88, 0xe0, 0x77, 0xb8, 0xbd, 0x24,
	0x37, 0x9e, 0xfc, 0x75, 0x73, 0x4d, 0x6b, 0x53, 0x48, 0x36, 0xa0, 0xad, 0xa6, 0xf6, 0x44, 0xc3,
	0x2f, 0x50, 0x5b, 0x54, 0x15, 0x1b, 0x45, 0xf6, 0x1a, 0xc1, 0xcd, 0x93, 0x0d, 0x11, 0x39, 0x3c,
	0xbe, 0x83, 0xbd, 0xd9, 0xad, 0xc0, 0x7b, 0x45, 0xd2, 0x8a, 0x65, 0x31, 0x57, 0xbe, 0xce, 0xa0,
	0x5d, 0xc2, 0xfe, 0x9c, 0x9e, 0x78, 0xc0, 0xd4, 0x1d, 0x61, 0xf9, 0x1d, 0x61, 0x17, 0xc9, 0x1d,
	0x31, 0x1f, 0x2c, 0x72, 0x9b, 0xd7, 0xdf, 0xda, 0xc2, 0x16, 0x54, 0x67, 0xd4, 0xc6, 0xe3, 0x22,
	0x61, 0x79, 0x07, 0xcc, 0x35, 0x55, 0x66, 0x50, 0x94, 0xc2, 0x8b, 0x28, 0x73, 0xba, 0xaf, 0x47,
	0x79, 0xc5, 0xbe, 0x9e, 0x0e, 0x5c, 0x39, 0x1c, 0x77, 0x58, 0x37, 0xf0, 0x6d, 0x7f, 0xd8, 0x1d,
	0x7a, 0x6e, 0xc7, 0xf6, 0xd5, 0x16, 0x3d, 0xe6, 0xc2, 0xb5, 0x93, 0xaf, 0xb8, 0xa5, 0x9d, 0x52,
	0x8a, 0xf0, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x46, 0x58, 0xb2, 0x65, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MConfigClient is the client API for MConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MConfigClient interface {
	WatchConfigStream(ctx context.Context, opts ...grpc.CallOption) (MConfig_WatchConfigStreamClient, error)
	GetNodeStoreData(ctx context.Context, in *GetNodeStoreDataRequest, opts ...grpc.CallOption) (*GetNodeStoreDataResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfiResponse, error)
	GetNodeDetail(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetNodeDetailResponse, error)
	DeletConfig(ctx context.Context, in *DeletConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeletFilter(ctx context.Context, in *DeletFilterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type mConfigClient struct {
	cc *grpc.ClientConn
}

func NewMConfigClient(cc *grpc.ClientConn) MConfigClient {
	return &mConfigClient{cc}
}

func (c *mConfigClient) WatchConfigStream(ctx context.Context, opts ...grpc.CallOption) (MConfig_WatchConfigStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MConfig_serviceDesc.Streams[0], "/api.sdk.MConfig/WatchConfigStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &mConfigWatchConfigStreamClient{stream}
	return x, nil
}

type MConfig_WatchConfigStreamClient interface {
	Send(*WatchConfigStreamRequest) error
	Recv() (*WatchConfigStreamResponse, error)
	grpc.ClientStream
}

type mConfigWatchConfigStreamClient struct {
	grpc.ClientStream
}

func (x *mConfigWatchConfigStreamClient) Send(m *WatchConfigStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mConfigWatchConfigStreamClient) Recv() (*WatchConfigStreamResponse, error) {
	m := new(WatchConfigStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mConfigClient) GetNodeStoreData(ctx context.Context, in *GetNodeStoreDataRequest, opts ...grpc.CallOption) (*GetNodeStoreDataResponse, error) {
	out := new(GetNodeStoreDataResponse)
	err := c.cc.Invoke(ctx, "/api.sdk.MConfig/GetNodeStoreData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfiResponse, error) {
	out := new(UpdateConfiResponse)
	err := c.cc.Invoke(ctx, "/api.sdk.MConfig/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigClient) GetNodeDetail(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetNodeDetailResponse, error) {
	out := new(GetNodeDetailResponse)
	err := c.cc.Invoke(ctx, "/api.sdk.MConfig/GetNodeDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigClient) DeletConfig(ctx context.Context, in *DeletConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.sdk.MConfig/DeletConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigClient) DeletFilter(ctx context.Context, in *DeletFilterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.sdk.MConfig/DeletFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MConfigServer is the server API for MConfig service.
type MConfigServer interface {
	WatchConfigStream(MConfig_WatchConfigStreamServer) error
	GetNodeStoreData(context.Context, *GetNodeStoreDataRequest) (*GetNodeStoreDataResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfiResponse, error)
	GetNodeDetail(context.Context, *empty.Empty) (*GetNodeDetailResponse, error)
	DeletConfig(context.Context, *DeletConfigRequest) (*empty.Empty, error)
	DeletFilter(context.Context, *DeletFilterRequest) (*empty.Empty, error)
}

// UnimplementedMConfigServer can be embedded to have forward compatible implementations.
type UnimplementedMConfigServer struct {
}

func (*UnimplementedMConfigServer) WatchConfigStream(srv MConfig_WatchConfigStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchConfigStream not implemented")
}
func (*UnimplementedMConfigServer) GetNodeStoreData(ctx context.Context, req *GetNodeStoreDataRequest) (*GetNodeStoreDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeStoreData not implemented")
}
func (*UnimplementedMConfigServer) UpdateConfig(ctx context.Context, req *UpdateConfigRequest) (*UpdateConfiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (*UnimplementedMConfigServer) GetNodeDetail(ctx context.Context, req *empty.Empty) (*GetNodeDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeDetail not implemented")
}
func (*UnimplementedMConfigServer) DeletConfig(ctx context.Context, req *DeletConfigRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletConfig not implemented")
}
func (*UnimplementedMConfigServer) DeletFilter(ctx context.Context, req *DeletFilterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletFilter not implemented")
}

func RegisterMConfigServer(s *grpc.Server, srv MConfigServer) {
	s.RegisterService(&_MConfig_serviceDesc, srv)
}

func _MConfig_WatchConfigStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MConfigServer).WatchConfigStream(&mConfigWatchConfigStreamServer{stream})
}

type MConfig_WatchConfigStreamServer interface {
	Send(*WatchConfigStreamResponse) error
	Recv() (*WatchConfigStreamRequest, error)
	grpc.ServerStream
}

type mConfigWatchConfigStreamServer struct {
	grpc.ServerStream
}

func (x *mConfigWatchConfigStreamServer) Send(m *WatchConfigStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mConfigWatchConfigStreamServer) Recv() (*WatchConfigStreamRequest, error) {
	m := new(WatchConfigStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MConfig_GetNodeStoreData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeStoreDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MConfigServer).GetNodeStoreData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sdk.MConfig/GetNodeStoreData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MConfigServer).GetNodeStoreData(ctx, req.(*GetNodeStoreDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MConfig_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MConfigServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sdk.MConfig/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MConfigServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MConfig_GetNodeDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MConfigServer).GetNodeDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sdk.MConfig/GetNodeDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MConfigServer).GetNodeDetail(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MConfig_DeletConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MConfigServer).DeletConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sdk.MConfig/DeletConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MConfigServer).DeletConfig(ctx, req.(*DeletConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MConfig_DeletFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MConfigServer).DeletFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sdk.MConfig/DeletFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MConfigServer).DeletFilter(ctx, req.(*DeletFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.sdk.MConfig",
	HandlerType: (*MConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeStoreData",
			Handler:    _MConfig_GetNodeStoreData_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _MConfig_UpdateConfig_Handler,
		},
		{
			MethodName: "GetNodeDetail",
			Handler:    _MConfig_GetNodeDetail_Handler,
		},
		{
			MethodName: "DeletConfig",
			Handler:    _MConfig_DeletConfig_Handler,
		},
		{
			MethodName: "DeletFilter",
			Handler:    _MConfig_DeletFilter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchConfigStream",
			Handler:       _MConfig_WatchConfigStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/v1/server.proto",
}