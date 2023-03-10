// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/coin_symbol.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type CoinSymbol struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol    string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	IsDeleted bool   `protobuf:"varint,3,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
}

func (m *CoinSymbol) Reset()         { *m = CoinSymbol{} }
func (m *CoinSymbol) String() string { return proto.CompactTextString(m) }
func (*CoinSymbol) ProtoMessage()    {}
func (*CoinSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ce43cd4e84a56d4, []int{0}
}
func (m *CoinSymbol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinSymbol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinSymbol.Merge(m, src)
}
func (m *CoinSymbol) XXX_Size() int {
	return m.Size()
}
func (m *CoinSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_CoinSymbol proto.InternalMessageInfo

func (m *CoinSymbol) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CoinSymbol) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *CoinSymbol) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func init() {
	proto.RegisterType((*CoinSymbol)(nil), "spidex.oracle.v1.CoinSymbol")
}

func init() { proto.RegisterFile("oracle/v1/coin_symbol.proto", fileDescriptor_2ce43cd4e84a56d4) }

var fileDescriptor_2ce43cd4e84a56d4 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0xcc, 0x8b, 0x2f, 0xae, 0xcc, 0x4d, 0xca,
	0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0xc8, 0x4c, 0x49, 0xad, 0xd0,
	0x83, 0xa8, 0xd1, 0x2b, 0x33, 0x54, 0x0a, 0xe2, 0xe2, 0x72, 0xce, 0xcf, 0xcc, 0x0b, 0x06, 0xab,
	0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c,
	0x11, 0x12, 0xe3, 0x62, 0x83, 0xe8, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84,
	0x64, 0xb8, 0x38, 0x33, 0x8b, 0x5d, 0x52, 0x73, 0x52, 0x4b, 0x52, 0x53, 0x24, 0x98, 0x15, 0x18,
	0x35, 0x38, 0x82, 0x10, 0x02, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x71, 0x4a,
	0x62, 0x41, 0x01, 0x94, 0xa5, 0x5f, 0xa1, 0x0f, 0x75, 0x7a, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12,
	0x1b, 0xd8, 0xc9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf3, 0x2d, 0x02, 0xd1, 0x00,
	0x00, 0x00,
}

func (m *CoinSymbol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinSymbol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinSymbol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsDeleted {
		i--
		if m.IsDeleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintCoinSymbol(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCoinSymbol(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoinSymbol(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinSymbol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoinSymbol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCoinSymbol(uint64(m.Id))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovCoinSymbol(uint64(l))
	}
	if m.IsDeleted {
		n += 2
	}
	return n
}

func sovCoinSymbol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinSymbol(x uint64) (n int) {
	return sovCoinSymbol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoinSymbol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinSymbol
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
			return fmt.Errorf("proto: CoinSymbol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinSymbol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinSymbol
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
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinSymbol
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
				return ErrInvalidLengthCoinSymbol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinSymbol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDeleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinSymbol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsDeleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCoinSymbol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinSymbol
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
func skipCoinSymbol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinSymbol
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
					return 0, ErrIntOverflowCoinSymbol
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
					return 0, ErrIntOverflowCoinSymbol
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
				return 0, ErrInvalidLengthCoinSymbol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinSymbol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinSymbol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinSymbol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinSymbol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinSymbol = fmt.Errorf("proto: unexpected end of group")
)
