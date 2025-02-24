// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: proto/blog.proto

package pb

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

const (
	Blogs_GetBlogList_FullMethodName = "/Blogs/GetBlogList"
)

// BlogsClient is the client API for Blogs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogsClient interface {
	GetBlogList(ctx context.Context, in *GetBlogListRequest, opts ...grpc.CallOption) (*GetBlogListResponse, error)
}

type blogsClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogsClient(cc grpc.ClientConnInterface) BlogsClient {
	return &blogsClient{cc}
}

func (c *blogsClient) GetBlogList(ctx context.Context, in *GetBlogListRequest, opts ...grpc.CallOption) (*GetBlogListResponse, error) {
	out := new(GetBlogListResponse)
	err := c.cc.Invoke(ctx, Blogs_GetBlogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogsServer is the server API for Blogs service.
// All implementations must embed UnimplementedBlogsServer
// for forward compatibility
type BlogsServer interface {
	GetBlogList(context.Context, *GetBlogListRequest) (*GetBlogListResponse, error)
	mustEmbedUnimplementedBlogsServer()
}

// UnimplementedBlogsServer must be embedded to have forward compatible implementations.
type UnimplementedBlogsServer struct {
}

func (UnimplementedBlogsServer) GetBlogList(context.Context, *GetBlogListRequest) (*GetBlogListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogList not implemented")
}
func (UnimplementedBlogsServer) mustEmbedUnimplementedBlogsServer() {}

// UnsafeBlogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogsServer will
// result in compilation errors.
type UnsafeBlogsServer interface {
	mustEmbedUnimplementedBlogsServer()
}

func RegisterBlogsServer(s grpc.ServiceRegistrar, srv BlogsServer) {
	s.RegisterService(&Blogs_ServiceDesc, srv)
}

func _Blogs_GetBlogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogsServer).GetBlogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blogs_GetBlogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogsServer).GetBlogList(ctx, req.(*GetBlogListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blogs_ServiceDesc is the grpc.ServiceDesc for Blogs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blogs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Blogs",
	HandlerType: (*BlogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlogList",
			Handler:    _Blogs_GetBlogList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blog.proto",
}
