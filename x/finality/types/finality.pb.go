// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/finality.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// IndexedBlock is the necessary metadata and finalization status of a block
type IndexedBlock struct {
	// height is the height of the block
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// app_hash is the AppHash of the block
	AppHash []byte `protobuf:"bytes,2,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	// finalized indicates whether the IndexedBlock is finalised by 2/3
	// finality providers or not
	Finalized bool `protobuf:"varint,3,opt,name=finalized,proto3" json:"finalized,omitempty"`
}

func (m *IndexedBlock) Reset()         { *m = IndexedBlock{} }
func (m *IndexedBlock) String() string { return proto.CompactTextString(m) }
func (*IndexedBlock) ProtoMessage()    {}
func (*IndexedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{0}
}
func (m *IndexedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedBlock.Merge(m, src)
}
func (m *IndexedBlock) XXX_Size() int {
	return m.Size()
}
func (m *IndexedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedBlock proto.InternalMessageInfo

func (m *IndexedBlock) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *IndexedBlock) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *IndexedBlock) GetFinalized() bool {
	if m != nil {
		return m.Finalized
	}
	return false
}

// Evidence is the evidence that a finality provider has signed finality
// signatures with correct public randomness on two conflicting Babylon headers
type Evidence struct {
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// block_height is the height of the conflicting blocks
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// master_pub_rand is the master public randomness the finality provider has committed to
	// encoded as a base58 string
	MasterPubRand string `protobuf:"bytes,3,opt,name=master_pub_rand,json=masterPubRand,proto3" json:"master_pub_rand,omitempty"`
	// canonical_app_hash is the AppHash of the canonical block
	CanonicalAppHash []byte `protobuf:"bytes,4,opt,name=canonical_app_hash,json=canonicalAppHash,proto3" json:"canonical_app_hash,omitempty"`
	// fork_app_hash is the AppHash of the fork block
	ForkAppHash []byte `protobuf:"bytes,5,opt,name=fork_app_hash,json=forkAppHash,proto3" json:"fork_app_hash,omitempty"`
	// canonical_finality_sig is the finality signature to the canonical block
	// where finality signature is an EOTS signature, i.e.,
	// the `s` in a Schnorr signature `(r, s)`
	// `r` is the public randomness that is already committed by the finality provider
	CanonicalFinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,6,opt,name=canonical_finality_sig,json=canonicalFinalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"canonical_finality_sig,omitempty"`
	// fork_finality_sig is the finality signature to the fork block
	// where finality signature is an EOTS signature
	ForkFinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,7,opt,name=fork_finality_sig,json=forkFinalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"fork_finality_sig,omitempty"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{1}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Evidence) GetMasterPubRand() string {
	if m != nil {
		return m.MasterPubRand
	}
	return ""
}

func (m *Evidence) GetCanonicalAppHash() []byte {
	if m != nil {
		return m.CanonicalAppHash
	}
	return nil
}

func (m *Evidence) GetForkAppHash() []byte {
	if m != nil {
		return m.ForkAppHash
	}
	return nil
}

// FinalityProviderSigningInfo defines a finality provider's signing info for monitoring their
// liveness activity.
type FinalityProviderSigningInfo struct {
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// start_height is the block height at which finality provider become active OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// jailed_until defines the timestamp until which the finality provider is jailed due to liveness downtime.
	JailedUntil time.Time `protobuf:"bytes,3,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until"`
	// missed_blocks_counter defines a counter to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter int64 `protobuf:"varint,4,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty"`
}

func (m *FinalityProviderSigningInfo) Reset()         { *m = FinalityProviderSigningInfo{} }
func (m *FinalityProviderSigningInfo) String() string { return proto.CompactTextString(m) }
func (*FinalityProviderSigningInfo) ProtoMessage()    {}
func (*FinalityProviderSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{2}
}
func (m *FinalityProviderSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalityProviderSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalityProviderSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalityProviderSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalityProviderSigningInfo.Merge(m, src)
}
func (m *FinalityProviderSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *FinalityProviderSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalityProviderSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FinalityProviderSigningInfo proto.InternalMessageInfo

func (m *FinalityProviderSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *FinalityProviderSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *FinalityProviderSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*IndexedBlock)(nil), "babylon.finality.v1.IndexedBlock")
	proto.RegisterType((*Evidence)(nil), "babylon.finality.v1.Evidence")
	proto.RegisterType((*FinalityProviderSigningInfo)(nil), "babylon.finality.v1.FinalityProviderSigningInfo")
}

func init() {
	proto.RegisterFile("babylon/finality/v1/finality.proto", fileDescriptor_ca5b87e52e3e6d02)
}

