// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv.proto

package srv

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

type NotifyReadyReq struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Nctxs                uint32   `protobuf:"varint,2,opt,name=nctxs,proto3" json:"nctxs,omitempty"`
	DrpcListenerSock     string   `protobuf:"bytes,3,opt,name=drpcListenerSock,proto3" json:"drpcListenerSock,omitempty"`
	InstanceIdx          uint32   `protobuf:"varint,4,opt,name=instanceIdx,proto3" json:"instanceIdx,omitempty"`
	Ntgts                uint32   `protobuf:"varint,5,opt,name=ntgts,proto3" json:"ntgts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReadyReq) Reset()         { *m = NotifyReadyReq{} }
func (m *NotifyReadyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReadyReq) ProtoMessage()    {}
func (*NotifyReadyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{0}
}

func (m *NotifyReadyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReadyReq.Unmarshal(m, b)
}
func (m *NotifyReadyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReadyReq.Marshal(b, m, deterministic)
}
func (m *NotifyReadyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReadyReq.Merge(m, src)
}
func (m *NotifyReadyReq) XXX_Size() int {
	return xxx_messageInfo_NotifyReadyReq.Size(m)
}
func (m *NotifyReadyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReadyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReadyReq proto.InternalMessageInfo

func (m *NotifyReadyReq) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NotifyReadyReq) GetNctxs() uint32 {
	if m != nil {
		return m.Nctxs
	}
	return 0
}

func (m *NotifyReadyReq) GetDrpcListenerSock() string {
	if m != nil {
		return m.DrpcListenerSock
	}
	return ""
}

func (m *NotifyReadyReq) GetInstanceIdx() uint32 {
	if m != nil {
		return m.InstanceIdx
	}
	return 0
}

func (m *NotifyReadyReq) GetNtgts() uint32 {
	if m != nil {
		return m.Ntgts
	}
	return 0
}

type BioErrorReq struct {
	UnmapErr             bool     `protobuf:"varint,1,opt,name=unmapErr,proto3" json:"unmapErr,omitempty"`
	ReadErr              bool     `protobuf:"varint,2,opt,name=readErr,proto3" json:"readErr,omitempty"`
	WriteErr             bool     `protobuf:"varint,3,opt,name=writeErr,proto3" json:"writeErr,omitempty"`
	TgtId                int32    `protobuf:"varint,4,opt,name=tgtId,proto3" json:"tgtId,omitempty"`
	InstanceIdx          uint32   `protobuf:"varint,5,opt,name=instanceIdx,proto3" json:"instanceIdx,omitempty"`
	DrpcListenerSock     string   `protobuf:"bytes,6,opt,name=drpcListenerSock,proto3" json:"drpcListenerSock,omitempty"`
	Uri                  string   `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BioErrorReq) Reset()         { *m = BioErrorReq{} }
func (m *BioErrorReq) String() string { return proto.CompactTextString(m) }
func (*BioErrorReq) ProtoMessage()    {}
func (*BioErrorReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{1}
}

func (m *BioErrorReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BioErrorReq.Unmarshal(m, b)
}
func (m *BioErrorReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BioErrorReq.Marshal(b, m, deterministic)
}
func (m *BioErrorReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BioErrorReq.Merge(m, src)
}
func (m *BioErrorReq) XXX_Size() int {
	return xxx_messageInfo_BioErrorReq.Size(m)
}
func (m *BioErrorReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BioErrorReq.DiscardUnknown(m)
}

var xxx_messageInfo_BioErrorReq proto.InternalMessageInfo

func (m *BioErrorReq) GetUnmapErr() bool {
	if m != nil {
		return m.UnmapErr
	}
	return false
}

func (m *BioErrorReq) GetReadErr() bool {
	if m != nil {
		return m.ReadErr
	}
	return false
}

func (m *BioErrorReq) GetWriteErr() bool {
	if m != nil {
		return m.WriteErr
	}
	return false
}

func (m *BioErrorReq) GetTgtId() int32 {
	if m != nil {
		return m.TgtId
	}
	return 0
}

func (m *BioErrorReq) GetInstanceIdx() uint32 {
	if m != nil {
		return m.InstanceIdx
	}
	return 0
}

func (m *BioErrorReq) GetDrpcListenerSock() string {
	if m != nil {
		return m.DrpcListenerSock
	}
	return ""
}

func (m *BioErrorReq) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type GetPoolSvcReq struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolSvcReq) Reset()         { *m = GetPoolSvcReq{} }
func (m *GetPoolSvcReq) String() string { return proto.CompactTextString(m) }
func (*GetPoolSvcReq) ProtoMessage()    {}
func (*GetPoolSvcReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{2}
}

