// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Params struct {
	TokensPerVote     uint64 `protobuf:"varint,1,opt,name=tokens_per_vote,json=tokensPerVote,proto3" json:"tokens_per_vote,omitempty"`
	MaxVotesPerSigner uint64 `protobuf:"varint,2,opt,name=max_votes_per_signer,json=maxVotesPerSigner,proto3" json:"max_votes_per_signer,omitempty"`
	MaxQuorums        uint64 `protobuf:"varint,3,opt,name=max_quorums,json=maxQuorums,proto3" json:"max_quorums,omitempty"`
	EpochBlocks       uint64 `protobuf:"varint,4,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty"`
	EncodedSlices     uint64 `protobuf:"varint,5,opt,name=encoded_slices,json=encodedSlices,proto3" json:"encoded_slices,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTokensPerVote() uint64 {
	if m != nil {
		return m.TokensPerVote
	}
	return 0
}

func (m *Params) GetMaxVotesPerSigner() uint64 {
	if m != nil {
		return m.MaxVotesPerSigner
	}
	return 0
}

func (m *Params) GetMaxQuorums() uint64 {
	if m != nil {
		return m.MaxQuorums
	}
	return 0
}

func (m *Params) GetEpochBlocks() uint64 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

func (m *Params) GetEncodedSlices() uint64 {
	if m != nil {
		return m.EncodedSlices
	}
	return 0
}

// GenesisState defines the dasigners module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// params epoch_number the epoch number
	EpochNumber uint64 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	// signers defines all signers information
	Signers []*Signer `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
	// quorums_by_epoch defines chosen quorums by epoch
	QuorumsByEpoch []*Quorums `protobuf:"bytes,4,rep,name=quorums_by_epoch,json=quorumsByEpoch,proto3" json:"quorums_by_epoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *GenesisState) GetSigners() []*Signer {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *GenesisState) GetQuorumsByEpoch() []*Quorums {
	if m != nil {
		return m.QuorumsByEpoch
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "zgc.dasigners.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "zgc.dasigners.v1.GenesisState")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/genesis.proto", fileDescriptor_896efa766aaca3be) }

var fileDescriptor_896efa766aaca3be = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb3, 0x24, 0x04, 0xc9, 0x29, 0xa5, 0x58, 0x3d, 0x6c, 0x7a, 0xd8, 0x84, 0x4a, 0x20,
	0x2e, 0xac, 0xdb, 0x22, 0xf1, 0x00, 0x41, 0x08, 0x71, 0x41, 0x25, 0x91, 0x38, 0x70, 0x59, 0x79,
	0x9d, 0xc1, 0x59, 0x35, 0xde, 0x59, 0xd6, 0xde, 0x28, 0xe9, 0x53, 0xf0, 0x58, 0x3d, 0x70, 0xe8,
	0x91, 0x13, 0x42, 0xc9, 0x8b, 0xa0, 0x1d, 0x9b, 0x56, 0xb4, 0xdc, 0xec, 0xff, 0xff, 0x66, 0xf4,
	0xcf, 0xd8, 0x2c, 0xb9, 0xd4, 0x4a, 0xcc, 0xa5, 0x2d, 0x74, 0x09, 0xb5, 0x15, 0xab, 0x53, 0xa1,
	0xa1, 0x04, 0x5b, 0xd8, 0xb4, 0xaa, 0xd1, 0x21, 0x3f, 0xb8, 0xd4, 0x2a, 0xbd, 0xf1, 0xd3, 0xd5,
	0xe9, 0xd1, 0x50, 0xa1, 0x35, 0x68, 0x33, 0xf2, 0x85, 0xbf, 0x78, 0xf8, 0xe8, 0x50, 0xa3, 0x46,
	0xaf, 0xb7, 0xa7, 0xa0, 0x0e, 0x35, 0xa2, 0x5e, 0x82, 0xa0, 0x5b, 0xde, 0x7c, 0x15, 0xb2, 0xdc,
	0x04, 0x6b, 0x74, 0xd7, 0x72, 0x85, 0x01, 0xeb, 0xa4, 0xa9, 0x02, 0x30, 0xbe, 0x17, 0xef, 0x36,
	0x0b, 0x11, 0xc7, 0x3f, 0x22, 0xd6, 0x3f, 0x97, 0xb5, 0x34, 0x96, 0xbf, 0x60, 0x4f, 0x1c, 0x5e,
	0x40, 0x69, 0xb3, 0x0a, 0xea, 0x6c, 0x85, 0x0e, 0xe2, 0x68, 0x1c, 0xbd, 0xec, 0x4d, 0x1f, 0x7b,
	0xf9, 0x1c, 0xea, 0xcf, 0xe8, 0x80, 0x0b, 0x76, 0x68, 0xe4, 0x9a, 0x00, 0x8f, 0xfa, 0x8e, 0xf1,
	0x03, 0x82, 0x9f, 0x1a, 0xb9, 0x6e, 0xb1, 0x16, 0x9f, 0x91, 0xc1, 0x47, 0x6c, 0xd0, 0x16, 0x7c,
	0x6b, 0xb0, 0x6e, 0x8c, 0x8d, 0xbb, 0xc4, 0x31, 0x23, 0xd7, 0x9f, 0xbc, 0xc2, 0x9f, 0xb1, 0x3d,
	0xa8, 0x50, 0x2d, 0xb2, 0x7c, 0x89, 0xea, 0xc2, 0xc6, 0x3d, 0x22, 0x06, 0xa4, 0x4d, 0x48, 0xe2,
	0xcf, 0xd9, 0x3e, 0x94, 0x0a, 0xe7, 0x30, 0xcf, 0xec, 0xb2, 0x50, 0x60, 0xe3, 0x87, 0x3e, 0x5b,
	0x50, 0x67, 0x24, 0x1e, 0x6f, 0x23, 0xb6, 0xf7, 0xde, 0xbf, 0xc0, 0xcc, 0x49, 0x07, 0xfc, 0x0d,
	0xeb, 0x57, 0x34, 0x1e, 0xcd, 0x32, 0x38, 0x8b, 0xd3, 0xbb, 0x2f, 0x92, 0xfa, 0xf1, 0x27, 0xbd,
	0xab, 0x5f, 0xa3, 0xce, 0x34, 0xd0, 0xb7, 0x91, 0xca, 0xc6, 0xe4, 0x37, 0xc3, 0xf9, 0x48, 0x1f,
	0x49, 0xe2, 0x67, 0xec, 0x51, 0xe8, 0x12, 0x77, 0xc7, 0xdd, 0xff, 0xf7, 0xf6, 0x1b, 0x98, 0xfe,
	0x05, 0xf9, 0x5b, 0x76, 0x10, 0xd6, 0x90, 0xe5, 0x9b, 0x8c, 0xba, 0xc5, 0x3d, 0x2a, 0x1e, 0xde,
	0x2f, 0x0e, 0xeb, 0x99, 0xee, 0x87, 0x92, 0xc9, 0xe6, 0x1d, 0x6d, 0xe4, 0xc3, 0xd5, 0x36, 0x89,
	0xae, 0xb7, 0x49, 0xf4, 0x7b, 0x9b, 0x44, 0xdf, 0x77, 0x49, 0xe7, 0x7a, 0x97, 0x74, 0x7e, 0xee,
	0x92, 0xce, 0x17, 0xa1, 0x0b, 0xb7, 0x68, 0xf2, 0x54, 0xa1, 0x11, 0x27, 0x7a, 0x29, 0x73, 0x2b,
	0x4e, 0xf4, 0x2b, 0xb5, 0x90, 0x45, 0x29, 0xd6, 0xff, 0x7e, 0x04, 0xb7, 0xa9, 0xc0, 0xe6, 0x7d,
	0xfa, 0x05, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xe5, 0x90, 0xe2, 0xc8, 0x02, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EncodedSlices != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EncodedSlices))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochBlocks != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxQuorums != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxQuorums))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxVotesPerSigner != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxVotesPerSigner))
		i--
		dAtA[i] = 0x10
	}
	if m.TokensPerVote != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TokensPerVote))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QuorumsByEpoch) > 0 {
		for iNdEx := len(m.QuorumsByEpoch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QuorumsByEpoch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EpochNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokensPerVote != 0 {
		n += 1 + sovGenesis(uint64(m.TokensPerVote))
	}
	if m.MaxVotesPerSigner != 0 {
		n += 1 + sovGenesis(uint64(m.MaxVotesPerSigner))
	}
	if m.MaxQuorums != 0 {
		n += 1 + sovGenesis(uint64(m.MaxQuorums))
	}
	if m.EpochBlocks != 0 {
		n += 1 + sovGenesis(uint64(m.EpochBlocks))
	}
	if m.EncodedSlices != 0 {
		n += 1 + sovGenesis(uint64(m.EncodedSlices))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovGenesis(uint64(m.EpochNumber))
	}
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.QuorumsByEpoch) > 0 {
		for _, e := range m.QuorumsByEpoch {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerVote", wireType)
			}
			m.TokensPerVote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokensPerVote |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotesPerSigner", wireType)
			}
			m.MaxVotesPerSigner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxVotesPerSigner |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxQuorums", wireType)
			}
			m.MaxQuorums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxQuorums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedSlices", wireType)
			}
			m.EncodedSlices = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncodedSlices |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &Signer{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumsByEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuorumsByEpoch = append(m.QuorumsByEpoch, &Quorums{})
			if err := m.QuorumsByEpoch[len(m.QuorumsByEpoch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
