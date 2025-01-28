// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/main.proto

package become_better

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BecomeBetter_MainCategories_FullMethodName = "/example.BecomeBetter/MainCategories"
	BecomeBetter_AddCategories_FullMethodName  = "/example.BecomeBetter/AddCategories"
	BecomeBetter_FillProgress_FullMethodName   = "/example.BecomeBetter/FillProgress"
)

// BecomeBetterClient is the client API for BecomeBetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BecomeBetterClient interface {
	MainCategories(ctx context.Context, in *MainCategoriesRequest, opts ...grpc.CallOption) (*MainCategoriesResponse, error)
	AddCategories(ctx context.Context, in *AddCategoryMessage, opts ...grpc.CallOption) (*MainCategoriesMessage, error)
	FillProgress(ctx context.Context, in *FillProgressRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type becomeBetterClient struct {
	cc grpc.ClientConnInterface
}

func NewBecomeBetterClient(cc grpc.ClientConnInterface) BecomeBetterClient {
	return &becomeBetterClient{cc}
}

func (c *becomeBetterClient) MainCategories(ctx context.Context, in *MainCategoriesRequest, opts ...grpc.CallOption) (*MainCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MainCategoriesResponse)
	err := c.cc.Invoke(ctx, BecomeBetter_MainCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *becomeBetterClient) AddCategories(ctx context.Context, in *AddCategoryMessage, opts ...grpc.CallOption) (*MainCategoriesMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MainCategoriesMessage)
	err := c.cc.Invoke(ctx, BecomeBetter_AddCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *becomeBetterClient) FillProgress(ctx context.Context, in *FillProgressRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BecomeBetter_FillProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BecomeBetterServer is the server API for BecomeBetter service.
// All implementations must embed UnimplementedBecomeBetterServer
// for forward compatibility.
type BecomeBetterServer interface {
	MainCategories(context.Context, *MainCategoriesRequest) (*MainCategoriesResponse, error)
	AddCategories(context.Context, *AddCategoryMessage) (*MainCategoriesMessage, error)
	FillProgress(context.Context, *FillProgressRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedBecomeBetterServer()
}

// UnimplementedBecomeBetterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBecomeBetterServer struct{}

func (UnimplementedBecomeBetterServer) MainCategories(context.Context, *MainCategoriesRequest) (*MainCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MainCategories not implemented")
}
func (UnimplementedBecomeBetterServer) AddCategories(context.Context, *AddCategoryMessage) (*MainCategoriesMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategories not implemented")
}
func (UnimplementedBecomeBetterServer) FillProgress(context.Context, *FillProgressRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FillProgress not implemented")
}
func (UnimplementedBecomeBetterServer) mustEmbedUnimplementedBecomeBetterServer() {}
func (UnimplementedBecomeBetterServer) testEmbeddedByValue()                      {}

// UnsafeBecomeBetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BecomeBetterServer will
// result in compilation errors.
type UnsafeBecomeBetterServer interface {
	mustEmbedUnimplementedBecomeBetterServer()
}

func RegisterBecomeBetterServer(s grpc.ServiceRegistrar, srv BecomeBetterServer) {
	// If the following call pancis, it indicates UnimplementedBecomeBetterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BecomeBetter_ServiceDesc, srv)
}

func _BecomeBetter_MainCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BecomeBetterServer).MainCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BecomeBetter_MainCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BecomeBetterServer).MainCategories(ctx, req.(*MainCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BecomeBetter_AddCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BecomeBetterServer).AddCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BecomeBetter_AddCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BecomeBetterServer).AddCategories(ctx, req.(*AddCategoryMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BecomeBetter_FillProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FillProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BecomeBetterServer).FillProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BecomeBetter_FillProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BecomeBetterServer).FillProgress(ctx, req.(*FillProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BecomeBetter_ServiceDesc is the grpc.ServiceDesc for BecomeBetter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BecomeBetter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.BecomeBetter",
	HandlerType: (*BecomeBetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MainCategories",
			Handler:    _BecomeBetter_MainCategories_Handler,
		},
		{
			MethodName: "AddCategories",
			Handler:    _BecomeBetter_AddCategories_Handler,
		},
		{
			MethodName: "FillProgress",
			Handler:    _BecomeBetter_FillProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}
