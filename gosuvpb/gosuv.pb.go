// Code generated by protoc-gen-go.
// source: gosuv.proto
// DO NOT EDIT!

/*
Package gosuvpb is a generated protocol buffer package.

It is generated from these files:
	gosuv.proto

It has these top-level messages:
	NopRequest
	Response
	Request
	TailRequest
	ProgramInfo
	ProgramStatus
	StatusResponse
	LogLine
*/
package gosuvpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NopRequest struct {
}

func (m *NopRequest) Reset()         { *m = NopRequest{} }
func (m *NopRequest) String() string { return proto.CompactTextString(m) }
func (*NopRequest) ProtoMessage()    {}

type Response struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type TailRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	Follow bool   `protobuf:"varint,3,opt,name=follow" json:"follow,omitempty"`
}

func (m *TailRequest) Reset()         { *m = TailRequest{} }
func (m *TailRequest) String() string { return proto.CompactTextString(m) }
func (*TailRequest) ProtoMessage()    {}

type ProgramInfo struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Directory string   `protobuf:"bytes,3,opt,name=directory" json:"directory,omitempty"`
	Command   []string `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	Environ   []string `protobuf:"bytes,4,rep,name=environ" json:"environ,omitempty"`
}

func (m *ProgramInfo) Reset()         { *m = ProgramInfo{} }
func (m *ProgramInfo) String() string { return proto.CompactTextString(m) }
func (*ProgramInfo) ProtoMessage()    {}

type ProgramStatus struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Extra  string `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
}

func (m *ProgramStatus) Reset()         { *m = ProgramStatus{} }
func (m *ProgramStatus) String() string { return proto.CompactTextString(m) }
func (*ProgramStatus) ProtoMessage()    {}

type StatusResponse struct {
	Programs []*ProgramStatus `protobuf:"bytes,1,rep,name=programs" json:"programs,omitempty"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}

func (m *StatusResponse) GetPrograms() []*ProgramStatus {
	if m != nil {
		return m.Programs
	}
	return nil
}

type LogLine struct {
	Line string `protobuf:"bytes,1,opt,name=line" json:"line,omitempty"`
}

func (m *LogLine) Reset()         { *m = LogLine{} }
func (m *LogLine) String() string { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GoSuv service

type GoSuvClient interface {
	Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Create(ctx context.Context, in *ProgramInfo, opts ...grpc.CallOption) (*Response, error)
}

type goSuvClient struct {
	cc *grpc.ClientConn
}

func NewGoSuvClient(cc *grpc.ClientConn) GoSuvClient {
	return &goSuvClient{cc}
}

func (c *goSuvClient) Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Create(ctx context.Context, in *ProgramInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoSuv service

type GoSuvServer interface {
	Shutdown(context.Context, *NopRequest) (*Response, error)
	Version(context.Context, *NopRequest) (*Response, error)
	Status(context.Context, *NopRequest) (*StatusResponse, error)
	Create(context.Context, *ProgramInfo) (*Response, error)
}

func RegisterGoSuvServer(s *grpc.Server, srv GoSuvServer) {
	s.RegisterService(&_GoSuv_serviceDesc, srv)
}

func _GoSuv_Shutdown_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Shutdown(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Version_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Version(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Status_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NopRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Create_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ProgramInfo)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GoSuv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.GoSuv",
	HandlerType: (*GoSuvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _GoSuv_Shutdown_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _GoSuv_Version_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _GoSuv_Status_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GoSuv_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Program service

type ProgramClient interface {
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (Program_TailClient, error)
}

type programClient struct {
	cc *grpc.ClientConn
}

func NewProgramClient(cc *grpc.ClientConn) ProgramClient {
	return &programClient{cc}
}

func (c *programClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (Program_TailClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Program_serviceDesc.Streams[0], c.cc, "/gosuvpb.Program/Tail", opts...)
	if err != nil {
		return nil, err
	}
	x := &programTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Program_TailClient interface {
	Recv() (*LogLine, error)
	grpc.ClientStream
}

type programTailClient struct {
	grpc.ClientStream
}

func (x *programTailClient) Recv() (*LogLine, error) {
	m := new(LogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Program service

type ProgramServer interface {
	Start(context.Context, *Request) (*Response, error)
	Stop(context.Context, *Request) (*Response, error)
	Tail(*TailRequest, Program_TailServer) error
}

func RegisterProgramServer(s *grpc.Server, srv ProgramServer) {
	s.RegisterService(&_Program_serviceDesc, srv)
}

func _Program_Start_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Start(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Stop_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Stop(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Tail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProgramServer).Tail(m, &programTailServer{stream})
}

type Program_TailServer interface {
	Send(*LogLine) error
	grpc.ServerStream
}

type programTailServer struct {
	grpc.ServerStream
}

func (x *programTailServer) Send(m *LogLine) error {
	return x.ServerStream.SendMsg(m)
}

var _Program_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.Program",
	HandlerType: (*ProgramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Program_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Program_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tail",
			Handler:       _Program_Tail_Handler,
			ServerStreams: true,
		},
	},
}