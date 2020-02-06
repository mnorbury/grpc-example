// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculatorpb

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

type SumRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SumStreamRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumStreamRequest) Reset()         { *m = SumStreamRequest{} }
func (m *SumStreamRequest) String() string { return proto.CompactTextString(m) }
func (*SumStreamRequest) ProtoMessage()    {}
func (*SumStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *SumStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumStreamRequest.Unmarshal(m, b)
}
func (m *SumStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumStreamRequest.Marshal(b, m, deterministic)
}
func (m *SumStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumStreamRequest.Merge(m, src)
}
func (m *SumStreamRequest) XXX_Size() int {
	return xxx_messageInfo_SumStreamRequest.Size(m)
}
func (m *SumStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumStreamRequest proto.InternalMessageInfo

func (m *SumStreamRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SumStreamResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumStreamResponse) Reset()         { *m = SumStreamResponse{} }
func (m *SumStreamResponse) String() string { return proto.CompactTextString(m) }
func (*SumStreamResponse) ProtoMessage()    {}
func (*SumStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *SumStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumStreamResponse.Unmarshal(m, b)
}
func (m *SumStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumStreamResponse.Marshal(b, m, deterministic)
}
func (m *SumStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumStreamResponse.Merge(m, src)
}
func (m *SumStreamResponse) XXX_Size() int {
	return xxx_messageInfo_SumStreamResponse.Size(m)
}
func (m *SumStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumStreamResponse proto.InternalMessageInfo

func (m *SumStreamResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FibonacciRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type FibonacciResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{5}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MaxRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxRequest) Reset()         { *m = MaxRequest{} }
func (m *MaxRequest) String() string { return proto.CompactTextString(m) }
func (*MaxRequest) ProtoMessage()    {}
func (*MaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{6}
}

func (m *MaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxRequest.Unmarshal(m, b)
}
func (m *MaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxRequest.Marshal(b, m, deterministic)
}
func (m *MaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxRequest.Merge(m, src)
}
func (m *MaxRequest) XXX_Size() int {
	return xxx_messageInfo_MaxRequest.Size(m)
}
func (m *MaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MaxRequest proto.InternalMessageInfo

func (m *MaxRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MaxResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxResponse) Reset()         { *m = MaxResponse{} }
func (m *MaxResponse) String() string { return proto.CompactTextString(m) }
func (*MaxResponse) ProtoMessage()    {}
func (*MaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{7}
}

func (m *MaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxResponse.Unmarshal(m, b)
}
func (m *MaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxResponse.Marshal(b, m, deterministic)
}
func (m *MaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxResponse.Merge(m, src)
}
func (m *MaxResponse) XXX_Size() int {
	return xxx_messageInfo_MaxResponse.Size(m)
}
func (m *MaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaxResponse proto.InternalMessageInfo

func (m *MaxResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*SumStreamRequest)(nil), "calculator.SumStreamRequest")
	proto.RegisterType((*SumStreamResponse)(nil), "calculator.SumStreamResponse")
	proto.RegisterType((*FibonacciRequest)(nil), "calculator.FibonacciRequest")
	proto.RegisterType((*FibonacciResponse)(nil), "calculator.FibonacciResponse")
	proto.RegisterType((*MaxRequest)(nil), "calculator.MaxRequest")
	proto.RegisterType((*MaxResponse)(nil), "calculator.MaxResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0x15, 0x07, 0xbe, 0x0d, 0xd9, 0x82, 0xcc, 0x31, 0x14, 0x24, 0x30, 0x28, 0x08,
	0x63, 0xe8, 0xc5, 0x83, 0x27, 0x05, 0x4f, 0xee, 0xb2, 0xdc, 0xbc, 0xbd, 0x84, 0x1c, 0x0a, 0xed,
	0x52, 0xd3, 0x64, 0xec, 0x3f, 0xf2, 0xdf, 0x94, 0xfe, 0xb0, 0x6d, 0x36, 0xd6, 0x1e, 0xbf, 0xef,
	0x7d, 0xf9, 0xb4, 0xef, 0xd3, 0xc2, 0x54, 0x62, 0x2c, 0x5d, 0x8c, 0x56, 0x9b, 0x75, 0x6a, 0xb4,
	0xd5, 0x14, 0x9a, 0x09, 0x0b, 0x01, 0xb8, 0x4b, 0x76, 0xea, 0xc7, 0xa9, 0xcc, 0xd2, 0x09, 0x10,
	0x5c, 0x90, 0x47, 0x12, 0x06, 0x3b, 0x82, 0x79, 0x12, 0x8b, 0x61, 0x99, 0x04, 0x5b, 0xc1, 0xb8,
	0x68, 0x66, 0xa9, 0xde, 0x67, 0x8a, 0xce, 0x61, 0x64, 0x54, 0xe6, 0x62, 0x5b, 0xf5, 0xab, 0xc4,
	0x42, 0x98, 0x72, 0x97, 0x70, 0x6b, 0x14, 0xd6, 0xd8, 0x5b, 0xb8, 0x3a, 0x60, 0xec, 0x54, 0x55,
	0x2d, 0x03, 0x7b, 0x82, 0x59, 0xab, 0xd9, 0x8f, 0xfd, 0x8c, 0x84, 0xde, 0xa3, 0x94, 0x51, 0x2f,
	0xb6, 0xd5, 0xec, 0xc1, 0x32, 0x80, 0x2d, 0x1e, 0xbb, 0x81, 0x2b, 0x18, 0x17, 0x9d, 0x6e, 0xd4,
	0xf3, 0xef, 0x10, 0x66, 0x1f, 0xb5, 0x58, 0xae, 0xcc, 0x21, 0x92, 0x8a, 0xbe, 0x42, 0xc0, 0x5d,
	0x42, 0xe7, 0xeb, 0xd6, 0x57, 0x68, 0x84, 0x2f, 0xef, 0xce, 0xe6, 0xe5, 0x53, 0xd8, 0x80, 0x7e,
	0xc1, 0x75, 0xad, 0x87, 0xde, 0x9f, 0xf4, 0x3c, 0xbf, 0xcb, 0x87, 0x0b, 0xdb, 0x7f, 0x56, 0x48,
	0x72, 0x5a, 0x6d, 0xc5, 0xa7, 0x9d, 0x6a, 0xf5, 0x69, 0x67, 0x2a, 0xd9, 0x60, 0x43, 0xe8, 0x1b,
	0x04, 0x5b, 0x3c, 0xfa, 0x57, 0x35, 0x1e, 0xfd, 0xab, 0x5a, 0xee, 0xf2, 0x37, 0xd9, 0x90, 0xf7,
	0x9b, 0xef, 0x49, 0xb3, 0x4f, 0x85, 0x18, 0x15, 0xbf, 0xe5, 0xcb, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0x15, 0x8d, 0x35, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Sum of two numbers - Unary operation
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Sum of a client stream - Client streaming
	SumStream(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_SumStreamClient, error)
	// Fibonacci generator - Server streaming
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error)
	// Maximum number - Bidirectional streaming
	Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) SumStream(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_SumStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/SumStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceSumStreamClient{stream}
	return x, nil
}

type CalculatorService_SumStreamClient interface {
	Send(*SumStreamRequest) error
	CloseAndRecv() (*SumStreamResponse, error)
	grpc.ClientStream
}

type calculatorServiceSumStreamClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceSumStreamClient) Send(m *SumStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceSumStreamClient) CloseAndRecv() (*SumStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Fibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_FibonacciClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type calculatorServiceFibonacciClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFibonacciClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxClient{stream}
	return x, nil
}

type CalculatorService_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Sum of two numbers - Unary operation
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Sum of a client stream - Client streaming
	SumStream(CalculatorService_SumStreamServer) error
	// Fibonacci generator - Server streaming
	Fibonacci(*FibonacciRequest, CalculatorService_FibonacciServer) error
	// Maximum number - Bidirectional streaming
	Max(CalculatorService_MaxServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) SumStream(srv CalculatorService_SumStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SumStream not implemented")
}
func (*UnimplementedCalculatorServiceServer) Fibonacci(req *FibonacciRequest, srv CalculatorService_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (*UnimplementedCalculatorServiceServer) Max(srv CalculatorService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_SumStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).SumStream(&calculatorServiceSumStreamServer{stream})
}

type CalculatorService_SumStreamServer interface {
	SendAndClose(*SumStreamResponse) error
	Recv() (*SumStreamRequest, error)
	grpc.ServerStream
}

type calculatorServiceSumStreamServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceSumStreamServer) SendAndClose(m *SumStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceSumStreamServer) Recv() (*SumStreamRequest, error) {
	m := new(SumStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Fibonacci(m, &calculatorServiceFibonacciServer{stream})
}

type CalculatorService_FibonacciServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type calculatorServiceFibonacciServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFibonacciServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Max(&calculatorServiceMaxServer{stream})
}

type CalculatorService_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SumStream",
			Handler:       _CalculatorService_SumStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Fibonacci",
			Handler:       _CalculatorService_Fibonacci_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CalculatorService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}