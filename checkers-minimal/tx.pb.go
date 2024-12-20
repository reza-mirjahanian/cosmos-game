// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alice/checkers/v1/tx.proto

package checkers

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// 2 - msg request type should be named as ReqCheckersTorram
type ReqCheckersTorram struct {
	// creator is the message sender.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Black   string `protobuf:"bytes,3,opt,name=black,proto3" json:"black,omitempty"`
	Red     string `protobuf:"bytes,4,opt,name=red,proto3" json:"red,omitempty"`
}

func (m *ReqCheckersTorram) Reset()         { *m = ReqCheckersTorram{} }
func (m *ReqCheckersTorram) String() string { return proto.CompactTextString(m) }
func (*ReqCheckersTorram) ProtoMessage()    {}
func (*ReqCheckersTorram) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e435dc85b9892e, []int{0}
}
func (m *ReqCheckersTorram) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReqCheckersTorram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReqCheckersTorram.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReqCheckersTorram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCheckersTorram.Merge(m, src)
}
func (m *ReqCheckersTorram) XXX_Size() int {
	return m.Size()
}
func (m *ReqCheckersTorram) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCheckersTorram.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCheckersTorram proto.InternalMessageInfo

func (m *ReqCheckersTorram) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ReqCheckersTorram) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ReqCheckersTorram) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *ReqCheckersTorram) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

// 3 - msg response type should be named as ResCheckersTorram
type ResCheckersTorram struct {
}

func (m *ResCheckersTorram) Reset()         { *m = ResCheckersTorram{} }
func (m *ResCheckersTorram) String() string { return proto.CompactTextString(m) }
func (*ResCheckersTorram) ProtoMessage()    {}
func (*ResCheckersTorram) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e435dc85b9892e, []int{1}
}
func (m *ResCheckersTorram) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResCheckersTorram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResCheckersTorram.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResCheckersTorram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResCheckersTorram.Merge(m, src)
}
func (m *ResCheckersTorram) XXX_Size() int {
	return m.Size()
}
func (m *ResCheckersTorram) XXX_DiscardUnknown() {
	xxx_messageInfo_ResCheckersTorram.DiscardUnknown(m)
}

var xxx_messageInfo_ResCheckersTorram proto.InternalMessageInfo

// 4 - In the keeper file, try to store more game data like game start time and end time.
type ResKillGame struct {
}

func (m *ResKillGame) Reset()         { *m = ResKillGame{} }
func (m *ResKillGame) String() string { return proto.CompactTextString(m) }
func (*ResKillGame) ProtoMessage()    {}
func (*ResKillGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e435dc85b9892e, []int{2}
}
func (m *ResKillGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResKillGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResKillGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResKillGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResKillGame.Merge(m, src)
}
func (m *ResKillGame) XXX_Size() int {
	return m.Size()
}
func (m *ResKillGame) XXX_DiscardUnknown() {
	xxx_messageInfo_ResKillGame.DiscardUnknown(m)
}

var xxx_messageInfo_ResKillGame proto.InternalMessageInfo

type ReqKillGame struct {
	Terminator string `protobuf:"bytes,1,opt,name=terminator,proto3" json:"terminator,omitempty"`
	Index      string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *ReqKillGame) Reset()         { *m = ReqKillGame{} }
func (m *ReqKillGame) String() string { return proto.CompactTextString(m) }
func (*ReqKillGame) ProtoMessage()    {}
func (*ReqKillGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e435dc85b9892e, []int{3}
}
func (m *ReqKillGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReqKillGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReqKillGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReqKillGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqKillGame.Merge(m, src)
}
func (m *ReqKillGame) XXX_Size() int {
	return m.Size()
}
func (m *ReqKillGame) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqKillGame.DiscardUnknown(m)
}

var xxx_messageInfo_ReqKillGame proto.InternalMessageInfo

