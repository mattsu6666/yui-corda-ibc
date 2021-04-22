// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/corda/v1/host.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Host struct {
	Participants []*Party                      `protobuf:"bytes,1,rep,name=participants,proto3" json:"participants,omitempty"`
	BaseId       *StateRef                     `protobuf:"bytes,2,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Notary       *Party                        `protobuf:"bytes,3,opt,name=notary,proto3" json:"notary,omitempty"`
	ClientIds    []string                      `protobuf:"bytes,4,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	ConnIds      []string                      `protobuf:"bytes,5,rep,name=conn_ids,json=connIds,proto3" json:"conn_ids,omitempty"`
	PortChanIds  []*Host_PortChannelIdentifier `protobuf:"bytes,6,rep,name=port_chan_ids,json=portChanIds,proto3" json:"port_chan_ids,omitempty"`
	BankIds      []string                      `protobuf:"bytes,7,rep,name=bank_ids,json=bankIds,proto3" json:"bank_ids,omitempty"`
	Id           string                        `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a17a943ca31e48, []int{0}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Host.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return m.Size()
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetParticipants() []*Party {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *Host) GetBaseId() *StateRef {
	if m != nil {
		return m.BaseId
	}
	return nil
}

func (m *Host) GetNotary() *Party {
	if m != nil {
		return m.Notary
	}
	return nil
}

func (m *Host) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *Host) GetConnIds() []string {
	if m != nil {
		return m.ConnIds
	}
	return nil
}

func (m *Host) GetPortChanIds() []*Host_PortChannelIdentifier {
	if m != nil {
		return m.PortChanIds
	}
	return nil
}

func (m *Host) GetBankIds() []string {
	if m != nil {
		return m.BankIds
	}
	return nil
}

func (m *Host) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Host_PortChannelIdentifier struct {
	PortId    string `protobuf:"bytes,1,opt,name=portId,proto3" json:"portId,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (m *Host_PortChannelIdentifier) Reset()         { *m = Host_PortChannelIdentifier{} }
func (m *Host_PortChannelIdentifier) String() string { return proto.CompactTextString(m) }
func (*Host_PortChannelIdentifier) ProtoMessage()    {}
func (*Host_PortChannelIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a17a943ca31e48, []int{0, 0}
}
func (m *Host_PortChannelIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Host_PortChannelIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Host_PortChannelIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Host_PortChannelIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host_PortChannelIdentifier.Merge(m, src)
}
func (m *Host_PortChannelIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *Host_PortChannelIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Host_PortChannelIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_Host_PortChannelIdentifier proto.InternalMessageInfo

func (m *Host_PortChannelIdentifier) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *Host_PortChannelIdentifier) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*Host)(nil), "ibc.lightclients.corda.v1.Host")
	proto.RegisterType((*Host_PortChannelIdentifier)(nil), "ibc.lightclients.corda.v1.Host.PortChannelIdentifier")
}

func init() {
	proto.RegisterFile("ibc/lightclients/corda/v1/host.proto", fileDescriptor_49a17a943ca31e48)
}

