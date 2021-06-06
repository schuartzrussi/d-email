// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: email/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateEmail struct {
	Creator               string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	From                  string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                    string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	SenderSignature       string   `protobuf:"bytes,4,opt,name=senderSignature,proto3" json:"senderSignature,omitempty"`
	SenderAddressVersion  uint64   `protobuf:"varint,5,opt,name=senderAddressVersion,proto3" json:"senderAddressVersion,omitempty"`
	Subject               string   `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                  string   `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ReplyTo               string   `protobuf:"bytes,8,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
	TrackIds              []string `protobuf:"bytes,9,rep,name=trackIds,proto3" json:"trackIds,omitempty"`
	SendedAt              string   `protobuf:"bytes,10,opt,name=sendedAt,proto3" json:"sendedAt,omitempty"`
	DecryptionKeys        []string `protobuf:"bytes,11,rep,name=decryptionKeys,proto3" json:"decryptionKeys,omitempty"`
	PreviousDecryptionKey string   `protobuf:"bytes,12,opt,name=previousDecryptionKey,proto3" json:"previousDecryptionKey,omitempty"`
	Id                    string   `protobuf:"bytes,13,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateEmail) Reset()         { *m = MsgCreateEmail{} }
func (m *MsgCreateEmail) String() string { return proto.CompactTextString(m) }
func (*MsgCreateEmail) ProtoMessage()    {}
func (*MsgCreateEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ca79922bd41532c, []int{0}
}
func (m *MsgCreateEmail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateEmail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateEmail.Merge(m, src)
}
func (m *MsgCreateEmail) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateEmail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateEmail proto.InternalMessageInfo

func (m *MsgCreateEmail) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateEmail) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgCreateEmail) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgCreateEmail) GetSenderSignature() string {
	if m != nil {
		return m.SenderSignature
	}
	return ""
}

func (m *MsgCreateEmail) GetSenderAddressVersion() uint64 {
	if m != nil {
		return m.SenderAddressVersion
	}
	return 0
}

func (m *MsgCreateEmail) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *MsgCreateEmail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgCreateEmail) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *MsgCreateEmail) GetTrackIds() []string {
	if m != nil {
		return m.TrackIds
	}
	return nil
}

func (m *MsgCreateEmail) GetSendedAt() string {
	if m != nil {
		return m.SendedAt
	}
	return ""
}

func (m *MsgCreateEmail) GetDecryptionKeys() []string {
	if m != nil {
		return m.DecryptionKeys
	}
	return nil
}

func (m *MsgCreateEmail) GetPreviousDecryptionKey() string {
	if m != nil {
		return m.PreviousDecryptionKey
	}
	return ""
}

func (m *MsgCreateEmail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgCreateEmailResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateEmailResponse) Reset()         { *m = MsgCreateEmailResponse{} }
func (m *MsgCreateEmailResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateEmailResponse) ProtoMessage()    {}
func (*MsgCreateEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ca79922bd41532c, []int{1}
}
func (m *MsgCreateEmailResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateEmailResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateEmailResponse.Merge(m, src)
}
func (m *MsgCreateEmailResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateEmailResponse proto.InternalMessageInfo

func (m *MsgCreateEmailResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateEmail)(nil), "schrsi.demail.email.MsgCreateEmail")
	proto.RegisterType((*MsgCreateEmailResponse)(nil), "schrsi.demail.email.MsgCreateEmailResponse")
}

func init() { proto.RegisterFile("email/tx.proto", fileDescriptor_8ca79922bd41532c) }