func (m *ReqKillGame) GetTerminator() string {
	if m != nil {
		return m.Terminator
	}
	return ""
}

func (m *ReqKillGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqCheckersTorram)(nil), "alice.checkers.v1.ReqCheckersTorram")
	proto.RegisterType((*ResCheckersTorram)(nil), "alice.checkers.v1.ResCheckersTorram")
	proto.RegisterType((*ResKillGame)(nil), "alice.checkers.v1.ResKillGame")
	proto.RegisterType((*ReqKillGame)(nil), "alice.checkers.v1.ReqKillGame")
}

func init() { proto.RegisterFile("alice/checkers/v1/tx.proto", fileDescriptor_d8e435dc85b9892e) }

var fileDescriptor_d8e435dc85b9892e = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcc, 0xc9, 0x4c,
	0x4e, 0xd5, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0xcb, 0xe9, 0xc1, 0xe4, 0xf4, 0xca, 0x0c,
	0xa5, 0xc4, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0x41, 0x4a, 0x73, 0x8b,
	0xd3, 0x21, 0x6a, 0xa5, 0x24, 0x21, 0x12, 0xf1, 0x60, 0x9e, 0x3e, 0x84, 0x03, 0x91, 0x52, 0x5a,
	0xcb, 0xc8, 0x25, 0x18, 0x94, 0x5a, 0xe8, 0x0c, 0x35, 0x26, 0x24, 0xbf, 0xa8, 0x28, 0x31, 0x57,
	0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x02, 0x8b,
	0x43, 0x38, 0x42, 0x7a, 0x5c, 0xac, 0x49, 0x39, 0x89, 0xc9, 0xd9, 0x12, 0xcc, 0x20, 0x51, 0x27,
	0x89, 0x4b, 0x5b, 0x74, 0x45, 0xa0, 0xd6, 0x38, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x07, 0x97,
	0x14, 0x65, 0xe6, 0xa5, 0x07, 0x41, 0x94, 0x09, 0x69, 0x71, 0x31, 0x17, 0xa5, 0xa6, 0x48, 0xb0,
	0x10, 0x50, 0x0d, 0x52, 0x64, 0xc5, 0xd3, 0xf4, 0x7c, 0x83, 0x16, 0xcc, 0x7e, 0x25, 0x61, 0x90,
	0x73, 0x8b, 0x51, 0x9d, 0xab, 0xc4, 0xcb, 0xc5, 0x1d, 0x94, 0x5a, 0xec, 0x9d, 0x99, 0x93, 0xe3,
	0x9e, 0x98, 0x9b, 0xaa, 0x14, 0x02, 0xe2, 0x16, 0xc2, 0xb8, 0x42, 0x72, 0x5c, 0x5c, 0x25, 0xa9,
	0x45, 0xb9, 0x99, 0x79, 0x48, 0xfe, 0x41, 0x12, 0xc1, 0xee, 0x25, 0x2b, 0x7e, 0x90, 0xb5, 0x48,
	0xca, 0x8c, 0x0e, 0x33, 0x72, 0xf1, 0xa1, 0x05, 0x53, 0x1c, 0x97, 0x00, 0x4c, 0xc4, 0x19, 0xe4,
	0xbe, 0x54, 0xf7, 0x5c, 0x21, 0x15, 0x3d, 0x8c, 0x88, 0xd1, 0xc3, 0x08, 0x60, 0x29, 0xec, 0xaa,
	0xd0, 0xfc, 0x25, 0xe4, 0xc5, 0xc5, 0x81, 0xf0, 0x05, 0x76, 0x73, 0x61, 0xf2, 0x52, 0xd8, 0xe5,
	0xe1, 0x81, 0x22, 0xc5, 0xda, 0xf0, 0x7c, 0x83, 0x16, 0xa3, 0x93, 0xf1, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x49, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xa3, 0xa6, 0xbb, 0x24, 0x36, 0x70, 0x5a, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x86, 0x1e, 0xdb, 0x7c, 0x90, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckersTorramClient is the client API for CheckersTorram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckersTorramClient interface {
	//  2 - rpc should be named as CheckersCreateGm
	CheckersCreateGm(ctx context.Context, in *ReqCheckersTorram, opts ...grpc.CallOption) (*ResCheckersTorram, error)
	KillGame(ctx context.Context, in *ReqKillGame, opts ...grpc.CallOption) (*ResKillGame, error)
}