var fileDescriptor_49a17a943ca31e48 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x41, 0x8f, 0xd3, 0x3c,
	0x10, 0x6d, 0xda, 0xfd, 0xda, 0x8d, 0xfb, 0xc1, 0xc1, 0x12, 0xab, 0x6c, 0x81, 0x50, 0x01, 0x87,
	0x4a, 0xa8, 0x8e, 0xb6, 0x08, 0x89, 0xc3, 0x9e, 0x76, 0x17, 0x41, 0x0f, 0x48, 0x25, 0x7b, 0x82,
	0x4b, 0xe5, 0xd8, 0x6e, 0x62, 0x91, 0xb5, 0x23, 0x67, 0x5a, 0x35, 0xff, 0x02, 0xfe, 0x15, 0xc7,
	0x95, 0xb8, 0x70, 0x44, 0xed, 0x1f, 0x41, 0x76, 0x52, 0x04, 0x12, 0xed, 0xde, 0x32, 0xf3, 0xde,
	0xbc, 0x19, 0x4f, 0xde, 0xa0, 0xe7, 0x32, 0x61, 0x51, 0x2e, 0xd3, 0x0c, 0x58, 0x2e, 0x85, 0x82,
	0x32, 0x62, 0xda, 0x70, 0x1a, 0xad, 0xce, 0xa2, 0x4c, 0x97, 0x40, 0x0a, 0xa3, 0x41, 0xe3, 0x53,
	0x99, 0x30, 0xf2, 0x27, 0x8b, 0x38, 0x16, 0x59, 0x9d, 0x0d, 0x5e, 0xec, 0x17, 0x70, 0x1f, 0x63,
	0xa8, 0x0a, 0x51, 0xd6, 0x3a, 0x83, 0x87, 0xa9, 0xd6, 0x69, 0x2e, 0x22, 0x17, 0x25, 0xcb, 0x45,
	0x24, 0x6e, 0x0a, 0xa8, 0x6a, 0xf0, 0xe9, 0xf7, 0x0e, 0x3a, 0x7a, 0xa7, 0x4b, 0xc0, 0x57, 0xe8,
	0xff, 0x82, 0x1a, 0x90, 0x4c, 0x16, 0x54, 0x41, 0x19, 0x78, 0xc3, 0xce, 0xa8, 0x3f, 0x19, 0x92,
	0xbd, 0x43, 0x90, 0x19, 0x35, 0x50, 0xc5, 0x7f, 0x55, 0xe1, 0x73, 0xd4, 0x4b, 0x68, 0x29, 0xe6,
	0x92, 0x07, 0xed, 0xa1, 0x37, 0xea, 0x4f, 0x9e, 0x1d, 0x10, 0xb8, 0x06, 0x0a, 0x22, 0x16, 0x8b,
	0xb8, 0x6b, 0x6b, 0xa6, 0x1c, 0xbf, 0x46, 0x5d, 0xa5, 0x81, 0x9a, 0x2a, 0xe8, 0xb8, 0xe2, 0xbb,
	0xbb, 0x37, 0x7c, 0xfc, 0x18, 0xa1, 0x9a, 0x31, 0x97, 0xbc, 0x0c, 0x8e, 0x86, 0x9d, 0x91, 0x1f,
	0xfb, 0x75, 0x66, 0xca, 0x4b, 0x7c, 0x8a, 0x8e, 0x99, 0x56, 0xca, 0x81, 0xff, 0x39, 0xb0, 0x67,
	0x63, 0x0b, 0x7d, 0x44, 0xf7, 0x0a, 0x6d, 0x60, 0xce, 0x32, 0x5a, 0xe3, 0x5d, 0xf7, 0xf0, 0x57,
	0x07, 0x5a, 0xdb, 0x7d, 0x91, 0x99, 0x36, 0x70, 0x99, 0x51, 0xa5, 0x44, 0x3e, 0xe5, 0x42, 0x81,
	0x5c, 0x48, 0x61, 0xe2, 0x7e, 0xd1, 0xa4, 0x9b, 0xae, 0x09, 0x55, 0x9f, 0x9d, 0x6a, 0xaf, 0xee,
	0x6a, 0x63, 0x0b, 0xdd, 0x47, 0x6d, 0xc9, 0x83, 0xe3, 0xa1, 0x37, 0xf2, 0xe3, 0xb6, 0xe4, 0x83,
	0xf7, 0xe8, 0xc1, 0x3f, 0x05, 0xf1, 0x09, 0xea, 0x5a, 0xc9, 0x29, 0x0f, 0x3c, 0x47, 0x6e, 0x22,
	0xfc, 0x08, 0xf9, 0x6c, 0x47, 0x76, 0xab, 0xb6, 0xef, 0xdd, 0x25, 0x26, 0x5f, 0x3d, 0xd4, 0xb7,
	0x53, 0x5e, 0x0b, 0xb3, 0x92, 0x4c, 0xe0, 0x73, 0x84, 0x2e, 0x8d, 0xa0, 0x20, 0xdc, 0xaf, 0x3e,
	0x21, 0xb5, 0x23, 0xc8, 0xce, 0x11, 0xe4, 0x8d, 0x75, 0xc4, 0x60, 0x4f, 0x1e, 0x5f, 0x21, 0xff,
	0xc3, 0x52, 0x98, 0xea, 0x60, 0xf1, 0x93, 0x3b, 0x16, 0x76, 0x91, 0x7f, 0xdb, 0x84, 0xde, 0xed,
	0x26, 0xf4, 0x7e, 0x6e, 0x42, 0xef, 0xcb, 0x36, 0x6c, 0xdd, 0x6e, 0xc3, 0xd6, 0x8f, 0x6d, 0xd8,
	0xba, 0xf0, 0x2d, 0x3e, 0xb3, 0x72, 0x9f, 0xde, 0xa6, 0x12, 0xb2, 0x65, 0x42, 0x98, 0xbe, 0x89,
	0x38, 0x05, 0xca, 0x32, 0x2a, 0x55, 0x4e, 0x93, 0xc6, 0xd6, 0xd6, 0xf6, 0xa9, 0x8e, 0xd6, 0xd1,
	0x6f, 0xff, 0x8f, 0x77, 0x07, 0xb0, 0x5e, 0x8f, 0xeb, 0x1b, 0x70, 0xd6, 0x4f, 0xba, 0x6e, 0xbc,
	0x97, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0xc0, 0x9c, 0x27, 0x6b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostServiceClient interface {
	// transactions
	CreateHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// queries
	QueryHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Host, error)
}