var fileDescriptor_ca5b87e52e3e6d02 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0x8d, 0x9b, 0xfe, 0xb9, 0x4c, 0x12, 0xf5, 0xaf, 0x5b, 0xaa, 0x10, 0x90, 0x13, 0xb2, 0x40,
	0x11, 0x42, 0x76, 0x9b, 0x56, 0x88, 0x2d, 0x46, 0x45, 0x0d, 0x20, 0x11, 0x39, 0x65, 0xc3, 0x66,
	0x34, 0xb6, 0xc7, 0xf6, 0x10, 0x7b, 0xc6, 0xb2, 0xc7, 0x51, 0xc3, 0x53, 0x74, 0xc5, 0x33, 0xb0,
	0xe4, 0x31, 0xba, 0x2c, 0x3b, 0xd4, 0x45, 0x40, 0xc9, 0x82, 0xd7, 0x40, 0x1e, 0xdb, 0x49, 0x59,
	0x81, 0x84, 0xd8, 0x44, 0x33, 0xe7, 0xbb, 0x9d, 0x93, 0x33, 0x9f, 0x41, 0xdf, 0x44, 0xe6, 0xdc,
	0x67, 0x54, 0x73, 0x08, 0x45, 0x3e, 0xe1, 0x73, 0x6d, 0x76, 0xb4, 0x3e, 0xab, 0x61, 0xc4, 0x38,
	0x93, 0xf7, 0xf2, 0x1c, 0x75, 0x8d, 0xcf, 0x8e, 0x3a, 0x5d, 0x97, 0x31, 0xd7, 0xc7, 0x9a, 0x48,
	0x31, 0x13, 0x47, 0xe3, 0x24, 0xc0, 0x31, 0x47, 0x41, 0x98, 0x55, 0x75, 0x76, 0x51, 0x40, 0x28,
	0xd3, 0xc4, 0x6f, 0x0e, 0xed, 0xbb, 0xcc, 0x65, 0xe2, 0xa8, 0xa5, 0xa7, 0x0c, 0xed, 0x43, 0xd0,
	0x1c, 0x51, 0x1b, 0x5f, 0x60, 0x5b, 0xf7, 0x99, 0x35, 0x95, 0x0f, 0x40, 0xc5, 0xc3, 0xc4, 0xf5,
	0x78, 0x5b, 0xea, 0x49, 0x83, 0x6d, 0x23, 0xbf, 0xc9, 0x77, 0x41, 0x0d, 0x85, 0x21, 0xf4, 0x50,
	0xec, 0xb5, 0xb7, 0x7a, 0xd2, 0xa0, 0x69, 0x54, 0x51, 0x18, 0x9e, 0xa1, 0xd8, 0x93, 0xef, 0x83,
	0x7a, 0xc6, 0xed, 0x03, 0xb6, 0xdb, 0xe5, 0x9e, 0x34, 0xa8, 0x19, 0x1b, 0xa0, 0xff, 0xa5, 0x0c,
	0x6a, 0xa7, 0x33, 0x62, 0x63, 0x6a, 0x61, 0xd9, 0x00, 0x75, 0x27, 0x84, 0x26, 0xb7, 0x60, 0x38,
	0x15, 0x03, 0x9a, 0xfa, 0x93, 0x9b, 0x45, 0x77, 0xe8, 0x12, 0xee, 0x25, 0xa6, 0x6a, 0xb1, 0x40,
	0xcb, 0xe5, 0x5a, 0x1e, 0x22, 0xb4, 0xb8, 0x68, 0x7c, 0x1e, 0xe2, 0x58, 0xd5, 0x47, 0xe3, 0xe3,
	0x93, 0xc3, 0x71, 0x62, 0xbe, 0xc2, 0x73, 0xa3, 0xea, 0x84, 0x3a, 0xb7, 0xc6, 0x53, 0xf9, 0x01,
	0x68, 0x9a, 0x29, 0x75, 0x98, 0xf3, 0xde, 0x12, 0xbc, 0x1b, 0x02, 0x3b, 0xcb, 0xc8, 0x3f, 0x04,
	0x3b, 0x01, 0x8a, 0x39, 0x8e, 0x60, 0x98, 0x98, 0x30, 0x42, 0x34, 0xe3, 0x59, 0x37, 0x5a, 0x19,
	0x3c, 0x4e, 0x4c, 0x03, 0x51, 0x5b, 0x7e, 0x0c, 0x64, 0x0b, 0x51, 0x46, 0x89, 0x85, 0x7c, 0xb8,
	0x96, 0xbb, 0x2d, 0xe4, 0xfe, 0xbf, 0x8e, 0x3c, 0xcb, 0x75, 0xf7, 0x41, 0xcb, 0x61, 0xd1, 0x74,
	0x93, 0xf8, 0x9f, 0x48, 0x6c, 0xa4, 0x60, 0x91, 0x43, 0xc1, 0xc1, 0xa6, 0x63, 0xe1, 0x20, 0x8c,
	0x89, 0xdb, 0xae, 0x08, 0xf5, 0x4f, 0x6f, 0x16, 0xdd, 0x93, 0x3f, 0x53, 0x3f, 0xb1, 0x3c, 0xca,
	0xa2, 0xe8, 0xf4, 0xcd, 0xf9, 0x64, 0x42, 0x5c, 0x63, 0x7f, 0xdd, 0xf7, 0x45, 0xde, 0x76, 0x42,
	0x5c, 0xd9, 0x06, 0xbb, 0x82, 0xd3, 0x2f, 0xa3, 0xaa, 0x7f, 0x39, 0x6a, 0x27, 0x6d, 0x79, 0x6b,
	0x4a, 0xff, 0xe3, 0x16, 0xb8, 0x57, 0xdc, 0xc7, 0x11, 0x4b, 0xdd, 0x8d, 0x26, 0xc4, 0xa5, 0x84,
	0xba, 0x23, 0xea, 0xb0, 0x7f, 0x65, 0x73, 0xcc, 0x51, 0xc4, 0x6f, 0xdb, 0x5c, 0x36, 0x1a, 0x02,
	0xcb, 0x6d, 0x7e, 0x0d, 0x9a, 0xef, 0x11, 0xf1, 0xb1, 0x0d, 0x13, 0xca, 0x89, 0x2f, 0x3c, 0x6e,
	0x0c, 0x3b, 0x6a, 0xb6, 0x2c, 0x6a, 0xb1, 0x2c, 0xea, 0x79, 0xb1, 0x2c, 0x7a, 0xeb, 0x6a, 0xd1,
	0x2d, 0x5d, 0x7e, 0xeb, 0x4a, 0x9f, 0x7e, 0x7c, 0x7e, 0x24, 0x19, 0x8d, 0xac, 0xfc, 0x6d, 0x5a,
	0x2d, 0x0f, 0xc1, 0x9d, 0x80, 0xc4, 0x31, 0xb6, 0xa1, 0x78, 0x4a, 0x31, 0xb4, 0x58, 0x42, 0x39,
	0x8e, 0xc4, 0x7b, 0x28, 0x1b, 0x7b, 0x59, 0x50, 0x6c, 0x4d, 0xfc, 0x3c, 0x0b, 0xe9, 0x2f, 0xaf,
	0x96, 0x8a, 0x74, 0xbd, 0x54, 0xa4, 0xef, 0x4b, 0x45, 0xba, 0x5c, 0x29, 0xa5, 0xeb, 0x95, 0x52,
	0xfa, 0xba, 0x52, 0x4a, 0xef, 0x0e, 0x7f, 0xa7, 0xfd, 0x62, 0xf3, 0x11, 0x10, 0x7f, 0x83, 0x59,
	0x11, 0x7c, 0x8f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd3, 0x3e, 0xc7, 0x25, 0x04, 0x00,
	0x00,
}

