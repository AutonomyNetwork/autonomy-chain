// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launchpad/v1beta1/events.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
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

type EventModule struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *EventModule) Reset()         { *m = EventModule{} }
func (m *EventModule) String() string { return proto.CompactTextString(m) }
func (*EventModule) ProtoMessage()    {}
func (*EventModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_57175d03a91888f4, []int{0}
}
func (m *EventModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventModule.Merge(m, src)
}
func (m *EventModule) XXX_Size() int {
	return m.Size()
}
func (m *EventModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EventModule.DiscardUnknown(m)
}

var xxx_messageInfo_EventModule proto.InternalMessageInfo

func (m *EventModule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EventCreateLaunchpad struct {
	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator         string    `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	TokenId         uint64    `protobuf:"varint,3,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	AccepetedDenoms []string  `protobuf:"bytes,4,rep,name=accepetedDenoms,proto3" json:"accepetedDenoms,omitempty"`
	Softcap         uint64    `protobuf:"varint,5,opt,name=softcap,proto3" json:"softcap,omitempty"`
	Hardcap         uint64    `protobuf:"varint,6,opt,name=hardcap,proto3" json:"hardcap,omitempty"`
	StartTime       time.Time `protobuf:"bytes,7,opt,name=startTime,proto3,stdtime" json:"startTime"`
	EndTime         time.Time `protobuf:"bytes,8,opt,name=endTime,proto3,stdtime" json:"endTime"`
	Status          string    `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *EventCreateLaunchpad) Reset()         { *m = EventCreateLaunchpad{} }
func (m *EventCreateLaunchpad) String() string { return proto.CompactTextString(m) }
func (*EventCreateLaunchpad) ProtoMessage()    {}
func (*EventCreateLaunchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_57175d03a91888f4, []int{1}
}
func (m *EventCreateLaunchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateLaunchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateLaunchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateLaunchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateLaunchpad.Merge(m, src)
}
func (m *EventCreateLaunchpad) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateLaunchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateLaunchpad.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateLaunchpad proto.InternalMessageInfo

type EventDepositToLaunchpad struct {
	Id        uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *EventDepositToLaunchpad) Reset()         { *m = EventDepositToLaunchpad{} }
func (m *EventDepositToLaunchpad) String() string { return proto.CompactTextString(m) }
func (*EventDepositToLaunchpad) ProtoMessage()    {}
func (*EventDepositToLaunchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_57175d03a91888f4, []int{2}
}
func (m *EventDepositToLaunchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDepositToLaunchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDepositToLaunchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDepositToLaunchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDepositToLaunchpad.Merge(m, src)
}
func (m *EventDepositToLaunchpad) XXX_Size() int {
	return m.Size()
}
func (m *EventDepositToLaunchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDepositToLaunchpad.DiscardUnknown(m)
}

var xxx_messageInfo_EventDepositToLaunchpad proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventModule)(nil), "launchpad.v1beta1.EventModule")
	proto.RegisterType((*EventCreateLaunchpad)(nil), "launchpad.v1beta1.EventCreateLaunchpad")
	proto.RegisterType((*EventDepositToLaunchpad)(nil), "launchpad.v1beta1.EventDepositToLaunchpad")
}

func init() { proto.RegisterFile("launchpad/v1beta1/events.proto", fileDescriptor_57175d03a91888f4) }

var fileDescriptor_57175d03a91888f4 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x93, 0xb6, 0xf4, 0x8f, 0x2b, 0x81, 0x88, 0x26, 0xc8, 0x2a, 0x94, 0x94, 0x9e, 0x72,
	0x59, 0xc2, 0xc6, 0x0d, 0x24, 0x24, 0xba, 0x71, 0x40, 0x02, 0x0e, 0x55, 0xb9, 0x70, 0x41, 0xae,
	0xfd, 0xae, 0x8d, 0xda, 0xf8, 0x8d, 0x62, 0x67, 0xb0, 0x6f, 0xc0, 0x71, 0x1f, 0x61, 0x67, 0x3e,
	0x00, 0x9f, 0x61, 0xc7, 0x1d, 0x39, 0x6d, 0xa8, 0xbd, 0xf0, 0x31, 0x90, 0x1d, 0xa7, 0x45, 0xc0,
	0x85, 0x53, 0xfc, 0xf8, 0xf7, 0xf8, 0xb1, 0xf3, 0xc4, 0x21, 0xc1, 0x8a, 0x96, 0x82, 0x2d, 0x72,
	0xca, 0x93, 0xb3, 0xc3, 0x19, 0x28, 0x7a, 0x98, 0xc0, 0x19, 0x08, 0x25, 0xe3, 0xbc, 0x40, 0x85,
	0xde, 0xfd, 0x2d, 0x8f, 0x2d, 0x1f, 0xec, 0xcd, 0x71, 0x8e, 0x86, 0x26, 0x7a, 0x54, 0x19, 0x07,
	0xe1, 0x1c, 0x71, 0xbe, 0x82, 0xc4, 0xa8, 0x59, 0x79, 0x9a, 0xa8, 0x34, 0x03, 0xa9, 0x68, 0x96,
	0x5b, 0x43, 0xc0, 0x50, 0x66, 0x28, 0x93, 0x19, 0x95, 0xb0, 0xdd, 0x8b, 0x61, 0x2a, 0xfe, 0xe2,
	0x62, 0xb9, 0xe5, 0x5a, 0x58, 0xbe, 0x5f, 0xf1, 0x8f, 0xd5, 0xce, 0x95, 0xa8, 0xd0, 0xe8, 0x31,
	0xe9, 0xbf, 0xd2, 0x87, 0x7e, 0x8b, 0xbc, 0x5c, 0x81, 0xe7, 0x91, 0x96, 0xa0, 0x19, 0xf8, 0xee,
	0xd0, 0x8d, 0x7a, 0x13, 0x33, 0x1e, 0xdd, 0x34, 0xc8, 0x9e, 0xf1, 0x1c, 0x17, 0x40, 0x15, 0xbc,
	0xa9, 0xdf, 0xca, 0xbb, 0x4b, 0x1a, 0x29, 0xb7, 0xd6, 0x46, 0xca, 0x3d, 0x9f, 0x74, 0x98, 0xb6,
	0x60, 0xe1, 0x37, 0xcc, 0x64, 0x2d, 0x35, 0x51, 0xb8, 0x04, 0xf1, 0x9a, 0xfb, 0xcd, 0xa1, 0x1b,
	0xb5, 0x26, 0xb5, 0xf4, 0x22, 0x72, 0x8f, 0x32, 0x06, 0x39, 0x28, 0xe0, 0x27, 0x20, 0x30, 0x93,
	0x7e, 0x6b, 0xd8, 0x8c, 0x7a, 0x93, 0x3f, 0xa7, 0x75, 0x86, 0xc4, 0x53, 0xc5, 0x68, 0xee, 0xdf,
	0xa9, 0x32, 0xac, 0xd4, 0x64, 0x41, 0x0b, 0xae, 0x49, 0xbb, 0x22, 0x56, 0x7a, 0x63, 0xd2, 0x93,
	0x8a, 0x16, 0x6a, 0x9a, 0x66, 0xe0, 0x77, 0x86, 0x6e, 0xd4, 0x3f, 0x1a, 0xc4, 0x55, 0xdb, 0x71,
	0xdd, 0x76, 0x3c, 0xad, 0xdb, 0x1e, 0x77, 0xaf, 0x6e, 0x42, 0xe7, 0xe2, 0x36, 0x74, 0x27, 0xbb,
	0x65, 0xde, 0x0b, 0xd2, 0x01, 0xc1, 0x4d, 0x42, 0xf7, 0x3f, 0x12, 0xea, 0x45, 0xde, 0x03, 0xd2,
	0x96, 0x8a, 0xaa, 0x52, 0xfa, 0x3d, 0x53, 0x8a, 0x55, 0xcf, 0xba, 0x5f, 0x2e, 0x43, 0xe7, 0xe7,
	0x65, 0xe8, 0x8c, 0xbe, 0xb9, 0xe4, 0xa1, 0x29, 0xf8, 0x04, 0x72, 0x94, 0xa9, 0x9a, 0xe2, 0xbf,
	0x3a, 0x6e, 0x99, 0x8e, 0x1f, 0x91, 0x1e, 0xaf, 0x5c, 0xdb, 0x96, 0x77, 0x13, 0x1e, 0x23, 0x6d,
	0x9a, 0x61, 0x29, 0x94, 0xdf, 0x1c, 0x36, 0xa3, 0xfe, 0xd1, 0x7e, 0x6c, 0x3f, 0xb6, 0xbe, 0x39,
	0xf5, 0x2d, 0x8c, 0x8f, 0x31, 0x15, 0xe3, 0x27, 0xfa, 0xa4, 0x5f, 0x6f, 0xc3, 0x68, 0x9e, 0xaa,
	0x45, 0x39, 0x8b, 0x19, 0x66, 0xf6, 0x66, 0xd8, 0xc7, 0x81, 0xe4, 0xcb, 0x44, 0x9d, 0xe7, 0x20,
	0xcd, 0x02, 0x39, 0xb1, 0xd1, 0xbb, 0x83, 0x8f, 0xdf, 0x5f, 0xad, 0x03, 0xf7, 0x7a, 0x1d, 0xb8,
	0x3f, 0xd6, 0x81, 0x7b, 0xb1, 0x09, 0x9c, 0xeb, 0x4d, 0xe0, 0x7c, 0xdf, 0x04, 0xce, 0x87, 0xe7,
	0xbf, 0xa5, 0xbe, 0x2c, 0x15, 0x0a, 0xcc, 0xce, 0xdf, 0x81, 0xfa, 0x84, 0xc5, 0x32, 0xa1, 0x56,
	0x1f, 0xb0, 0x05, 0x4d, 0x45, 0xf2, 0x39, 0xd9, 0xfd, 0x47, 0x66, 0xbb, 0x59, 0xdb, 0x14, 0xfb,
	0xf4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x44, 0xe7, 0xe0, 0x61, 0x03, 0x00, 0x00,
}

func (m *EventModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateLaunchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateLaunchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateLaunchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x4a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvents(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEvents(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	if m.Hardcap != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Hardcap))
		i--
		dAtA[i] = 0x30
	}
	if m.Softcap != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Softcap))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AccepetedDenoms) > 0 {
		for iNdEx := len(m.AccepetedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AccepetedDenoms[iNdEx])
			copy(dAtA[i:], m.AccepetedDenoms[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.AccepetedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.TokenId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDepositToLaunchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDepositToLaunchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDepositToLaunchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCreateLaunchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovEvents(uint64(m.TokenId))
	}
	if len(m.AccepetedDenoms) > 0 {
		for _, s := range m.AccepetedDenoms {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.Softcap != 0 {
		n += 1 + sovEvents(uint64(m.Softcap))
	}
	if m.Hardcap != 0 {
		n += 1 + sovEvents(uint64(m.Hardcap))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovEvents(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovEvents(uint64(l))
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventDepositToLaunchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateLaunchpad) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateLaunchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateLaunchpad: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccepetedDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccepetedDenoms = append(m.AccepetedDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Softcap", wireType)
			}
			m.Softcap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hardcap", wireType)
			}
			m.Hardcap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventDepositToLaunchpad) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDepositToLaunchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDepositToLaunchpad: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