type hostServiceClient struct {
	cc grpc1.ClientConn
}

func NewHostServiceClient(cc grpc1.ClientConn) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) CreateHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ibc.lightclients.corda.v1.HostService/CreateHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) QueryHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Host, error) {
	out := new(Host)
	err := c.cc.Invoke(ctx, "/ibc.lightclients.corda.v1.HostService/QueryHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
type HostServiceServer interface {
	// transactions
	CreateHost(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// queries
	QueryHost(context.Context, *emptypb.Empty) (*Host, error)
}

// UnimplementedHostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (*UnimplementedHostServiceServer) CreateHost(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHost not implemented")
}
func (*UnimplementedHostServiceServer) QueryHost(ctx context.Context, req *emptypb.Empty) (*Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHost not implemented")
}

func RegisterHostServiceServer(s grpc1.Server, srv HostServiceServer) {
	s.RegisterService(&_HostService_serviceDesc, srv)
}

func _HostService_CreateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).CreateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.lightclients.corda.v1.HostService/CreateHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).CreateHost(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_QueryHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).QueryHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.lightclients.corda.v1.HostService/QueryHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).QueryHost(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.lightclients.corda.v1.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHost",
			Handler:    _HostService_CreateHost_Handler,
		},
		{
			MethodName: "QueryHost",
			Handler:    _HostService_QueryHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/lightclients/corda/v1/host.proto",
}

func (m *Host) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Host) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Host) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BankIds) > 0 {
		for iNdEx := len(m.BankIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BankIds[iNdEx])
			copy(dAtA[i:], m.BankIds[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.BankIds[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PortChanIds) > 0 {
		for iNdEx := len(m.PortChanIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PortChanIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHost(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ConnIds) > 0 {
		for iNdEx := len(m.ConnIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ConnIds[iNdEx])
			copy(dAtA[i:], m.ConnIds[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.ConnIds[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClientIds) > 0 {
		for iNdEx := len(m.ClientIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientIds[iNdEx])
			copy(dAtA[i:], m.ClientIds[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.ClientIds[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Notary != nil {
		{
			size, err := m.Notary.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BaseId != nil {
		{
			size, err := m.BaseId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Participants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHost(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Host_PortChannelIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Host_PortChannelIdentifier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Host_PortChannelIdentifier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintHost(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintHost(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Host) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if m.BaseId != nil {
		l = m.BaseId.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if m.Notary != nil {
		l = m.Notary.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if len(m.ClientIds) > 0 {
		for _, s := range m.ClientIds {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if len(m.ConnIds) > 0 {
		for _, s := range m.ConnIds {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if len(m.PortChanIds) > 0 {
		for _, e := range m.PortChanIds {
			l = e.Size()
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if len(m.BankIds) > 0 {
		for _, s := range m.BankIds {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func (m *Host_PortChannelIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Host) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: Host: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Host: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, &Party{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notary", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Notary == nil {
				m.Notary = &Party{}
			}
			if err := m.Notary.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIds = append(m.ClientIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnIds = append(m.ConnIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortChanIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortChanIds = append(m.PortChanIds, &Host_PortChannelIdentifier{})
			if err := m.PortChanIds[len(m.PortChanIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BankIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BankIds = append(m.BankIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func (m *Host_PortChannelIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: PortChannelIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortChannelIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHost = fmt.Errorf("proto: unexpected end of group")
)
