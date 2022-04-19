// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: announcer/api.proto

package announcer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnnounceServiceClient is the client API for AnnounceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnounceServiceClient interface {
	//add Ip from ino net interface
	AddIP(ctx context.Context, in *AddIpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//remove Ip from net interface
	RemoveIP(ctx context.Context, in *RemoveIpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//get current state
	GetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStateResponse, error)
}

type announceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnounceServiceClient(cc grpc.ClientConnInterface) AnnounceServiceClient {
	return &announceServiceClient{cc}
}

func (c *announceServiceClient) AddIP(ctx context.Context, in *AddIpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/announcer.AnnounceService/AddIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announceServiceClient) RemoveIP(ctx context.Context, in *RemoveIpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/announcer.AnnounceService/RemoveIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announceServiceClient) GetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/announcer.AnnounceService/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnounceServiceServer is the server API for AnnounceService service.
// All implementations must embed UnimplementedAnnounceServiceServer
// for forward compatibility
type AnnounceServiceServer interface {
	//add Ip from ino net interface
	AddIP(context.Context, *AddIpRequest) (*emptypb.Empty, error)
	//remove Ip from net interface
	RemoveIP(context.Context, *RemoveIpRequest) (*emptypb.Empty, error)
	//get current state
	GetState(context.Context, *emptypb.Empty) (*GetStateResponse, error)
	mustEmbedUnimplementedAnnounceServiceServer()
}

// UnimplementedAnnounceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnnounceServiceServer struct {
}

func (UnimplementedAnnounceServiceServer) AddIP(context.Context, *AddIpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIP not implemented")
}
func (UnimplementedAnnounceServiceServer) RemoveIP(context.Context, *RemoveIpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIP not implemented")
}
func (UnimplementedAnnounceServiceServer) GetState(context.Context, *emptypb.Empty) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedAnnounceServiceServer) mustEmbedUnimplementedAnnounceServiceServer() {}

// UnsafeAnnounceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnounceServiceServer will
// result in compilation errors.
type UnsafeAnnounceServiceServer interface {
	mustEmbedUnimplementedAnnounceServiceServer()
}

func RegisterAnnounceServiceServer(s grpc.ServiceRegistrar, srv AnnounceServiceServer) {
	s.RegisterService(&AnnounceService_ServiceDesc, srv)
}

func _AnnounceService_AddIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnounceServiceServer).AddIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcer.AnnounceService/AddIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnounceServiceServer).AddIP(ctx, req.(*AddIpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnounceService_RemoveIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveIpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnounceServiceServer).RemoveIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcer.AnnounceService/RemoveIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnounceServiceServer).RemoveIP(ctx, req.(*RemoveIpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnounceService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnounceServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcer.AnnounceService/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnounceServiceServer).GetState(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AnnounceService_ServiceDesc is the grpc.ServiceDesc for AnnounceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnnounceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "announcer.AnnounceService",
	HandlerType: (*AnnounceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIP",
			Handler:    _AnnounceService_AddIP_Handler,
		},
		{
			MethodName: "RemoveIP",
			Handler:    _AnnounceService_RemoveIP_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _AnnounceService_GetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announcer/api.proto",
}
