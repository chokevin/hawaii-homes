// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: homeserver.proto

package homeserver

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
	HomeService_GetProperty_FullMethodName = "/homeserver.HomeService/GetProperty"
)

// HomeServiceClient is the client API for HomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The real estate service definition.
type HomeServiceClient interface {
	// Gets property details by ID.
	GetProperty(ctx context.Context, in *GetPropertyRequest, opts ...grpc.CallOption) (*GetPropertyResponse, error)
}

type homeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeServiceClient(cc grpc.ClientConnInterface) HomeServiceClient {
	return &homeServiceClient{cc}
}

func (c *homeServiceClient) GetProperty(ctx context.Context, in *GetPropertyRequest, opts ...grpc.CallOption) (*GetPropertyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPropertyResponse)
	err := c.cc.Invoke(ctx, HomeService_GetProperty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServiceServer is the server API for HomeService service.
// All implementations must embed UnimplementedHomeServiceServer
// for forward compatibility.
//
// The real estate service definition.
type HomeServiceServer interface {
	// Gets property details by ID.
	GetProperty(context.Context, *GetPropertyRequest) (*GetPropertyResponse, error)
	mustEmbedUnimplementedHomeServiceServer()
}

// UnimplementedHomeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHomeServiceServer struct{}

func (UnimplementedHomeServiceServer) GetProperty(context.Context, *GetPropertyRequest) (*GetPropertyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProperty not implemented")
}
func (UnimplementedHomeServiceServer) mustEmbedUnimplementedHomeServiceServer() {}
func (UnimplementedHomeServiceServer) testEmbeddedByValue()                     {}

// UnsafeHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServiceServer will
// result in compilation errors.
type UnsafeHomeServiceServer interface {
	mustEmbedUnimplementedHomeServiceServer()
}

func RegisterHomeServiceServer(s grpc.ServiceRegistrar, srv HomeServiceServer) {
	// If the following call pancis, it indicates UnimplementedHomeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HomeService_ServiceDesc, srv)
}

func _HomeService_GetProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServiceServer).GetProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HomeService_GetProperty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServiceServer).GetProperty(ctx, req.(*GetPropertyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeService_ServiceDesc is the grpc.ServiceDesc for HomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homeserver.HomeService",
	HandlerType: (*HomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProperty",
			Handler:    _HomeService_GetProperty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homeserver.proto",
}
