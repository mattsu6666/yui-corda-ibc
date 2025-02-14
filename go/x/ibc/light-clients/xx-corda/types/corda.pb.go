// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/corda/v1/corda.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type ClientState struct {
	BaseId    *StateRef  `protobuf:"bytes,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	NotaryKey *PublicKey `protobuf:"bytes,2,opt,name=notary_key,json=notaryKey,proto3" json:"notary_key,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43a9edf2977fc463, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func (m *ClientState) GetBaseId() *StateRef {
	if m != nil {
		return m.BaseId
	}
	return nil
}

func (m *ClientState) GetNotaryKey() *PublicKey {
	if m != nil {
		return m.NotaryKey
	}
	return nil
}

type ConsensusState struct {
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43a9edf2977fc463, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

type Header struct {
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_43a9edf2977fc463, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.corda.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.corda.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.corda.v1.Header")
}

func init() {
	proto.RegisterFile("ibc/lightclients/corda/v1/corda.proto", fileDescriptor_43a9edf2977fc463)
}

var fileDescriptor_43a9edf2977fc463 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x9b, 0x77, 0xe8, 0xab, 0x29, 0x88, 0x74, 0xaa, 0x0e, 0x41, 0xaa, 0x82, 0x20, 0x97,
	0x50, 0x5d, 0x9d, 0xec, 0xa2, 0x74, 0x50, 0xce, 0xcd, 0xa5, 0x24, 0xb9, 0xbf, 0x77, 0xc1, 0xf3,
	0x72, 0x24, 0xb9, 0x72, 0xf9, 0x16, 0xfd, 0x58, 0x8e, 0x1d, 0x1d, 0xe5, 0xee, 0x8b, 0x48, 0x13,
	0x14, 0x17, 0x75, 0x0b, 0xe4, 0x79, 0x7e, 0x0f, 0xff, 0x1f, 0x3e, 0x55, 0x42, 0xb2, 0x52, 0xe5,
	0x85, 0x93, 0xa5, 0x82, 0xca, 0x59, 0x26, 0xb5, 0xc9, 0x38, 0x5b, 0xcd, 0xe2, 0x83, 0xd6, 0x46,
	0x3b, 0x3d, 0x3e, 0x50, 0x42, 0xd2, 0xef, 0x31, 0x1a, 0x7f, 0x57, 0xb3, 0xc3, 0xf3, 0x3f, 0x08,
	0x89, 0xf3, 0x35, 0xd8, 0xc8, 0x99, 0xae, 0x11, 0x1e, 0xcd, 0x43, 0xec, 0xc1, 0x71, 0x07, 0xe3,
	0x2b, 0xfc, 0x5f, 0x70, 0x0b, 0x4b, 0x95, 0x4d, 0xd0, 0x11, 0x3a, 0x1b, 0x5d, 0x1c, 0xd3, 0x1f,
	0x97, 0x68, 0xa8, 0xa4, 0xf0, 0x94, 0x0e, 0xb7, 0x9d, 0xdb, 0x6c, 0x3c, 0xc7, 0xb8, 0xd2, 0x8e,
	0x1b, 0xbf, 0x7c, 0x06, 0x3f, 0xf9, 0x17, 0x00, 0x27, 0xbf, 0x00, 0xee, 0x1b, 0x51, 0x2a, 0xb9,
	0x00, 0x9f, 0xee, 0xc6, 0xde, 0x02, 0xfc, 0x74, 0x1f, 0xef, 0xcd, 0x75, 0x65, 0xa1, 0xb2, 0x8d,
	0x0d, 0x0b, 0xd3, 0x1d, 0x3c, 0xbc, 0x01, 0x9e, 0x81, 0xb9, 0x56, 0xaf, 0x1d, 0x41, 0x9b, 0x8e,
	0xa0, 0xf7, 0x8e, 0xa0, 0x75, 0x4f, 0x06, 0x9b, 0x9e, 0x0c, 0xde, 0x7a, 0x32, 0x78, 0xbc, 0xcb,
	0x95, 0x2b, 0x1a, 0x41, 0xa5, 0x7e, 0x61, 0x85, 0xaf, 0xc1, 0x94, 0x90, 0xe5, 0x60, 0x92, 0x92,
	0x0b, 0xcb, 0x7c, 0xa3, 0x92, 0x78, 0xfb, 0xd6, 0x4d, 0xae, 0x59, 0xcb, 0xbe, 0x24, 0x25, 0x9f,
	0x96, 0xda, 0x36, 0x66, 0x58, 0xf0, 0x23, 0x86, 0x41, 0xd0, 0xe5, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0xc4, 0xa2, 0xc9, 0x91, 0x01, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NotaryKey != nil {
		{
			size, err := m.NotaryKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCorda(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BaseId != nil {
		{
			size, err := m.BaseId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCorda(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintCorda(dAtA []byte, offset int, v uint64) int {
	offset -= sovCorda(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseId != nil {
		l = m.BaseId.Size()
		n += 1 + l + sovCorda(uint64(l))
	}
	if m.NotaryKey != nil {
		l = m.NotaryKey.Size()
		n += 1 + l + sovCorda(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovCorda(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCorda(x uint64) (n int) {
	return sovCorda(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCorda
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCorda
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
				return ErrInvalidLengthCorda
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCorda
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseId == nil {
				m.BaseId = &StateRef{}
			}
			if err := m.BaseId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotaryKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCorda
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
				return ErrInvalidLengthCorda
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCorda
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NotaryKey == nil {
				m.NotaryKey = &PublicKey{}
			}
			if err := m.NotaryKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCorda(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCorda
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCorda
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCorda(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCorda
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCorda
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCorda(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCorda
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
func skipCorda(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCorda
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
					return 0, ErrIntOverflowCorda
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
					return 0, ErrIntOverflowCorda
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
				return 0, ErrInvalidLengthCorda
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCorda
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCorda
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCorda        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCorda          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCorda = fmt.Errorf("proto: unexpected end of group")
)
