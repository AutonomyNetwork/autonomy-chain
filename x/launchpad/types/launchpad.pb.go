// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launchpad/v1beta1/launchpad.proto

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

type Launchpad struct {
	Id              uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator         string    `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	TokenId         uint64    `protobuf:"varint,3,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	AccepetedDenoms []string  `protobuf:"bytes,4,rep,name=accepetedDenoms,proto3" json:"accepetedDenoms,omitempty"`
	Softcap         uint64    `protobuf:"varint,5,opt,name=softcap,proto3" json:"softcap,omitempty"`
	Hardcap         uint64    `protobuf:"varint,6,opt,name=hardcap,proto3" json:"hardcap,omitempty"`
	StartTime       time.Time `protobuf:"bytes,7,opt,name=startTime,proto3,stdtime" json:"startTime"`
	EndTime         time.Time `protobuf:"bytes,8,opt,name=endTime,proto3,stdtime" json:"endTime"`
	Status          string    `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Launchpad) Reset()         { *m = Launchpad{} }
func (m *Launchpad) String() string { return proto.CompactTextString(m) }
func (*Launchpad) ProtoMessage()    {}
func (*Launchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8919382d78c3e36, []int{0}
}
func (m *Launchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Launchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Launchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Launchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Launchpad.Merge(m, src)
}
func (m *Launchpad) XXX_Size() int {
	return m.Size()
}
func (m *Launchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_Launchpad.DiscardUnknown(m)
}

var xxx_messageInfo_Launchpad proto.InternalMessageInfo

type DepositToLaunchpad struct {
	Id        uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *DepositToLaunchpad) Reset()         { *m = DepositToLaunchpad{} }
func (m *DepositToLaunchpad) String() string { return proto.CompactTextString(m) }
func (*DepositToLaunchpad) ProtoMessage()    {}
func (*DepositToLaunchpad) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8919382d78c3e36, []int{1}
}
func (m *DepositToLaunchpad) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositToLaunchpad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositToLaunchpad.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositToLaunchpad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositToLaunchpad.Merge(m, src)
}
func (m *DepositToLaunchpad) XXX_Size() int {
	return m.Size()
}
func (m *DepositToLaunchpad) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositToLaunchpad.DiscardUnknown(m)
}

var xxx_messageInfo_DepositToLaunchpad proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Launchpad)(nil), "launchpad.v1beta1.Launchpad")
	proto.RegisterType((*DepositToLaunchpad)(nil), "launchpad.v1beta1.DepositToLaunchpad")
}

func init() { proto.RegisterFile("launchpad/v1beta1/launchpad.proto", fileDescriptor_a8919382d78c3e36) }

var fileDescriptor_a8919382d78c3e36 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x6f, 0xd4, 0x40,
	0x10, 0xb5, 0xef, 0xc2, 0xdd, 0x79, 0x23, 0x81, 0xb0, 0x10, 0x72, 0x4e, 0xc8, 0x3e, 0x52, 0xb9,
	0x89, 0x4d, 0x42, 0x07, 0x12, 0x12, 0x47, 0x1a, 0x24, 0x44, 0x61, 0x1d, 0x0d, 0x0d, 0x5a, 0xaf,
	0x37, 0xbe, 0xd5, 0xc5, 0x3b, 0x96, 0x77, 0x0c, 0xe4, 0x1f, 0x50, 0xe6, 0x27, 0xa4, 0xa6, 0xe3,
	0x5f, 0xa4, 0x8c, 0x44, 0x43, 0x45, 0xd0, 0x5d, 0xc3, 0xcf, 0x40, 0xf6, 0xae, 0xcf, 0x88, 0x8f,
	0x22, 0x95, 0xfd, 0xe6, 0xbd, 0x79, 0x33, 0x7a, 0xbb, 0x4b, 0x1e, 0x9e, 0xd2, 0x5a, 0xb2, 0x65,
	0x49, 0xb3, 0xf8, 0xfd, 0x61, 0xca, 0x91, 0x1e, 0xc6, 0xdb, 0x4a, 0x54, 0x56, 0x80, 0xe0, 0xde,
	0xed, 0x0b, 0x46, 0x32, 0xbd, 0x97, 0x43, 0x0e, 0x2d, 0x1b, 0x37, 0x7f, 0x5a, 0x38, 0x0d, 0x72,
	0x80, 0xfc, 0x94, 0xc7, 0x2d, 0x4a, 0xeb, 0x93, 0x18, 0x45, 0xc1, 0x15, 0xd2, 0xa2, 0x34, 0x02,
	0x9f, 0x81, 0x2a, 0x40, 0xc5, 0x29, 0x55, 0x7c, 0x3b, 0x8e, 0x81, 0x90, 0x7f, 0xf1, 0x72, 0xb5,
	0xe5, 0x1b, 0x60, 0xf8, 0x3d, 0xcd, 0xbf, 0xd3, 0x93, 0x35, 0x30, 0xd4, 0x3f, 0x37, 0xda, 0xff,
	0x3a, 0x20, 0xce, 0xab, 0x6e, 0x7b, 0xf7, 0x36, 0x19, 0x88, 0xcc, 0xb3, 0x67, 0x76, 0xb8, 0x93,
	0x0c, 0x44, 0xe6, 0x7a, 0x64, 0xcc, 0x2a, 0x4e, 0x11, 0x2a, 0x6f, 0x30, 0xb3, 0x43, 0x27, 0xe9,
	0x60, 0xc3, 0x20, 0xac, 0xb8, 0x7c, 0x99, 0x79, 0xc3, 0x56, 0xde, 0x41, 0x37, 0x24, 0x77, 0x28,
	0x63, 0xbc, 0xe4, 0xc8, 0xb3, 0x63, 0x2e, 0xa1, 0x50, 0xde, 0xce, 0x6c, 0x18, 0x3a, 0xc9, 0x9f,
	0xe5, 0xc6, 0x43, 0xc1, 0x09, 0x32, 0x5a, 0x7a, 0xb7, 0xb4, 0x87, 0x81, 0x0d, 0xb3, 0xa4, 0x55,
	0xd6, 0x30, 0x23, 0xcd, 0x18, 0xe8, 0xce, 0x89, 0xa3, 0x90, 0x56, 0xb8, 0x10, 0x05, 0xf7, 0xc6,
	0x33, 0x3b, 0xdc, 0x3d, 0x9a, 0x46, 0x3a, 0xd5, 0xa8, 0x4b, 0x35, 0x5a, 0x74, 0xa9, 0xce, 0x27,
	0x97, 0xdf, 0x03, 0xeb, 0xfc, 0x3a, 0xb0, 0x93, 0xbe, 0xcd, 0x7d, 0x46, 0xc6, 0x5c, 0x66, 0xad,
	0xc3, 0xe4, 0x06, 0x0e, 0x5d, 0x93, 0x7b, 0x9f, 0x8c, 0x14, 0x52, 0xac, 0x95, 0xe7, 0xb4, 0xa1,
	0x18, 0xf4, 0x64, 0xf2, 0xe9, 0x22, 0xb0, 0x7e, 0x5e, 0x04, 0xd6, 0xfe, 0x17, 0x9b, 0xb8, 0xc7,
	0xbc, 0x04, 0x25, 0x70, 0x01, 0xff, 0x8f, 0xf7, 0x01, 0x71, 0x32, 0xad, 0xda, 0x06, 0xdc, 0x17,
	0x5c, 0x46, 0x46, 0xb4, 0x80, 0x5a, 0xa2, 0x37, 0x9c, 0x0d, 0xc3, 0xdd, 0xa3, 0xbd, 0xc8, 0x9c,
	0x67, 0x73, 0x39, 0xba, 0x8b, 0x16, 0xbd, 0x00, 0x21, 0xe7, 0x8f, 0x9a, 0x25, 0x3f, 0x5f, 0x07,
	0x61, 0x2e, 0x70, 0x59, 0xa7, 0x11, 0x83, 0xc2, 0x1c, 0xbe, 0xf9, 0x1c, 0xa8, 0x6c, 0x15, 0xe3,
	0x59, 0xc9, 0x55, 0xdb, 0xa0, 0x12, 0x63, 0xdd, 0xef, 0x3c, 0x7f, 0x73, 0xb9, 0xf6, 0xed, 0xab,
	0xb5, 0x6f, 0xff, 0x58, 0xfb, 0xf6, 0xf9, 0xc6, 0xb7, 0xae, 0x36, 0xbe, 0xf5, 0x6d, 0xe3, 0x5b,
	0x6f, 0x9f, 0xfe, 0xe6, 0xfa, 0xbc, 0x46, 0x90, 0x50, 0x9c, 0xbd, 0xe6, 0xf8, 0x01, 0xaa, 0x55,
	0x4c, 0x0d, 0x3e, 0x60, 0x4b, 0x2a, 0x64, 0xfc, 0xb1, 0x7f, 0x1b, 0x7a, 0x5c, 0x3a, 0x6a, 0x33,
	0x7d, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x74, 0x84, 0x1b, 0x47, 0x03, 0x00, 0x00,
}

func (m *Launchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Launchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Launchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintLaunchpad(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x4a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLaunchpad(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLaunchpad(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	if m.Hardcap != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.Hardcap))
		i--
		dAtA[i] = 0x30
	}
	if m.Softcap != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.Softcap))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AccepetedDenoms) > 0 {
		for iNdEx := len(m.AccepetedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AccepetedDenoms[iNdEx])
			copy(dAtA[i:], m.AccepetedDenoms[iNdEx])
			i = encodeVarintLaunchpad(dAtA, i, uint64(len(m.AccepetedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.TokenId != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLaunchpad(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DepositToLaunchpad) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositToLaunchpad) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositToLaunchpad) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintLaunchpad(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintLaunchpad(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLaunchpad(dAtA []byte, offset int, v uint64) int {
	offset -= sovLaunchpad(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Launchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLaunchpad(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLaunchpad(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovLaunchpad(uint64(m.TokenId))
	}
	if len(m.AccepetedDenoms) > 0 {
		for _, s := range m.AccepetedDenoms {
			l = len(s)
			n += 1 + l + sovLaunchpad(uint64(l))
		}
	}
	if m.Softcap != 0 {
		n += 1 + sovLaunchpad(uint64(m.Softcap))
	}
	if m.Hardcap != 0 {
		n += 1 + sovLaunchpad(uint64(m.Hardcap))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovLaunchpad(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLaunchpad(uint64(l))
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovLaunchpad(uint64(l))
	}
	return n
}

func (m *DepositToLaunchpad) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLaunchpad(uint64(m.Id))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovLaunchpad(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovLaunchpad(uint64(l))
		}
	}
	return n
}

func sovLaunchpad(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLaunchpad(x uint64) (n int) {
	return sovLaunchpad(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Launchpad) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLaunchpad
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
			return fmt.Errorf("proto: Launchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Launchpad: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
					return ErrIntOverflowLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
					return ErrIntOverflowLaunchpad
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
					return ErrIntOverflowLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLaunchpad(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLaunchpad
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
func (m *DepositToLaunchpad) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLaunchpad
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
			return fmt.Errorf("proto: DepositToLaunchpad: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositToLaunchpad: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
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
			skippy, err := skipLaunchpad(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLaunchpad
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
func skipLaunchpad(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLaunchpad
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
					return 0, ErrIntOverflowLaunchpad
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
					return 0, ErrIntOverflowLaunchpad
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
				return 0, ErrInvalidLengthLaunchpad
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLaunchpad
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLaunchpad
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLaunchpad        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLaunchpad          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLaunchpad = fmt.Errorf("proto: unexpected end of group")
)