var fileDescriptor_8ca79922bd41532c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x8e, 0xda, 0x30,
	0x14, 0x25, 0x84, 0xf2, 0x30, 0x6d, 0x2a, 0xb9, 0x0f, 0x59, 0x2c, 0x22, 0x44, 0xa5, 0x36, 0x52,
	0xd5, 0x44, 0xa2, 0xfd, 0x01, 0xfa, 0x58, 0x54, 0x15, 0x9b, 0xb4, 0xea, 0x62, 0x36, 0xa3, 0x24,
	0xbe, 0x04, 0xcf, 0x40, 0x1c, 0xd9, 0xce, 0x88, 0xfc, 0xc5, 0xac, 0xe7, 0x8b, 0x66, 0xc9, 0x72,
	0x96, 0x23, 0xf8, 0x91, 0x91, 0x9d, 0x80, 0x00, 0xb1, 0x98, 0x8d, 0xe5, 0xf3, 0xf0, 0xf5, 0xd1,
	0xbd, 0x17, 0x39, 0xb0, 0x8c, 0xd8, 0x22, 0x50, 0x2b, 0x3f, 0x17, 0x5c, 0x71, 0xfc, 0x46, 0x26,
	0x73, 0x21, 0x99, 0x4f, 0x0d, 0xef, 0x9b, 0x73, 0x74, 0x67, 0x23, 0x67, 0x2a, 0xd3, 0x1f, 0x02,
	0x22, 0x05, 0xbf, 0x34, 0x85, 0x09, 0xea, 0x24, 0x1a, 0x72, 0x41, 0xac, 0xa1, 0xe5, 0xf5, 0xc2,
	0x1d, 0xc4, 0x18, 0xb5, 0x66, 0x82, 0x2f, 0x49, 0xd3, 0xd0, 0xe6, 0x8e, 0x1d, 0xd4, 0x54, 0x9c,
	0xd8, 0x86, 0x69, 0x2a, 0x8e, 0x3d, 0xf4, 0x5a, 0x42, 0x46, 0x41, 0xfc, 0x65, 0x69, 0x16, 0xa9,
	0x42, 0x00, 0x69, 0x19, 0xf1, 0x94, 0xc6, 0x63, 0xf4, 0xb6, 0xa2, 0x26, 0x94, 0x0a, 0x90, 0xf2,
	0x3f, 0x08, 0xc9, 0x78, 0x46, 0x5e, 0x0c, 0x2d, 0xaf, 0x15, 0x9e, 0xd5, 0x74, 0x36, 0x59, 0xc4,
	0x57, 0x90, 0x28, 0xd2, 0xae, 0xb2, 0xd5, 0x50, 0x67, 0x8b, 0x39, 0x2d, 0x49, 0xa7, 0xca, 0xa6,
	0xef, 0xda, 0x2d, 0x20, 0x5f, 0x94, 0xff, 0x38, 0xe9, 0x56, 0xee, 0x1a, 0xe2, 0x01, 0xea, 0x2a,
	0x11, 0x25, 0xd7, 0xbf, 0xa9, 0x24, 0xbd, 0xa1, 0xed, 0xf5, 0xc2, 0x3d, 0xd6, 0x9a, 0xf9, 0x9b,
	0x4e, 0x14, 0x41, 0xe6, 0xd9, 0x1e, 0xe3, 0x8f, 0xc8, 0xa1, 0x90, 0x88, 0x32, 0x57, 0x8c, 0x67,
	0x7f, 0xa0, 0x94, 0xa4, 0x6f, 0x5e, 0x9f, 0xb0, 0xf8, 0x1b, 0x7a, 0x97, 0x0b, 0xb8, 0x61, 0xbc,
	0x90, 0x3f, 0x0f, 0x15, 0xf2, 0xd2, 0x14, 0x3c, 0x2f, 0xea, 0x5e, 0x32, 0x4a, 0x5e, 0x55, 0xbd,
	0x64, 0x74, 0xe4, 0xa1, 0xf7, 0xc7, 0xb3, 0x09, 0x41, 0xe6, 0x3c, 0x93, 0x50, 0x3b, 0xad, 0x9d,
	0x73, 0x3c, 0x43, 0xf6, 0x54, 0xa6, 0xf8, 0x12, 0xf5, 0x0f, 0x27, 0xf9, 0xc1, 0x3f, 0x33, 0x72,
	0xff, 0xb8, 0xe4, 0xe0, 0xf3, 0x33, 0x4c, 0xbb, 0x7f, 0xbf, 0x4f, 0xee, 0x37, 0xae, 0xb5, 0xde,
	0xb8, 0xd6, 0xe3, 0xc6, 0xb5, 0x6e, 0xb7, 0x6e, 0x63, 0xbd, 0x75, 0x1b, 0x0f, 0x5b, 0xb7, 0x71,
	0xf1, 0x29, 0x65, 0x6a, 0x5e, 0xc4, 0x7e, 0xc2, 0x97, 0x41, 0x55, 0x30, 0xa0, 0x5f, 0xaa, 0x0d,
	0x5c, 0x05, 0xf5, 0x26, 0x96, 0x39, 0xc8, 0xb8, 0x6d, 0xb6, 0xf1, 0xeb, 0x53, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x12, 0x1f, 0xbf, 0x9f, 0x02, 0x00, 0x00,
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
	CreateEmail(ctx context.Context, in *MsgCreateEmail, opts ...grpc.CallOption) (*MsgCreateEmailResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateEmail(ctx context.Context, in *MsgCreateEmail, opts ...grpc.CallOption) (*MsgCreateEmailResponse, error) {
	out := new(MsgCreateEmailResponse)
	err := c.cc.Invoke(ctx, "/schrsi.demail.email.Msg/CreateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateEmail(context.Context, *MsgCreateEmail) (*MsgCreateEmailResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateEmail(ctx context.Context, req *MsgCreateEmail) (*MsgCreateEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmail not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schrsi.demail.email.Msg/CreateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateEmail(ctx, req.(*MsgCreateEmail))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schrsi.demail.email.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmail",
			Handler:    _Msg_CreateEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email/tx.proto",
}

func (m *MsgCreateEmail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateEmail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateEmail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.PreviousDecryptionKey) > 0 {
		i -= len(m.PreviousDecryptionKey)
		copy(dAtA[i:], m.PreviousDecryptionKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PreviousDecryptionKey)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.DecryptionKeys) > 0 {
		for iNdEx := len(m.DecryptionKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DecryptionKeys[iNdEx])
			copy(dAtA[i:], m.DecryptionKeys[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.DecryptionKeys[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.SendedAt) > 0 {
		i -= len(m.SendedAt)
		copy(dAtA[i:], m.SendedAt)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SendedAt)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.TrackIds) > 0 {
		for iNdEx := len(m.TrackIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TrackIds[iNdEx])
			copy(dAtA[i:], m.TrackIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.TrackIds[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ReplyTo) > 0 {
		i -= len(m.ReplyTo)
		copy(dAtA[i:], m.ReplyTo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ReplyTo)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x32
	}
	if m.SenderAddressVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SenderAddressVersion))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SenderSignature) > 0 {
		i -= len(m.SenderSignature)
		copy(dAtA[i:], m.SenderSignature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SenderSignature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgCreateEmailResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateEmailResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateEmailResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
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
func (m *MsgCreateEmail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SenderSignature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.SenderAddressVersion != 0 {
		n += 1 + sovTx(uint64(m.SenderAddressVersion))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ReplyTo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.TrackIds) > 0 {
		for _, s := range m.TrackIds {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.SendedAt)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.DecryptionKeys) > 0 {
		for _, s := range m.DecryptionKeys {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.PreviousDecryptionKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateEmailResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateEmail) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateEmail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateEmail: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderSignature", wireType)
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
			m.SenderSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderAddressVersion", wireType)
			}
			m.SenderAddressVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderAddressVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
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
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
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
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyTo", wireType)
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
			m.ReplyTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackIds", wireType)
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
			m.TrackIds = append(m.TrackIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendedAt", wireType)
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
			m.SendedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecryptionKeys", wireType)
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
			m.DecryptionKeys = append(m.DecryptionKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousDecryptionKey", wireType)
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
			m.PreviousDecryptionKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateEmailResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateEmailResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateEmailResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