func (m *GetPoolSvcReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolSvcReq.Unmarshal(m, b)
}
func (m *GetPoolSvcReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolSvcReq.Marshal(b, m, deterministic)
}
func (m *GetPoolSvcReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolSvcReq.Merge(m, src)
}
func (m *GetPoolSvcReq) XXX_Size() int {
	return xxx_messageInfo_GetPoolSvcReq.Size(m)
}
func (m *GetPoolSvcReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolSvcReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolSvcReq proto.InternalMessageInfo

func (m *GetPoolSvcReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type GetPoolSvcResp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Svcreps              []uint32 `protobuf:"varint,2,rep,packed,name=svcreps,proto3" json:"svcreps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolSvcResp) Reset()         { *m = GetPoolSvcResp{} }
func (m *GetPoolSvcResp) String() string { return proto.CompactTextString(m) }
func (*GetPoolSvcResp) ProtoMessage()    {}
func (*GetPoolSvcResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{3}
}

func (m *GetPoolSvcResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolSvcResp.Unmarshal(m, b)
}
func (m *GetPoolSvcResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolSvcResp.Marshal(b, m, deterministic)
}
func (m *GetPoolSvcResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolSvcResp.Merge(m, src)
}
func (m *GetPoolSvcResp) XXX_Size() int {
	return xxx_messageInfo_GetPoolSvcResp.Size(m)
}
func (m *GetPoolSvcResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolSvcResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolSvcResp proto.InternalMessageInfo

func (m *GetPoolSvcResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetPoolSvcResp) GetSvcreps() []uint32 {
	if m != nil {
		return m.Svcreps
	}
	return nil
}

// following messages support the deprecated C mgmt API and
// maintain compatibility with the iosrv MS
type PoolCreateUpcall struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Svcreps              []uint32 `protobuf:"varint,2,rep,packed,name=svcreps,proto3" json:"svcreps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolCreateUpcall) Reset()         { *m = PoolCreateUpcall{} }
func (m *PoolCreateUpcall) String() string { return proto.CompactTextString(m) }
func (*PoolCreateUpcall) ProtoMessage()    {}
func (*PoolCreateUpcall) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{4}
}

func (m *PoolCreateUpcall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolCreateUpcall.Unmarshal(m, b)
}
func (m *PoolCreateUpcall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolCreateUpcall.Marshal(b, m, deterministic)
}
func (m *PoolCreateUpcall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolCreateUpcall.Merge(m, src)
}
func (m *PoolCreateUpcall) XXX_Size() int {
	return xxx_messageInfo_PoolCreateUpcall.Size(m)
}
func (m *PoolCreateUpcall) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolCreateUpcall.DiscardUnknown(m)
}

var xxx_messageInfo_PoolCreateUpcall proto.InternalMessageInfo

func (m *PoolCreateUpcall) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PoolCreateUpcall) GetSvcreps() []uint32 {
	if m != nil {
		return m.Svcreps
	}
	return nil
}

type PoolDestroyUpcall struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolDestroyUpcall) Reset()         { *m = PoolDestroyUpcall{} }
func (m *PoolDestroyUpcall) String() string { return proto.CompactTextString(m) }
func (*PoolDestroyUpcall) ProtoMessage()    {}
func (*PoolDestroyUpcall) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{5}
}

func (m *PoolDestroyUpcall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolDestroyUpcall.Unmarshal(m, b)
}
func (m *PoolDestroyUpcall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolDestroyUpcall.Marshal(b, m, deterministic)
}
func (m *PoolDestroyUpcall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolDestroyUpcall.Merge(m, src)
}
func (m *PoolDestroyUpcall) XXX_Size() int {
	return xxx_messageInfo_PoolDestroyUpcall.Size(m)
}
func (m *PoolDestroyUpcall) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolDestroyUpcall.DiscardUnknown(m)
}

var xxx_messageInfo_PoolDestroyUpcall proto.InternalMessageInfo

func (m *PoolDestroyUpcall) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type PoolListUpcall struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Npools               uint64   `protobuf:"varint,2,opt,name=npools,proto3" json:"npools,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolListUpcall) Reset()         { *m = PoolListUpcall{} }
func (m *PoolListUpcall) String() string { return proto.CompactTextString(m) }
func (*PoolListUpcall) ProtoMessage()    {}
func (*PoolListUpcall) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{6}
}

func (m *PoolListUpcall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListUpcall.Unmarshal(m, b)
}
func (m *PoolListUpcall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListUpcall.Marshal(b, m, deterministic)
}
func (m *PoolListUpcall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListUpcall.Merge(m, src)
}
func (m *PoolListUpcall) XXX_Size() int {
	return xxx_messageInfo_PoolListUpcall.Size(m)
}
func (m *PoolListUpcall) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListUpcall.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListUpcall proto.InternalMessageInfo

func (m *PoolListUpcall) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *PoolListUpcall) GetNpools() uint64 {
	if m != nil {
		return m.Npools
	}
	return 0
}

// copied from mgmt/pool.proto
type PoolListUpcallResp struct {
	Status               int32                      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Pools                []*PoolListUpcallResp_Pool `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PoolListUpcallResp) Reset()         { *m = PoolListUpcallResp{} }
