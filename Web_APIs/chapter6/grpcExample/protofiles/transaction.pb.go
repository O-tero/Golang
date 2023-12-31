//  Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package datafiles

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the 
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TransactionRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransactionRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransactionResponse struct {
	Confirmation         bool     `protobuf:"varint,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetConfirmation() bool {
	if m != nil {
		return m.Confirmation
	}
	return false
}

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "datafiles.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "datafiles.TransactionResponse")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c,
	0x49, 0x2c, 0x49, 0x4c, 0xcb, 0xcc, 0x49, 0x2d, 0x56, 0x0a, 0xe0, 0x12, 0x0a, 0x41, 0xc8, 0x07,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x15, 0xe5, 0xe7, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x7c, 0x5c, 0x4c, 0x25, 0xf9, 0x12, 0x4c, 0x60,
	0x11, 0xa6, 0x92, 0x7c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66,
	0x05, 0x46, 0x0d, 0xa6, 0x20, 0x28, 0x4f, 0xc9, 0x92, 0x4b, 0x18, 0xc5, 0xc4, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x25, 0x2e, 0x9e, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0xa2, 0xdc, 0x44, 0x90,
	0x38, 0xd8, 0x68, 0x8e, 0x20, 0x14, 0x31, 0xa3, 0x34, 0x2e, 0x01, 0xdf, 0xfc, 0xbc, 0xd4, 0x4a,
	0x24, 0xfd, 0x42, 0x41, 0x5c, 0xfc, 0xbe, 0x89, 0xd9, 0xa9, 0xc8, 0x42, 0xb2, 0x7a, 0x70, 0xf7,
	0xeb, 0x61, 0x3a, 0x5e, 0x4a, 0x0e, 0x97, 0x34, 0xc4, 0x25, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0x60,
	0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x7c, 0x5e, 0x9b, 0x1b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MoneyTransactionClient is the client API for MoneyTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoneyTransactionClient interface {
	MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type moneyTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewMoneyTransactionClient(cc grpc.ClientConnInterface) MoneyTransactionClient {
	return &moneyTransactionClient{cc}
}

func (c *moneyTransactionClient) MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/datafiles.MoneyTransaction/MakeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyTransactionServer is the server API for MoneyTransaction service.
type MoneyTransactionServer interface {
	MakeTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
}

// UnimplementedMoneyTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}

func (*UnimplementedMoneyTransactionServer) MakeTransaction(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

func RegisterMoneyTransactionServer(s *grpc.Server, srv MoneyTransactionServer) {
	s.RegisterService(&_MoneyTransaction_serviceDesc, srv)
}

func _MoneyTransaction_MakeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datafiles.MoneyTransaction/MakeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoneyTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datafiles.MoneyTransaction",
	HandlerType: (*MoneyTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeTransaction",
			Handler:    _MoneyTransaction_MakeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
