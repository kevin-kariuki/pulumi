// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LogSeverity is the severity level of a log message.  Errors are fatal; all others are informational.
type LogSeverity int32

const (
	LogSeverity_DEBUG   LogSeverity = 0
	LogSeverity_INFO    LogSeverity = 1
	LogSeverity_WARNING LogSeverity = 2
	LogSeverity_ERROR   LogSeverity = 3
)

var LogSeverity_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var LogSeverity_value = map[string]int32{
	"DEBUG":   0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x LogSeverity) String() string {
	return proto.EnumName(LogSeverity_name, int32(x))
}
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_engine_fba9c4feb5d1b1dc, []int{0}
}

type LogRequest struct {
	// the logging level of this message.
	Severity LogSeverity `protobuf:"varint,1,opt,name=severity,enum=pulumirpc.LogSeverity" json:"severity,omitempty"`
	// the contents of the logged message.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// the (optional) resource urn this log is associated with.
	Urn string `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	// the (optional) stream id that a stream of log messages can be associated with. This allows
	// clients to not have to buffer a large set of log messages that they all want to be
	// conceptually connected.  Instead the messages can be sent as chunks (with the same stream id)
	// and the end display can show the messages as they arrive, while still stitching them together
	// into one total log message.
	//
	// 0/not-given means: do not associate with any stream.
	StreamId             int32    `protobuf:"varint,4,opt,name=streamId" json:"streamId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_fba9c4feb5d1b1dc, []int{0}
}
func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (dst *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(dst, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetSeverity() LogSeverity {
	if m != nil {
		return m.Severity
	}
	return LogSeverity_DEBUG
}

func (m *LogRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *LogRequest) GetStreamId() int32 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func init() {
	proto.RegisterType((*LogRequest)(nil), "pulumirpc.LogRequest")
	proto.RegisterEnum("pulumirpc.LogSeverity", LogSeverity_name, LogSeverity_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.Engine/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(context.Context, *LogRequest) (*empty.Empty, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Engine/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Engine_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor_engine_fba9c4feb5d1b1dc) }

var fileDescriptor_engine_fba9c4feb5d1b1dc = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcd, 0x4b, 0xcf,
	0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x28, 0xcd, 0x29, 0xcd, 0xcd,
	0x2c, 0x2a, 0x48, 0x96, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x24, 0x95,
	0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0xd4, 0x29, 0x75, 0x30, 0x72, 0x71, 0xf9, 0xe4,
	0xa7, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x71, 0x71, 0x14, 0xa7, 0x96, 0xa5,
	0x16, 0x65, 0x96, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x89, 0xe9, 0xc1, 0x4d, 0xd2,
	0x03, 0x2a, 0x0c, 0x86, 0xca, 0x06, 0xc1, 0xd5, 0x09, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x01, 0xb5, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x02, 0x5c, 0xcc, 0xa5,
	0x45, 0x79, 0x12, 0xcc, 0x60, 0x51, 0x10, 0x53, 0x48, 0x0a, 0x68, 0x7e, 0x49, 0x51, 0x6a, 0x62,
	0xae, 0x67, 0x8a, 0x04, 0x0b, 0x50, 0x98, 0x35, 0x08, 0xce, 0xd7, 0xb2, 0xe2, 0xe2, 0x46, 0xb2,
	0x40, 0x88, 0x93, 0x8b, 0xd5, 0xc5, 0xd5, 0x29, 0xd4, 0x5d, 0x80, 0x41, 0x88, 0x83, 0x8b, 0xc5,
	0xd3, 0xcf, 0xcd, 0x5f, 0x80, 0x51, 0x88, 0x9b, 0x8b, 0x3d, 0xdc, 0x31, 0xc8, 0xcf, 0xd3, 0xcf,
	0x5d, 0x80, 0x09, 0xa4, 0xc2, 0x35, 0x28, 0xc8, 0x3f, 0x48, 0x80, 0xd9, 0xc8, 0x81, 0x8b, 0xcd,
	0x15, 0xec, 0x7d, 0x21, 0x33, 0x2e, 0x66, 0xa0, 0x29, 0x42, 0xa2, 0xa8, 0xce, 0x86, 0xfa, 0x4f,
	0x4a, 0x4c, 0x0f, 0x12, 0x18, 0x7a, 0xb0, 0xc0, 0xd0, 0x73, 0x05, 0x05, 0x86, 0x12, 0x43, 0x12,
	0x1b, 0x58, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x1d, 0x42, 0x33, 0x47, 0x01, 0x00,
	0x00,
}
