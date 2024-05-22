// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/evmutil/v1beta1/conversion_pair.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// ConversionPair defines a 0gChain ERC20 address and corresponding denom that is
// allowed to be converted between ERC20 and sdk.Coin
type ConversionPair struct {
	// ERC20 address of the token on the 0gChain EVM
	ZgChainERC20Address HexBytes `protobuf:"bytes,1,opt,name=zgchain_erc20_address,json=zgchainErc20Address,proto3,casttype=HexBytes" json:"zgchain_erc20_address,omitempty"`
	// Denom of the corresponding sdk.Coin
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ConversionPair) Reset()         { *m = ConversionPair{} }
func (m *ConversionPair) String() string { return proto.CompactTextString(m) }
func (*ConversionPair) ProtoMessage()    {}
func (*ConversionPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bad9d4ffa6874ec, []int{0}
}
func (m *ConversionPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConversionPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConversionPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConversionPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionPair.Merge(m, src)
}
func (m *ConversionPair) XXX_Size() int {
	return m.Size()
}
func (m *ConversionPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionPair.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionPair proto.InternalMessageInfo

// AllowedCosmosCoinERC20Token defines allowed cosmos-sdk denom & metadata
// for evm token representations of sdk assets.
// NOTE: once evm token contracts are deployed, changes to metadata for a given
// cosmos_denom will not change metadata of deployed contract.
type AllowedCosmosCoinERC20Token struct {
	// Denom of the sdk.Coin
	CosmosDenom string `protobuf:"bytes,1,opt,name=cosmos_denom,json=cosmosDenom,proto3" json:"cosmos_denom,omitempty"`
	// Name of ERC20 contract
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of ERC20 contract
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Number of decimals ERC20 contract is deployed with.
	Decimals uint32 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *AllowedCosmosCoinERC20Token) Reset()         { *m = AllowedCosmosCoinERC20Token{} }
func (m *AllowedCosmosCoinERC20Token) String() string { return proto.CompactTextString(m) }
func (*AllowedCosmosCoinERC20Token) ProtoMessage()    {}
func (*AllowedCosmosCoinERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bad9d4ffa6874ec, []int{1}
}
func (m *AllowedCosmosCoinERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedCosmosCoinERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedCosmosCoinERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedCosmosCoinERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedCosmosCoinERC20Token.Merge(m, src)
}
func (m *AllowedCosmosCoinERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *AllowedCosmosCoinERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedCosmosCoinERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedCosmosCoinERC20Token proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionPair)(nil), "zgc.evmutil.v1beta1.ConversionPair")
	proto.RegisterType((*AllowedCosmosCoinERC20Token)(nil), "zgc.evmutil.v1beta1.AllowedCosmosCoinERC20Token")
}

func init() {
	proto.RegisterFile("zgc/evmutil/v1beta1/conversion_pair.proto", fileDescriptor_6bad9d4ffa6874ec)
}

var fileDescriptor_6bad9d4ffa6874ec = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0x3f, 0x4f, 0xf2, 0x40,
	0x1c, 0xc7, 0x7b, 0xcf, 0xc3, 0x43, 0x78, 0x4e, 0x74, 0x38, 0xd0, 0x34, 0x98, 0x1c, 0x88, 0x0b,
	0x9a, 0xd8, 0x16, 0xdc, 0xdc, 0xa0, 0x92, 0x38, 0x38, 0x98, 0xc6, 0xc4, 0x84, 0xa5, 0xb9, 0x5e,
	0x2f, 0x47, 0x63, 0xdb, 0x23, 0xbd, 0x82, 0xc0, 0x6e, 0xe2, 0x64, 0x7c, 0x09, 0x8e, 0xbe, 0x14,
	0x47, 0x46, 0x27, 0x82, 0xe5, 0x5d, 0x38, 0x99, 0xfe, 0x09, 0xdb, 0xef, 0xfb, 0xfd, 0x7d, 0xf2,
	0xc9, 0xfd, 0x81, 0x67, 0x4b, 0x4e, 0x75, 0x36, 0x0b, 0xa6, 0xb1, 0xe7, 0xeb, 0xb3, 0xae, 0xc3,
	0x62, 0xd2, 0xd5, 0xa9, 0x08, 0x67, 0x2c, 0x92, 0x9e, 0x08, 0xed, 0x09, 0xf1, 0x22, 0x6d, 0x12,
	0x89, 0x58, 0xa0, 0xda, 0x92, 0x53, 0xad, 0x40, 0xb5, 0x02, 0x6d, 0xd4, 0xb9, 0xe0, 0x22, 0xdb,
	0xeb, 0xe9, 0x94, 0xa3, 0xed, 0x67, 0x00, 0x0f, 0xcc, 0x9d, 0xe4, 0x8e, 0x78, 0x11, 0x7a, 0x80,
	0x87, 0x4b, 0x4e, 0xc7, 0xc4, 0x0b, 0x6d, 0x16, 0xd1, 0x9e, 0x61, 0x13, 0xd7, 0x8d, 0x98, 0x94,
	0x2a, 0x68, 0x81, 0x4e, 0x75, 0x70, 0x9a, 0xac, 0x9b, 0xb5, 0x11, 0x37, 0x53, 0x60, 0x68, 0x99,
	0x3d, 0xa3, 0x9f, 0xaf, 0x7f, 0xd6, 0xcd, 0xca, 0x0d, 0x9b, 0x0f, 0x16, 0x31, 0x93, 0x56, 0xad,
	0x30, 0x0c, 0x53, 0x41, 0x01, 0xa0, 0x3a, 0xfc, 0xe7, 0xb2, 0x50, 0x04, 0xea, 0x9f, 0x16, 0xe8,
	0xfc, 0xb7, 0xf2, 0x70, 0x55, 0x7a, 0x79, 0x6f, 0x2a, 0xed, 0x57, 0x00, 0x8f, 0xfb, 0xbe, 0x2f,
	0x9e, 0x98, 0x6b, 0x0a, 0x19, 0x08, 0x69, 0x8a, 0x42, 0x7f, 0x2f, 0x1e, 0x59, 0x88, 0x4e, 0x60,
	0x95, 0x66, 0xbd, 0x9d, 0x2b, 0x40, 0xa6, 0xd8, 0xcb, 0xbb, 0xeb, 0xb4, 0x42, 0x08, 0x96, 0x42,
	0x12, 0xb0, 0xc2, 0x9e, 0xcd, 0xe8, 0x08, 0x96, 0xe5, 0x22, 0x70, 0x84, 0xaf, 0xfe, 0xcd, 0xda,
	0x22, 0xa1, 0x06, 0xac, 0xb8, 0x8c, 0x7a, 0x01, 0xf1, 0xa5, 0x5a, 0x6a, 0x81, 0xce, 0xbe, 0xb5,
	0xcb, 0xf9, 0x81, 0x06, 0xb7, 0x9b, 0x6f, 0x0c, 0x3e, 0x12, 0x0c, 0x3e, 0x13, 0x0c, 0x56, 0x09,
	0x06, 0x9b, 0x04, 0x83, 0xb7, 0x2d, 0x56, 0x56, 0x5b, 0xac, 0x7c, 0x6d, 0xb1, 0x32, 0x3a, 0xe7,
	0x5e, 0x3c, 0x9e, 0x3a, 0x1a, 0x15, 0x81, 0x6e, 0x70, 0x9f, 0x38, 0x52, 0x37, 0xf8, 0x45, 0x76,
	0x6d, 0x7d, 0xbe, 0xfb, 0xa9, 0x78, 0x31, 0x61, 0xd2, 0x29, 0x67, 0xaf, 0x7d, 0xf9, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x07, 0x29, 0xab, 0xc5, 0x01, 0x00, 0x00,
}

