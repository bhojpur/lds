// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LdsUIClient is the client API for LdsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LdsUIClient interface {
	// ListDetectorSpecs returns a list of Detector(s) that can be started through the UI.
	ListDetectorSpecs(ctx context.Context, in *ListDetectorSpecsRequest, opts ...grpc.CallOption) (LdsUI_ListDetectorSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type ldsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewLdsUIClient(cc grpc.ClientConnInterface) LdsUIClient {
	return &ldsUIClient{cc}
}

func (c *ldsUIClient) ListDetectorSpecs(ctx context.Context, in *ListDetectorSpecsRequest, opts ...grpc.CallOption) (LdsUI_ListDetectorSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LdsUI_ServiceDesc.Streams[0], "/v1.LdsUI/ListDetectorSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &ldsUIListDetectorSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LdsUI_ListDetectorSpecsClient interface {
	Recv() (*ListDetectorSpecsResponse, error)
	grpc.ClientStream
}

type ldsUIListDetectorSpecsClient struct {
	grpc.ClientStream
}

func (x *ldsUIListDetectorSpecsClient) Recv() (*ListDetectorSpecsResponse, error) {
	m := new(ListDetectorSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ldsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.LdsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LdsUIServer is the server API for LdsUI service.
// All implementations must embed UnimplementedLdsUIServer
// for forward compatibility
type LdsUIServer interface {
	// ListDetectorSpecs returns a list of Detector(s) that can be started through the UI.
	ListDetectorSpecs(*ListDetectorSpecsRequest, LdsUI_ListDetectorSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedLdsUIServer()
}

// UnimplementedLdsUIServer must be embedded to have forward compatible implementations.
type UnimplementedLdsUIServer struct {
}

func (UnimplementedLdsUIServer) ListDetectorSpecs(*ListDetectorSpecsRequest, LdsUI_ListDetectorSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDetectorSpecs not implemented")
}
func (UnimplementedLdsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedLdsUIServer) mustEmbedUnimplementedLdsUIServer() {}

// UnsafeLdsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LdsUIServer will
// result in compilation errors.
type UnsafeLdsUIServer interface {
	mustEmbedUnimplementedLdsUIServer()
}

func RegisterLdsUIServer(s grpc.ServiceRegistrar, srv LdsUIServer) {
	s.RegisterService(&LdsUI_ServiceDesc, srv)
}

func _LdsUI_ListDetectorSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDetectorSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LdsUIServer).ListDetectorSpecs(m, &ldsUIListDetectorSpecsServer{stream})
}

type LdsUI_ListDetectorSpecsServer interface {
	Send(*ListDetectorSpecsResponse) error
	grpc.ServerStream
}

type ldsUIListDetectorSpecsServer struct {
	grpc.ServerStream
}

func (x *ldsUIListDetectorSpecsServer) Send(m *ListDetectorSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LdsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LdsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LdsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LdsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LdsUI_ServiceDesc is the grpc.ServiceDesc for LdsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LdsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.LdsUI",
	HandlerType: (*LdsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _LdsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListDetectorSpecs",
			Handler:       _LdsUI_ListDetectorSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lds-ui.proto",
}