func (m *IndexedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Finalized {
		i--
		if m.Finalized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ForkFinalitySig != nil {
		{
			size := m.ForkFinalitySig.Size()
			i -= size
			if _, err := m.ForkFinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CanonicalFinalitySig != nil {
		{
			size := m.CanonicalFinalitySig.Size()
			i -= size
			if _, err := m.CanonicalFinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ForkAppHash) > 0 {
		i -= len(m.ForkAppHash)
		copy(dAtA[i:], m.ForkAppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.ForkAppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CanonicalAppHash) > 0 {
		i -= len(m.CanonicalAppHash)
		copy(dAtA[i:], m.CanonicalAppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.CanonicalAppHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MasterPubRand) > 0 {
		i -= len(m.MasterPubRand)
		copy(dAtA[i:], m.MasterPubRand)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.MasterPubRand)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FinalityProviderSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalityProviderSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalityProviderSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintFinality(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.StartHeight != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFinality(dAtA []byte, offset int, v uint64) int {
	offset -= sovFinality(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovFinality(uint64(m.Height))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.Finalized {
		n += 2
	}
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovFinality(uint64(m.BlockHeight))
	}
	l = len(m.MasterPubRand)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	l = len(m.CanonicalAppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	l = len(m.ForkAppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.CanonicalFinalitySig != nil {
		l = m.CanonicalFinalitySig.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.ForkFinalitySig != nil {
		l = m.ForkFinalitySig.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	return n
}

func (m *FinalityProviderSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovFinality(uint64(m.StartHeight))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovFinality(uint64(l))
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovFinality(uint64(m.MissedBlocksCounter))
	}
	return n
}

func sovFinality(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFinality(x uint64) (n int) {
	return sovFinality(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: IndexedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
			m.Finalized = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterPubRand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterPubRand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalAppHash = append(m.CanonicalAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CanonicalAppHash == nil {
				m.CanonicalAppHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForkAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForkAppHash = append(m.ForkAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ForkAppHash == nil {
				m.ForkAppHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalFinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.CanonicalFinalitySig = &v
			if err := m.CanonicalFinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForkFinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.ForkFinalitySig = &v
			if err := m.ForkFinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func (m *FinalityProviderSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: FinalityProviderSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalityProviderSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func skipFinality(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
				return 0, ErrInvalidLengthFinality
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFinality
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFinality
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFinality        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFinality          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFinality = fmt.Errorf("proto: unexpected end of group")
)
