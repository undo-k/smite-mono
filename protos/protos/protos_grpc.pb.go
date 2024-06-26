// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: protos.proto

package protos

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

// GodCacheClient is the client API for GodCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GodCacheClient interface {
	FetchGod(ctx context.Context, in *GodRequest, opts ...grpc.CallOption) (*God, error)
	FetchAllGods(ctx context.Context, in *GodRequest, opts ...grpc.CallOption) (*GodList, error)
	PutGod(ctx context.Context, in *God, opts ...grpc.CallOption) (*Response, error)
}

type godCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewGodCacheClient(cc grpc.ClientConnInterface) GodCacheClient {
	return &godCacheClient{cc}
}

func (c *godCacheClient) FetchGod(ctx context.Context, in *GodRequest, opts ...grpc.CallOption) (*God, error) {
	out := new(God)
	err := c.cc.Invoke(ctx, "/protos.GodCache/FetchGod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godCacheClient) FetchAllGods(ctx context.Context, in *GodRequest, opts ...grpc.CallOption) (*GodList, error) {
	out := new(GodList)
	err := c.cc.Invoke(ctx, "/protos.GodCache/FetchAllGods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godCacheClient) PutGod(ctx context.Context, in *God, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.GodCache/PutGod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GodCacheServer is the server API for GodCache service.
// All implementations must embed UnimplementedGodCacheServer
// for forward compatibility
type GodCacheServer interface {
	FetchGod(context.Context, *GodRequest) (*God, error)
	FetchAllGods(context.Context, *GodRequest) (*GodList, error)
	PutGod(context.Context, *God) (*Response, error)
	mustEmbedUnimplementedGodCacheServer()
}

// UnimplementedGodCacheServer must be embedded to have forward compatible implementations.
type UnimplementedGodCacheServer struct {
}

func (UnimplementedGodCacheServer) FetchGod(context.Context, *GodRequest) (*God, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGod not implemented")
}
func (UnimplementedGodCacheServer) FetchAllGods(context.Context, *GodRequest) (*GodList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAllGods not implemented")
}
func (UnimplementedGodCacheServer) PutGod(context.Context, *God) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutGod not implemented")
}
func (UnimplementedGodCacheServer) mustEmbedUnimplementedGodCacheServer() {}

// UnsafeGodCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GodCacheServer will
// result in compilation errors.
type UnsafeGodCacheServer interface {
	mustEmbedUnimplementedGodCacheServer()
}

func RegisterGodCacheServer(s grpc.ServiceRegistrar, srv GodCacheServer) {
	s.RegisterService(&GodCache_ServiceDesc, srv)
}

func _GodCache_FetchGod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodCacheServer).FetchGod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GodCache/FetchGod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodCacheServer).FetchGod(ctx, req.(*GodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GodCache_FetchAllGods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodCacheServer).FetchAllGods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GodCache/FetchAllGods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodCacheServer).FetchAllGods(ctx, req.(*GodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GodCache_PutGod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(God)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodCacheServer).PutGod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GodCache/PutGod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodCacheServer).PutGod(ctx, req.(*God))
	}
	return interceptor(ctx, in, info, handler)
}

// GodCache_ServiceDesc is the grpc.ServiceDesc for GodCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GodCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GodCache",
	HandlerType: (*GodCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchGod",
			Handler:    _GodCache_FetchGod_Handler,
		},
		{
			MethodName: "FetchAllGods",
			Handler:    _GodCache_FetchAllGods_Handler,
		},
		{
			MethodName: "PutGod",
			Handler:    _GodCache_PutGod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos.proto",
}

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorClient interface {
	FetchData(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateResponse, error)
}

type aggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorClient(cc grpc.ClientConnInterface) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) FetchData(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateResponse, error) {
	out := new(AggregateResponse)
	err := c.cc.Invoke(ctx, "/protos.Aggregator/FetchData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
// All implementations must embed UnimplementedAggregatorServer
// for forward compatibility
type AggregatorServer interface {
	FetchData(context.Context, *AggregateRequest) (*AggregateResponse, error)
	mustEmbedUnimplementedAggregatorServer()
}

// UnimplementedAggregatorServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServer struct {
}

func (UnimplementedAggregatorServer) FetchData(context.Context, *AggregateRequest) (*AggregateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchData not implemented")
}
func (UnimplementedAggregatorServer) mustEmbedUnimplementedAggregatorServer() {}

// UnsafeAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServer will
// result in compilation errors.
type UnsafeAggregatorServer interface {
	mustEmbedUnimplementedAggregatorServer()
}

func RegisterAggregatorServer(s grpc.ServiceRegistrar, srv AggregatorServer) {
	s.RegisterService(&Aggregator_ServiceDesc, srv)
}

func _Aggregator_FetchData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).FetchData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Aggregator/FetchData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).FetchData(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aggregator_ServiceDesc is the grpc.ServiceDesc for Aggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchData",
			Handler:    _Aggregator_FetchData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos.proto",
}