func (m *PoolListUpcallResp) String() string { return proto.CompactTextString(m) }
func (*PoolListUpcallResp) ProtoMessage()    {}
func (*PoolListUpcallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{7}
}

func (m *PoolListUpcallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListUpcallResp.Unmarshal(m, b)
}
func (m *PoolListUpcallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListUpcallResp.Marshal(b, m, deterministic)
}
func (m *PoolListUpcallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListUpcallResp.Merge(m, src)
}
func (m *PoolListUpcallResp) XXX_Size() int {
	return xxx_messageInfo_PoolListUpcallResp.Size(m)
}
func (m *PoolListUpcallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListUpcallResp.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListUpcallResp proto.InternalMessageInfo

func (m *PoolListUpcallResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PoolListUpcallResp) GetPools() []*PoolListUpcallResp_Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

type PoolListUpcallResp_Pool struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Svcreps              []uint32 `protobuf:"varint,2,rep,packed,name=svcreps,proto3" json:"svcreps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolListUpcallResp_Pool) Reset()         { *m = PoolListUpcallResp_Pool{} }
func (m *PoolListUpcallResp_Pool) String() string { return proto.CompactTextString(m) }
func (*PoolListUpcallResp_Pool) ProtoMessage()    {}
func (*PoolListUpcallResp_Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bbe8325d22c1a26, []int{7, 0}
}

func (m *PoolListUpcallResp_Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolListUpcallResp_Pool.Unmarshal(m, b)
}
func (m *PoolListUpcallResp_Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolListUpcallResp_Pool.Marshal(b, m, deterministic)
}
func (m *PoolListUpcallResp_Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolListUpcallResp_Pool.Merge(m, src)
}
func (m *PoolListUpcallResp_Pool) XXX_Size() int {
	return xxx_messageInfo_PoolListUpcallResp_Pool.Size(m)
}
func (m *PoolListUpcallResp_Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolListUpcallResp_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_PoolListUpcallResp_Pool proto.InternalMessageInfo

func (m *PoolListUpcallResp_Pool) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PoolListUpcallResp_Pool) GetSvcreps() []uint32 {
	if m != nil {
		return m.Svcreps
	}
	return nil
}

func init() {
	proto.RegisterType((*NotifyReadyReq)(nil), "srv.NotifyReadyReq")
	proto.RegisterType((*BioErrorReq)(nil), "srv.BioErrorReq")
	proto.RegisterType((*GetPoolSvcReq)(nil), "srv.GetPoolSvcReq")
	proto.RegisterType((*GetPoolSvcResp)(nil), "srv.GetPoolSvcResp")
	proto.RegisterType((*PoolCreateUpcall)(nil), "srv.PoolCreateUpcall")
	proto.RegisterType((*PoolDestroyUpcall)(nil), "srv.PoolDestroyUpcall")
	proto.RegisterType((*PoolListUpcall)(nil), "srv.PoolListUpcall")
	proto.RegisterType((*PoolListUpcallResp)(nil), "srv.PoolListUpcallResp")
	proto.RegisterType((*PoolListUpcallResp_Pool)(nil), "srv.PoolListUpcallResp.Pool")
}

func init() {
	proto.RegisterFile("srv.proto", fileDescriptor_2bbe8325d22c1a26)
}

var fileDescriptor_2bbe8325d22c1a26 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdd, 0xaa, 0xd3, 0x40,
	0x10, 0xc7, 0x49, 0x93, 0xf4, 0x63, 0x4a, 0x4b, 0x5d, 0x8a, 0x84, 0xe2, 0x45, 0x89, 0x17, 0x06,
	0x2f, 0x7a, 0x51, 0xbd, 0x16, 0xa9, 0x16, 0x29, 0x88, 0xc8, 0x16, 0x1f, 0x20, 0x26, 0x6b, 0x09,
	0xc6, 0xec, 0x3a, 0xbb, 0x89, 0xed, 0x8b, 0x78, 0x75, 0x5e, 0xeb, 0xbc, 0xcf, 0x61, 0x36, 0x49,
	0xe9, 0x39, 0xfd, 0x80, 0x73, 0xb7, 0xbf, 0x9d, 0x99, 0xff, 0xcc, 0x7f, 0xb2, 0x81, 0x81, 0xc6,
	0x6a, 0xa1, 0x50, 0x1a, 0xc9, 0x5c, 0x8d, 0x55, 0x78, 0xe7, 0xc0, 0xf8, 0x9b, 0x34, 0xd9, 0xaf,
	0x03, 0x17, 0x71, 0x7a, 0xe0, 0xe2, 0x2f, 0x9b, 0x80, 0x5b, 0x62, 0x16, 0x38, 0x73, 0x27, 0x1a,
	0x70, 0x3a, 0xb2, 0x29, 0xf8, 0x45, 0x62, 0xf6, 0x3a, 0xe8, 0xcc, 0x9d, 0x68, 0xc4, 0x6b, 0x60,
	0x6f, 0x61, 0x92, 0xa2, 0x4a, 0xbe, 0x66, 0xda, 0x88, 0x42, 0xe0, 0x56, 0x26, 0xbf, 0x03, 0xd7,
	0x16, 0x9d, 0xdd, 0xb3, 0x39, 0x0c, 0xb3, 0x42, 0x9b, 0xb8, 0x48, 0xc4, 0x26, 0xdd, 0x07, 0x9e,
	0xd5, 0x39, 0xbd, 0xb2, 0x3d, 0xcc, 0xce, 0xe8, 0xc0, 0x6f, 0x7a, 0x10, 0x84, 0xf7, 0x0e, 0x0c,
	0x57, 0x99, 0x5c, 0x23, 0x4a, 0xa4, 0xd9, 0x66, 0xd0, 0x2f, 0x8b, 0x3f, 0xb1, 0x5a, 0x23, 0xda,
	0x01, 0xfb, 0xfc, 0xc8, 0x2c, 0x80, 0x1e, 0x8a, 0x38, 0xa5, 0x50, 0xc7, 0x86, 0x5a, 0xa4, 0xaa,
	0x7f, 0x98, 0x19, 0x41, 0x21, 0xb7, 0xae, 0x6a, 0x99, 0xfa, 0x9a, 0x9d, 0xd9, 0xa4, 0x76, 0x26,
	0x9f, 0xd7, 0xf0, 0x74, 0x5e, 0xff, 0x7c, 0xde, 0x4b, 0xee, 0xbb, 0x57, 0xdc, 0x37, 0x1b, 0xed,
	0x1d, 0x37, 0x1a, 0xbe, 0x86, 0xd1, 0x17, 0x61, 0xbe, 0x4b, 0x99, 0x6f, 0xab, 0x84, 0x8c, 0x31,
	0xf0, 0xca, 0x32, 0x4b, 0x9b, 0xad, 0xdb, 0x73, 0xb8, 0x82, 0xf1, 0x69, 0x92, 0x56, 0xec, 0x25,
	0x74, 0xb5, 0x89, 0x4d, 0xa9, 0x6d, 0x9e, 0xcf, 0x1b, 0x22, 0xeb, 0xba, 0x4a, 0x50, 0x28, 0xfa,
	0x44, 0x6e, 0x34, 0xe2, 0x2d, 0x86, 0x1f, 0x61, 0x42, 0x02, 0x9f, 0x50, 0xc4, 0x46, 0xfc, 0x50,
	0x49, 0x9c, 0xe7, 0x97, 0x7a, 0xdd, 0x50, 0x78, 0x03, 0x2f, 0x48, 0xe1, 0xb3, 0xd0, 0x06, 0xe5,
	0xe1, 0xba, 0x44, 0xf8, 0x01, 0xc6, 0x94, 0x48, 0xce, 0x9b, 0xac, 0x29, 0xf8, 0x3b, 0x94, 0xa5,
	0x6a, 0xd2, 0x6a, 0x20, 0x13, 0x85, 0x92, 0x32, 0xaf, 0x9f, 0x93, 0xc7, 0x1b, 0x0a, 0xff, 0x3b,
	0xc0, 0x1e, 0x0b, 0xdc, 0xf4, 0xbc, 0x04, 0xbf, 0x55, 0x71, 0xa3, 0xe1, 0xf2, 0xd5, 0x82, 0x5e,
	0xf6, 0x79, 0xbd, 0xbd, 0xe2, 0x75, 0xea, 0xec, 0x3d, 0x78, 0x84, 0xcf, 0xdb, 0xc0, 0xcf, 0xae,
	0xfd, 0x5f, 0xde, 0x3d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x79, 0xa1, 0x28, 0x81, 0x3c, 0x03, 0x00,
	0x00,
}
