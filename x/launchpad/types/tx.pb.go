// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launchpad/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # proto/tx/message
type MsgCreateLaunchpad struct {
	Creator         string    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TokenId         uint64    `protobuf:"varint,2,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	AccepetedDenoms []string  `protobuf:"bytes,3,rep,name=accepetedDenoms,proto3" json:"accepetedDenoms,omitempty"`
	Softcap         uint64    `protobuf:"varint,4,opt,name=softcap,proto3" json:"softcap,omitempty"`
	Hardcap         uint64    `protobuf:"varint,5,opt,name=hardcap,proto3" json:"hardcap,omitempty"`
	StartTime       time.Time `protobuf:"bytes,6,opt,name=startTime,proto3,stdtime" json:"startTime"`
	EndTime         time.Time `protobuf:"bytes,7,opt,name=endTime,proto3,stdtime" json:"endTime"`
	Status          string    `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *MsgCreateLaunchpad) Reset()         { *m = MsgCreateLaunchpad{} }
func (m *MsgCreateLaunchpad) String() string { return proto.CompactTextString(m) }
func (*MsgCreateLaunchpad) ProtoMessage()    {}
func (*MsgCreateLaunchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3007a6ba7585d16, []int{0}
}
func (m *MsgCreateLaunchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateLaunchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateLaunchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateLaunchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateLaunchpad.Merge(m, src)
}
func (m *MsgCreateLaunchpad) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateLaunchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateLaunchpad.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateLaunchpad proto.InternalMessageInfo

type MsgCreateLaunchpadResponse struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgCreateLaunchpadResponse) Reset()         { *m = MsgCreateLaunchpadResponse{} }
func (m *MsgCreateLaunchpadResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateLaunchpadResponse) ProtoMessage()    {}
func (*MsgCreateLaunchpadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3007a6ba7585d16, []int{1}
}
func (m *MsgCreateLaunchpadResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateLaunchpadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateLaunchpadResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateLaunchpadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateLaunchpadResponse.Merge(m, src)
}
func (m *MsgCreateLaunchpadResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateLaunchpadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateLaunchpadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateLaunchpadResponse proto.InternalMessageInfo

func (m *MsgCreateLaunchpadResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgCreateLaunchpadResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgDepositToLaunchpad struct {
	Id        uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MsgDepositToLaunchpad) Reset()         { *m = MsgDepositToLaunchpad{} }
func (m *MsgDepositToLaunchpad) String() string { return proto.CompactTextString(m) }
func (*MsgDepositToLaunchpad) ProtoMessage()    {}
func (*MsgDepositToLaunchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3007a6ba7585d16, []int{2}
}
func (m *MsgDepositToLaunchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositToLaunchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositToLaunchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositToLaunchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositToLaunchpad.Merge(m, src)
}
func (m *MsgDepositToLaunchpad) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositToLaunchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositToLaunchpad.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositToLaunchpad proto.InternalMessageInfo

type MsgDepositToLaunchpadResponse struct {
	Id        uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MsgDepositToLaunchpadResponse) Reset()         { *m = MsgDepositToLaunchpadResponse{} }
func (m *MsgDepositToLaunchpadResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositToLaunchpadResponse) ProtoMessage()    {}
func (*MsgDepositToLaunchpadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3007a6ba7585d16, []int{3}
}
func (m *MsgDepositToLaunchpadResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositToLaunchpadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositToLaunchpadResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositToLaunchpadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositToLaunchpadResponse.Merge(m, src)
}
func (m *MsgDepositToLaunchpadResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositToLaunchpadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositToLaunchpadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositToLaunchpadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateLaunchpad)(nil), "launchpad.v1beta1.MsgCreateLaunchpad")
	proto.RegisterType((*MsgCreateLaunchpadResponse)(nil), "launchpad.v1beta1.MsgCreateLaunchpadResponse")
	proto.RegisterType((*MsgDepositToLaunchpad)(nil), "launchpad.v1beta1.MsgDepositToLaunchpad")
	proto.RegisterType((*MsgDepositToLaunchpadResponse)(nil), "launchpad.v1beta1.MsgDepositToLaunchpadResponse")
}