type checkersTorramClient struct {
	cc grpc1.ClientConn
}

func NewCheckersTorramClient(cc grpc1.ClientConn) CheckersTorramClient {
	return &checkersTorramClient{cc}
}

func (c *checkersTorramClient) CheckersCreateGm(ctx context.Context, in *ReqCheckersTorram, opts ...grpc.CallOption) (*ResCheckersTorram, error) {
	out := new(ResCheckersTorram)
	err := c.cc.Invoke(ctx, "/alice.checkers.v1.CheckersTorram/CheckersCreateGm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkersTorramClient) KillGame(ctx context.Context, in *ReqKillGame, opts ...grpc.CallOption) (*ResKillGame, error) {
	out := new(ResKillGame)
	err := c.cc.Invoke(ctx, "/alice.checkers.v1.CheckersTorram/KillGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckersTorramServer is the server API for CheckersTorram service.
type CheckersTorramServer interface {
	//  2 - rpc should be named as CheckersCreateGm
	CheckersCreateGm(context.Context, *ReqCheckersTorram) (*ResCheckersTorram, error)
	KillGame(context.Context, *ReqKillGame) (*ResKillGame, error)
}

// UnimplementedCheckersTorramServer can be embedded to have forward compatible implementations.
type UnimplementedCheckersTorramServer struct {
}

func (*UnimplementedCheckersTorramServer) CheckersCreateGm(ctx context.Context, req *ReqCheckersTorram) (*ResCheckersTorram, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckersCreateGm not implemented")
}
func (*UnimplementedCheckersTorramServer) KillGame(ctx context.Context, req *ReqKillGame) (*ResKillGame, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillGame not implemented")
}

func RegisterCheckersTorramServer(s grpc1.Server, srv CheckersTorramServer) {
	s.RegisterService(&_CheckersTorram_serviceDesc, srv)
}

func _CheckersTorram_CheckersCreateGm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCheckersTorram)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckersTorramServer).CheckersCreateGm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.checkers.v1.CheckersTorram/CheckersCreateGm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckersTorramServer).CheckersCreateGm(ctx, req.(*ReqCheckersTorram))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckersTorram_KillGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqKillGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckersTorramServer).KillGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.checkers.v1.CheckersTorram/KillGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckersTorramServer).KillGame(ctx, req.(*ReqKillGame))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckersTorram_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alice.checkers.v1.CheckersTorram",
	HandlerType: (*CheckersTorramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckersCreateGm",
			Handler:    _CheckersTorram_CheckersCreateGm_Handler,
		},
		{
			MethodName: "KillGame",
			Handler:    _CheckersTorram_KillGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alice/checkers/v1/tx.proto",
}

func (m *ReqCheckersTorram) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqCheckersTorram) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReqCheckersTorram) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResCheckersTorram) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResCheckersTorram) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResCheckersTorram) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ResKillGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResKillGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResKillGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ReqKillGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqKillGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReqKillGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Terminator) > 0 {
		i -= len(m.Terminator)
		copy(dAtA[i:], m.Terminator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Terminator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReqCheckersTorram) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ResCheckersTorram) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ResKillGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReqKillGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Terminator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReqCheckersTorram) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReqCheckersTorram: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqCheckersTorram: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResCheckersTorram) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResCheckersTorram: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResCheckersTorram: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResKillGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResKillGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResKillGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReqKillGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReqKillGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqKillGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Terminator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
