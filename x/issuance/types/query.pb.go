// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issuance/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// this line is used by starport scaffolding # 3
type QueryGetTokenRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetTokenRequest) Reset()         { *m = QueryGetTokenRequest{} }
func (m *QueryGetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenRequest) ProtoMessage()    {}
func (*QueryGetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9428b94c0c172dcd, []int{0}
}
func (m *QueryGetTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenRequest.Merge(m, src)
}
func (m *QueryGetTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenRequest proto.InternalMessageInfo

func (m *QueryGetTokenRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetTokenResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *QueryGetTokenResponse) Reset()         { *m = QueryGetTokenResponse{} }
func (m *QueryGetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenResponse) ProtoMessage()    {}
func (*QueryGetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9428b94c0c172dcd, []int{1}
}
func (m *QueryGetTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenResponse.Merge(m, src)
}
func (m *QueryGetTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenResponse proto.InternalMessageInfo

func (m *QueryGetTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type QueryAllTokenRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenRequest) Reset()         { *m = QueryAllTokenRequest{} }
func (m *QueryAllTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenRequest) ProtoMessage()    {}
func (*QueryAllTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9428b94c0c172dcd, []int{2}
}
func (m *QueryAllTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenRequest.Merge(m, src)
}
func (m *QueryAllTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenRequest proto.InternalMessageInfo

func (m *QueryAllTokenRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTokenResponse struct {
	Token      []*Token            `protobuf:"bytes,1,rep,name=Token,proto3" json:"Token,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenResponse) Reset()         { *m = QueryAllTokenResponse{} }
func (m *QueryAllTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenResponse) ProtoMessage()    {}
func (*QueryAllTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9428b94c0c172dcd, []int{3}
}
func (m *QueryAllTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenResponse.Merge(m, src)
}
func (m *QueryAllTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenResponse proto.InternalMessageInfo

func (m *QueryAllTokenResponse) GetToken() []*Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *QueryAllTokenResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetTokenRequest)(nil), "AutonomyNetwork.autonomychain.issuance.v1.QueryGetTokenRequest")
	proto.RegisterType((*QueryGetTokenResponse)(nil), "AutonomyNetwork.autonomychain.issuance.v1.QueryGetTokenResponse")
	proto.RegisterType((*QueryAllTokenRequest)(nil), "AutonomyNetwork.autonomychain.issuance.v1.QueryAllTokenRequest")
	proto.RegisterType((*QueryAllTokenResponse)(nil), "AutonomyNetwork.autonomychain.issuance.v1.QueryAllTokenResponse")
}

func init() { proto.RegisterFile("issuance/v1/query.proto", fileDescriptor_9428b94c0c172dcd) }

var fileDescriptor_9428b94c0c172dcd = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0x87, 0xeb, 0x94, 0x22, 0x64, 0x24, 0x06, 0x0b, 0x04, 0xaa, 0x50, 0x84, 0x32, 0x94, 0x3f,
	0x12, 0x76, 0x53, 0x10, 0x43, 0x17, 0x28, 0x43, 0xbb, 0x21, 0xa8, 0x3a, 0x31, 0x80, 0x9c, 0xd4,
	0x4a, 0xad, 0xa6, 0x76, 0x5a, 0x3b, 0x85, 0x0a, 0xb1, 0xf0, 0x04, 0x48, 0x2c, 0x3c, 0x02, 0x2f,
	0x02, 0x62, 0xac, 0xc4, 0xc2, 0x88, 0x5a, 0x1e, 0x04, 0xd5, 0x36, 0xb4, 0x09, 0xf7, 0xea, 0xde,
	0xdc, 0x3b, 0xc6, 0x39, 0xe7, 0x77, 0xbe, 0x2f, 0x3e, 0x81, 0xd7, 0xb9, 0x52, 0x39, 0x15, 0x31,
	0x23, 0xcb, 0x90, 0xcc, 0x73, 0xb6, 0x58, 0xe1, 0x6c, 0x21, 0xb5, 0x44, 0x77, 0x7b, 0xb9, 0x96,
	0x42, 0xce, 0x56, 0xcf, 0x98, 0x7e, 0x23, 0x17, 0x53, 0x4c, 0xdd, 0x73, 0x3c, 0xa1, 0x5c, 0xe0,
	0xbf, 0x6d, 0x78, 0x19, 0x36, 0x6f, 0x26, 0x52, 0x26, 0x29, 0x23, 0x34, 0xe3, 0x84, 0x0a, 0x21,
	0x35, 0xd5, 0x5c, 0x0a, 0x65, 0x83, 0x9a, 0xf7, 0x62, 0xa9, 0x66, 0x52, 0x91, 0x88, 0x2a, 0x66,
	0x27, 0x90, 0x65, 0x18, 0x31, 0x4d, 0x43, 0x92, 0xd1, 0x84, 0x0b, 0x53, 0xec, 0x6a, 0x0b, 0x34,
	0x5a, 0x4e, 0x99, 0x7b, 0x11, 0xb4, 0xe0, 0xd5, 0x17, 0xbb, 0xd6, 0x01, 0xd3, 0xa3, 0xdd, 0xf1,
	0x90, 0xcd, 0x73, 0xa6, 0x34, 0xba, 0x02, 0x3d, 0x3e, 0xbe, 0x01, 0x6e, 0x81, 0x3b, 0x17, 0x86,
	0x1e, 0x1f, 0x07, 0xaf, 0xe1, 0xb5, 0x52, 0x9d, 0xca, 0xa4, 0x50, 0x0c, 0xf5, 0x61, 0xc3, 0x1c,
	0x98, 0xda, 0xcb, 0x9d, 0x36, 0x3e, 0xb5, 0x1e, 0xb6, 0x41, 0xb6, 0x3d, 0x78, 0xe5, 0x40, 0x7a,
	0x69, 0x5a, 0x00, 0xe9, 0x43, 0xb8, 0xb7, 0x71, 0x43, 0x5a, 0xd8, 0xaa, 0xe3, 0x9d, 0x3a, 0xb6,
	0x1f, 0xd7, 0xa9, 0xe3, 0xe7, 0x34, 0x61, 0xae, 0x77, 0x78, 0xd0, 0x19, 0x7c, 0x01, 0xce, 0x60,
	0x3f, 0xe0, 0x7f, 0x83, 0xfa, 0x39, 0x0c, 0xd0, 0xa0, 0x40, 0xea, 0x19, 0xd2, 0xdb, 0x27, 0x92,
	0x5a, 0x88, 0x43, 0xd4, 0xce, 0xe7, 0x3a, 0x6c, 0x18, 0x54, 0xf4, 0x0d, 0x38, 0x36, 0xf4, 0xb8,
	0x02, 0xd5, 0x51, 0x17, 0xda, 0x7c, 0x72, 0xf6, 0x00, 0x8b, 0x18, 0x74, 0x3f, 0xfc, 0xf8, 0xfd,
	0xc9, 0x7b, 0x88, 0x3a, 0xa4, 0x94, 0x44, 0x0a, 0x49, 0xe4, 0xdf, 0xaa, 0x99, 0x3d, 0x23, 0xef,
	0xf8, 0xf8, 0x3d, 0xfa, 0x0a, 0xe0, 0x25, 0x93, 0xd6, 0x4b, 0xd3, 0xea, 0x2e, 0xa5, 0x9d, 0xa8,
	0xee, 0x52, 0xbe, 0xf3, 0xe0, 0x91, 0x71, 0x69, 0x23, 0x5c, 0xc9, 0x45, 0x3d, 0x1d, 0x7d, 0xdf,
	0xf8, 0x60, 0xbd, 0xf1, 0xc1, 0xaf, 0x8d, 0x0f, 0x3e, 0x6e, 0xfd, 0xda, 0x7a, 0xeb, 0xd7, 0x7e,
	0x6e, 0xfd, 0xda, 0xcb, 0x6e, 0xc2, 0xf5, 0x24, 0x8f, 0x70, 0x2c, 0x67, 0xc7, 0x66, 0xde, 0xb7,
	0xa1, 0x6f, 0x0f, 0x62, 0x57, 0x19, 0x53, 0xd1, 0x45, 0xf3, 0x2f, 0x3e, 0xf8, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0xb5, 0x46, 0x73, 0x34, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a token by id.
	Token(ctx context.Context, in *QueryGetTokenRequest, opts ...grpc.CallOption) (*QueryGetTokenResponse, error)
	// Queries a list of token items.
	TokenAll(ctx context.Context, in *QueryAllTokenRequest, opts ...grpc.CallOption) (*QueryAllTokenResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Token(ctx context.Context, in *QueryGetTokenRequest, opts ...grpc.CallOption) (*QueryGetTokenResponse, error) {
	out := new(QueryGetTokenResponse)
	err := c.cc.Invoke(ctx, "/AutonomyNetwork.autonomychain.issuance.v1.Query/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TokenAll(ctx context.Context, in *QueryAllTokenRequest, opts ...grpc.CallOption) (*QueryAllTokenResponse, error) {
	out := new(QueryAllTokenResponse)
	err := c.cc.Invoke(ctx, "/AutonomyNetwork.autonomychain.issuance.v1.Query/TokenAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a token by id.
	Token(context.Context, *QueryGetTokenRequest) (*QueryGetTokenResponse, error)
	// Queries a list of token items.
	TokenAll(context.Context, *QueryAllTokenRequest) (*QueryAllTokenResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Token(ctx context.Context, req *QueryGetTokenRequest) (*QueryGetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (*UnimplementedQueryServer) TokenAll(ctx context.Context, req *QueryAllTokenRequest) (*QueryAllTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutonomyNetwork.autonomychain.issuance.v1.Query/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Token(ctx, req.(*QueryGetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TokenAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AutonomyNetwork.autonomychain.issuance.v1.Query/TokenAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenAll(ctx, req.(*QueryAllTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AutonomyNetwork.autonomychain.issuance.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _Query_Token_Handler,
		},
		{
			MethodName: "TokenAll",
			Handler:    _Query_TokenAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issuance/v1/query.proto",
}

func (m *QueryGetTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Token != nil {
		{
			size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token) > 0 {
		for iNdEx := len(m.Token) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Token[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Token != nil {
		l = m.Token.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Token) > 0 {
		for _, e := range m.Token {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetTokenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Token == nil {
				m.Token = &Token{}
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllTokenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token, &Token{})
			if err := m.Token[len(m.Token)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