func init() { proto.RegisterFile("launchpad/v1beta1/tx.proto", fileDescriptor_d3007a6ba7585d16) }

var fileDescriptor_d3007a6ba7585d16 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x7d, 0x69, 0x7e, 0xf9, 0x73, 0x95, 0x7e, 0x15, 0x27, 0x40, 0xae, 0x05, 0x76, 0x14,
	0x09, 0xc9, 0x4b, 0xec, 0x36, 0x6c, 0x20, 0x21, 0x91, 0x56, 0x48, 0x48, 0x84, 0xc1, 0x0a, 0x0b,
	0x0b, 0xba, 0xd8, 0x57, 0xc7, 0x4a, 0x7d, 0x8f, 0xe5, 0x3b, 0x43, 0xfb, 0x0e, 0x18, 0xfb, 0x12,
	0x3a, 0x33, 0x33, 0xf2, 0x02, 0x3a, 0x30, 0x74, 0x64, 0xa2, 0x55, 0xb2, 0xf0, 0x32, 0x90, 0xed,
	0x73, 0xd2, 0x26, 0x41, 0x2a, 0x1b, 0x53, 0xf2, 0xbd, 0xef, 0x73, 0xdf, 0xbb, 0xfb, 0x3c, 0x77,
	0xc6, 0xc6, 0x31, 0xcd, 0xb8, 0x3f, 0x49, 0x68, 0xe0, 0x7e, 0xdc, 0x1f, 0x33, 0x49, 0xf7, 0x5d,
	0x79, 0xe2, 0x24, 0x29, 0x48, 0x20, 0xf7, 0x16, 0x9e, 0xa3, 0x3c, 0xe3, 0x7e, 0x08, 0x21, 0x14,
	0xae, 0x9b, 0xff, 0x2b, 0x0b, 0x0d, 0x2b, 0x04, 0x08, 0x8f, 0x99, 0x5b, 0xa8, 0x71, 0x76, 0xe4,
	0xca, 0x28, 0x66, 0x42, 0xd2, 0x38, 0x51, 0x05, 0xa6, 0x0f, 0x22, 0x06, 0xe1, 0x8e, 0xa9, 0x60,
	0x8b, 0x75, 0x7c, 0x88, 0xf8, 0x9a, 0xcf, 0xa7, 0x0b, 0x3f, 0x17, 0xca, 0xdf, 0x2d, 0xfd, 0x0f,
	0xe5, 0xca, 0xa5, 0x28, 0xad, 0xee, 0xf7, 0x1a, 0x26, 0x43, 0x11, 0x1e, 0xa4, 0x8c, 0x4a, 0xf6,
	0xa6, 0xda, 0x30, 0xd1, 0x71, 0xd3, 0xcf, 0x87, 0x20, 0xd5, 0x51, 0x07, 0xd9, 0x6d, 0xaf, 0x92,
	0xb9, 0x23, 0x61, 0xca, 0xf8, 0xeb, 0x40, 0xaf, 0x75, 0x90, 0x5d, 0xf7, 0x2a, 0x49, 0x6c, 0xbc,
	0x43, 0x7d, 0x9f, 0x25, 0x4c, 0xb2, 0xe0, 0x90, 0x71, 0x88, 0x85, 0xbe, 0xd5, 0xd9, 0xb2, 0xdb,
	0xde, 0xea, 0x70, 0x9e, 0x21, 0xe0, 0x48, 0xfa, 0x34, 0xd1, 0xeb, 0x65, 0x86, 0x92, 0xb9, 0x33,
	0xa1, 0x69, 0x90, 0x3b, 0xff, 0x95, 0x8e, 0x92, 0x64, 0x80, 0xdb, 0x42, 0xd2, 0x54, 0x8e, 0xa2,
	0x98, 0xe9, 0x8d, 0x0e, 0xb2, 0xb7, 0xfb, 0x86, 0x53, 0x82, 0x73, 0x2a, 0x70, 0xce, 0xa8, 0x02,
	0x37, 0x68, 0x5d, 0xfc, 0xb4, 0xb4, 0xb3, 0x2b, 0x0b, 0x79, 0xcb, 0x69, 0xe4, 0x05, 0x6e, 0x32,
	0x1e, 0x14, 0x09, 0xcd, 0xbf, 0x48, 0xa8, 0x26, 0x91, 0x87, 0xb8, 0x21, 0x24, 0x95, 0x99, 0xd0,
	0x5b, 0x05, 0x14, 0xa5, 0x9e, 0xb5, 0x3e, 0x9f, 0x5b, 0xda, 0xaf, 0x73, 0x4b, 0xeb, 0xbe, 0xc2,
	0xc6, 0x3a, 0x4d, 0x8f, 0x89, 0x04, 0xb8, 0x60, 0xe4, 0x7f, 0x5c, 0x8b, 0x82, 0x02, 0x68, 0xdd,
	0xab, 0x45, 0xb7, 0x28, 0xd7, 0x6e, 0x51, 0xee, 0x7e, 0x45, 0xf8, 0xc1, 0x50, 0x84, 0x87, 0x2c,
	0x01, 0x11, 0xc9, 0x11, 0x2c, 0x3b, 0xb3, 0x9a, 0xf1, 0x08, 0xb7, 0x83, 0xb2, 0x6a, 0x91, 0xb2,
	0x1c, 0x20, 0x3e, 0x6e, 0xd0, 0x18, 0x32, 0x2e, 0x8b, 0x56, 0x6c, 0xf7, 0x77, 0x1d, 0xd5, 0xfd,
	0xfc, 0x2a, 0x55, 0xd7, 0xd2, 0x39, 0x80, 0x88, 0x0f, 0xf6, 0xf2, 0xf3, 0x7e, 0xb9, 0xb2, 0xec,
	0x30, 0x92, 0x93, 0x6c, 0xec, 0xf8, 0x10, 0xab, 0xab, 0xa2, 0x7e, 0x7a, 0x22, 0x98, 0xba, 0xf2,
	0x34, 0x61, 0xa2, 0x98, 0x20, 0x3c, 0x15, 0x7d, 0xe3, 0xf8, 0xdf, 0x10, 0x7e, 0xbc, 0x71, 0xdb,
	0x7f, 0x44, 0xf0, 0x2f, 0x6d, 0xbf, 0x7f, 0x8d, 0xf0, 0xd6, 0x50, 0x84, 0x24, 0xc4, 0x3b, 0xab,
	0x0f, 0xe2, 0x89, 0xb3, 0xf6, 0x9a, 0x9d, 0xf5, 0x4e, 0x1b, 0xbd, 0x3b, 0x95, 0x2d, 0x68, 0x24,
	0x98, 0x6c, 0x68, 0xb1, 0xbd, 0x39, 0x64, 0xbd, 0xd2, 0xd8, 0xbb, 0x6b, 0x65, 0xb5, 0xe2, 0xe0,
	0xdd, 0xc5, 0xcc, 0x44, 0x97, 0x33, 0x13, 0x5d, 0xcf, 0x4c, 0x74, 0x36, 0x37, 0xb5, 0xcb, 0xb9,
	0xa9, 0xfd, 0x98, 0x9b, 0xda, 0xfb, 0xe7, 0x37, 0xc0, 0xbd, 0xcc, 0x24, 0x70, 0x88, 0x4f, 0xdf,
	0x32, 0xf9, 0x09, 0xd2, 0xa9, 0x4b, 0x95, 0xee, 0xf9, 0x13, 0x1a, 0x71, 0xf7, 0xc4, 0x5d, 0x7e,
	0xf6, 0x0a, 0xa2, 0xe3, 0x46, 0xf1, 0x80, 0x9e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xb6,
	0xbe, 0x8d, 0x10, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateLaunchpad(ctx context.Context, in *MsgCreateLaunchpad, opts ...grpc.CallOption) (*MsgCreateLaunchpadResponse, error)
	DepositToLaunchpad(ctx context.Context, in *MsgDepositToLaunchpad, opts ...grpc.CallOption) (*MsgDepositToLaunchpadResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateLaunchpad(ctx context.Context, in *MsgCreateLaunchpad, opts ...grpc.CallOption) (*MsgCreateLaunchpadResponse, error) {
	out := new(MsgCreateLaunchpadResponse)
	err := c.cc.Invoke(ctx, "/launchpad.v1beta1.Msg/CreateLaunchpad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepositToLaunchpad(ctx context.Context, in *MsgDepositToLaunchpad, opts ...grpc.CallOption) (*MsgDepositToLaunchpadResponse, error) {
	out := new(MsgDepositToLaunchpadResponse)
	err := c.cc.Invoke(ctx, "/launchpad.v1beta1.Msg/DepositToLaunchpad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateLaunchpad(context.Context, *MsgCreateLaunchpad) (*MsgCreateLaunchpadResponse, error)
	DepositToLaunchpad(context.Context, *MsgDepositToLaunchpad) (*MsgDepositToLaunchpadResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateLaunchpad(ctx context.Context, req *MsgCreateLaunchpad) (*MsgCreateLaunchpadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaunchpad not implemented")
}
func (*UnimplementedMsgServer) DepositToLaunchpad(ctx context.Context, req *MsgDepositToLaunchpad) (*MsgDepositToLaunchpadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositToLaunchpad not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateLaunchpad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLaunchpad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLaunchpad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/launchpad.v1beta1.Msg/CreateLaunchpad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLaunchpad(ctx, req.(*MsgCreateLaunchpad))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepositToLaunchpad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositToLaunchpad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositToLaunchpad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/launchpad.v1beta1.Msg/DepositToLaunchpad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositToLaunchpad(ctx, req.(*MsgDepositToLaunchpad))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "launchpad.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaunchpad",
			Handler:    _Msg_CreateLaunchpad_Handler,
		},
		{
			MethodName: "DepositToLaunchpad",
			Handler:    _Msg_DepositToLaunchpad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "launchpad/v1beta1/tx.proto",
}

func (m *MsgCreateLaunchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateLaunchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateLaunchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x42
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTx(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if m.Hardcap != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Hardcap))
		i--
		dAtA[i] = 0x28
	}
	if m.Softcap != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Softcap))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AccepetedDenoms) > 0 {
		for iNdEx := len(m.AccepetedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AccepetedDenoms[iNdEx])
			copy(dAtA[i:], m.AccepetedDenoms[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.AccepetedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.TokenId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x10
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

func (m *MsgCreateLaunchpadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateLaunchpadResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateLaunchpadResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositToLaunchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositToLaunchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositToLaunchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositToLaunchpadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositToLaunchpadResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositToLaunchpadResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
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
func (m *MsgCreateLaunchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovTx(uint64(m.TokenId))
	}
	if len(m.AccepetedDenoms) > 0 {
		for _, s := range m.AccepetedDenoms {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Softcap != 0 {
		n += 1 + sovTx(uint64(m.Softcap))
	}
	if m.Hardcap != 0 {
		n += 1 + sovTx(uint64(m.Hardcap))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateLaunchpadResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDepositToLaunchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgDepositToLaunchpadResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateLaunchpad) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateLaunchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateLaunchpad: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccepetedDenoms", wireType)
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
			m.AccepetedDenoms = append(m.AccepetedDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Softcap", wireType)
			}
			m.Softcap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Softcap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hardcap", wireType)
			}
			m.Hardcap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hardcap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateLaunchpadResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateLaunchpadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateLaunchpadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
func (m *MsgDepositToLaunchpad) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositToLaunchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositToLaunchpad: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgDepositToLaunchpadResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositToLaunchpadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositToLaunchpadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