func (this *ConversionPair) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ConversionPair")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ConversionPair but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ConversionPair but is not nil && this == nil")
	}
	if !bytes.Equal(this.ZgChainERC20Address, that1.ZgChainERC20Address) {
		return fmt.Errorf("ZgChainERC20Address this(%v) Not Equal that(%v)", this.ZgChainERC20Address, that1.ZgChainERC20Address)
	}
	if this.Denom != that1.Denom {
		return fmt.Errorf("Denom this(%v) Not Equal that(%v)", this.Denom, that1.Denom)
	}
	return nil
}
func (this *ConversionPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.ZgChainERC20Address, that1.ZgChainERC20Address) {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (this *AllowedCosmosCoinERC20Token) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AllowedCosmosCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedCosmosCoinERC20Token)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AllowedCosmosCoinERC20Token")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AllowedCosmosCoinERC20Token but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AllowedCosmosCoinERC20Token but is not nil && this == nil")
	}
	if this.CosmosDenom != that1.CosmosDenom {
		return fmt.Errorf("CosmosDenom this(%v) Not Equal that(%v)", this.CosmosDenom, that1.CosmosDenom)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Symbol != that1.Symbol {
		return fmt.Errorf("Symbol this(%v) Not Equal that(%v)", this.Symbol, that1.Symbol)
	}
	if this.Decimals != that1.Decimals {
		return fmt.Errorf("Decimals this(%v) Not Equal that(%v)", this.Decimals, that1.Decimals)
	}
	return nil
}
func (this *AllowedCosmosCoinERC20Token) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AllowedCosmosCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedCosmosCoinERC20Token)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CosmosDenom != that1.CosmosDenom {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Decimals != that1.Decimals {
		return false
	}
	return true
}
func (m *ConversionPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConversionPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConversionPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ZgChainERC20Address) > 0 {
		i -= len(m.ZgChainERC20Address)
		copy(dAtA[i:], m.ZgChainERC20Address)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.ZgChainERC20Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllowedCosmosCoinERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedCosmosCoinERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedCosmosCoinERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintConversionPair(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CosmosDenom) > 0 {
		i -= len(m.CosmosDenom)
		copy(dAtA[i:], m.CosmosDenom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.CosmosDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConversionPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovConversionPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConversionPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ZgChainERC20Address)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	return n
}

func (m *AllowedCosmosCoinERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CosmosDenom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovConversionPair(uint64(m.Decimals))
	}
	return n
}

func sovConversionPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConversionPair(x uint64) (n int) {
	return sovConversionPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConversionPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: ConversionPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConversionPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZgChainERC20Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZgChainERC20Address = append(m.ZgChainERC20Address[:0], dAtA[iNdEx:postIndex]...)
			if m.ZgChainERC20Address == nil {
				m.ZgChainERC20Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func (m *AllowedCosmosCoinERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: AllowedCosmosCoinERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedCosmosCoinERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func skipConversionPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
				return 0, ErrInvalidLengthConversionPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConversionPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConversionPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConversionPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConversionPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConversionPair = fmt.Errorf("proto: unexpected end of group")
)
