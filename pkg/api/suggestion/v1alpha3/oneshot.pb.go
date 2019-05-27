// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oneshot.proto

/*
Package api_v1_alpha3 is a generated protocol buffer package.

It is generated from these files:
	oneshot.proto

It has these top-level messages:
	GetOneshotSuggestionRequest
	GetOneshotSuggestionReply
*/
package api_v1_alpha3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GetOneshotSuggestionRequest struct {
}

func (m *GetOneshotSuggestionRequest) Reset()                    { *m = GetOneshotSuggestionRequest{} }
func (m *GetOneshotSuggestionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOneshotSuggestionRequest) ProtoMessage()               {}
func (*GetOneshotSuggestionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetOneshotSuggestionReply struct {
	OnnxModel []byte `protobuf:"bytes,1,opt,name=onnx_model,json=onnxModel,proto3" json:"onnx_model,omitempty"`
}

func (m *GetOneshotSuggestionReply) Reset()                    { *m = GetOneshotSuggestionReply{} }
func (m *GetOneshotSuggestionReply) String() string            { return proto.CompactTextString(m) }
func (*GetOneshotSuggestionReply) ProtoMessage()               {}
func (*GetOneshotSuggestionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetOneshotSuggestionReply) GetOnnxModel() []byte {
	if m != nil {
		return m.OnnxModel
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOneshotSuggestionRequest)(nil), "api.v1.alpha3.GetOneshotSuggestionRequest")
	proto.RegisterType((*GetOneshotSuggestionReply)(nil), "api.v1.alpha3.GetOneshotSuggestionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OneshotSuggestion service

type OneshotSuggestionClient interface {
	GetSuggestions(ctx context.Context, in *GetOneshotSuggestionRequest, opts ...grpc.CallOption) (OneshotSuggestion_GetSuggestionsClient, error)
}

type oneshotSuggestionClient struct {
	cc *grpc.ClientConn
}

func NewOneshotSuggestionClient(cc *grpc.ClientConn) OneshotSuggestionClient {
	return &oneshotSuggestionClient{cc}
}

func (c *oneshotSuggestionClient) GetSuggestions(ctx context.Context, in *GetOneshotSuggestionRequest, opts ...grpc.CallOption) (OneshotSuggestion_GetSuggestionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_OneshotSuggestion_serviceDesc.Streams[0], c.cc, "/api.v1.alpha3.OneshotSuggestion/GetSuggestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &oneshotSuggestionGetSuggestionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OneshotSuggestion_GetSuggestionsClient interface {
	Recv() (*GetOneshotSuggestionReply, error)
	grpc.ClientStream
}

type oneshotSuggestionGetSuggestionsClient struct {
	grpc.ClientStream
}

func (x *oneshotSuggestionGetSuggestionsClient) Recv() (*GetOneshotSuggestionReply, error) {
	m := new(GetOneshotSuggestionReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for OneshotSuggestion service

type OneshotSuggestionServer interface {
	GetSuggestions(*GetOneshotSuggestionRequest, OneshotSuggestion_GetSuggestionsServer) error
}

func RegisterOneshotSuggestionServer(s *grpc.Server, srv OneshotSuggestionServer) {
	s.RegisterService(&_OneshotSuggestion_serviceDesc, srv)
}

func _OneshotSuggestion_GetSuggestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOneshotSuggestionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OneshotSuggestionServer).GetSuggestions(m, &oneshotSuggestionGetSuggestionsServer{stream})
}

type OneshotSuggestion_GetSuggestionsServer interface {
	Send(*GetOneshotSuggestionReply) error
	grpc.ServerStream
}

type oneshotSuggestionGetSuggestionsServer struct {
	grpc.ServerStream
}

func (x *oneshotSuggestionGetSuggestionsServer) Send(m *GetOneshotSuggestionReply) error {
	return x.ServerStream.SendMsg(m)
}

var _OneshotSuggestion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.alpha3.OneshotSuggestion",
	HandlerType: (*OneshotSuggestionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSuggestions",
			Handler:       _OneshotSuggestion_GetSuggestions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "oneshot.proto",
}

func init() { proto.RegisterFile("oneshot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcf, 0x4b, 0x2d,
	0xce, 0xc8, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4, 0x2b,
	0x33, 0xd4, 0x4b, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x96, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0x28, 0x56, 0x92, 0xe5, 0x92, 0x76, 0x4f, 0x2d, 0xf1, 0x87, 0x18, 0x10, 0x5c, 0x9a, 0x9e,
	0x9e, 0x5a, 0x0c, 0x92, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xb2, 0xe2, 0x92, 0xc4,
	0x2e, 0x5d, 0x90, 0x53, 0x29, 0x24, 0xcb, 0xc5, 0x95, 0x9f, 0x97, 0x57, 0x11, 0x9f, 0x9b, 0x9f,
	0x92, 0x9a, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x09, 0x12, 0xf1, 0x05, 0x09, 0x18,
	0xd5, 0x72, 0x09, 0x62, 0x68, 0x14, 0xca, 0xe0, 0xe2, 0x73, 0x4f, 0x45, 0x12, 0x28, 0x16, 0xd2,
	0xd2, 0x43, 0x71, 0xaf, 0x1e, 0x1e, 0xe7, 0x48, 0x69, 0x10, 0xa5, 0xb6, 0x20, 0xa7, 0xd2, 0x80,
	0x31, 0x89, 0x0d, 0xec, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x33, 0x0a, 0xd1,
	0x1e, 0x01, 0x00, 0x00,
}
