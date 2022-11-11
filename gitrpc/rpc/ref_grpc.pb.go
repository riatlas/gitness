// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: ref.proto

package rpc

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

// ReferenceServiceClient is the client API for ReferenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReferenceServiceClient interface {
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error)
	DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*DeleteBranchResponse, error)
	ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (ReferenceService_ListBranchesClient, error)
	ListCommitTags(ctx context.Context, in *ListCommitTagsRequest, opts ...grpc.CallOption) (ReferenceService_ListCommitTagsClient, error)
}

type referenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReferenceServiceClient(cc grpc.ClientConnInterface) ReferenceServiceClient {
	return &referenceServiceClient{cc}
}

func (c *referenceServiceClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error) {
	out := new(CreateBranchResponse)
	err := c.cc.Invoke(ctx, "/rpc.ReferenceService/CreateBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceClient) DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*DeleteBranchResponse, error) {
	out := new(DeleteBranchResponse)
	err := c.cc.Invoke(ctx, "/rpc.ReferenceService/DeleteBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceClient) ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (ReferenceService_ListBranchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReferenceService_ServiceDesc.Streams[0], "/rpc.ReferenceService/ListBranches", opts...)
	if err != nil {
		return nil, err
	}
	x := &referenceServiceListBranchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReferenceService_ListBranchesClient interface {
	Recv() (*ListBranchesResponse, error)
	grpc.ClientStream
}

type referenceServiceListBranchesClient struct {
	grpc.ClientStream
}

func (x *referenceServiceListBranchesClient) Recv() (*ListBranchesResponse, error) {
	m := new(ListBranchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *referenceServiceClient) ListCommitTags(ctx context.Context, in *ListCommitTagsRequest, opts ...grpc.CallOption) (ReferenceService_ListCommitTagsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReferenceService_ServiceDesc.Streams[1], "/rpc.ReferenceService/ListCommitTags", opts...)
	if err != nil {
		return nil, err
	}
	x := &referenceServiceListCommitTagsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReferenceService_ListCommitTagsClient interface {
	Recv() (*ListCommitTagsResponse, error)
	grpc.ClientStream
}

type referenceServiceListCommitTagsClient struct {
	grpc.ClientStream
}

func (x *referenceServiceListCommitTagsClient) Recv() (*ListCommitTagsResponse, error) {
	m := new(ListCommitTagsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReferenceServiceServer is the server API for ReferenceService service.
// All implementations must embed UnimplementedReferenceServiceServer
// for forward compatibility
type ReferenceServiceServer interface {
	CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error)
	DeleteBranch(context.Context, *DeleteBranchRequest) (*DeleteBranchResponse, error)
	ListBranches(*ListBranchesRequest, ReferenceService_ListBranchesServer) error
	ListCommitTags(*ListCommitTagsRequest, ReferenceService_ListCommitTagsServer) error
	mustEmbedUnimplementedReferenceServiceServer()
}

// UnimplementedReferenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReferenceServiceServer struct {
}

func (UnimplementedReferenceServiceServer) CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (UnimplementedReferenceServiceServer) DeleteBranch(context.Context, *DeleteBranchRequest) (*DeleteBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranch not implemented")
}
func (UnimplementedReferenceServiceServer) ListBranches(*ListBranchesRequest, ReferenceService_ListBranchesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBranches not implemented")
}
func (UnimplementedReferenceServiceServer) ListCommitTags(*ListCommitTagsRequest, ReferenceService_ListCommitTagsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCommitTags not implemented")
}
func (UnimplementedReferenceServiceServer) mustEmbedUnimplementedReferenceServiceServer() {}

// UnsafeReferenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReferenceServiceServer will
// result in compilation errors.
type UnsafeReferenceServiceServer interface {
	mustEmbedUnimplementedReferenceServiceServer()
}

func RegisterReferenceServiceServer(s grpc.ServiceRegistrar, srv ReferenceServiceServer) {
	s.RegisterService(&ReferenceService_ServiceDesc, srv)
}

func _ReferenceService_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ReferenceService/CreateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceService_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ReferenceService/DeleteBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceServer).DeleteBranch(ctx, req.(*DeleteBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceService_ListBranches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBranchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReferenceServiceServer).ListBranches(m, &referenceServiceListBranchesServer{stream})
}

type ReferenceService_ListBranchesServer interface {
	Send(*ListBranchesResponse) error
	grpc.ServerStream
}

type referenceServiceListBranchesServer struct {
	grpc.ServerStream
}

func (x *referenceServiceListBranchesServer) Send(m *ListBranchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReferenceService_ListCommitTags_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCommitTagsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReferenceServiceServer).ListCommitTags(m, &referenceServiceListCommitTagsServer{stream})
}

type ReferenceService_ListCommitTagsServer interface {
	Send(*ListCommitTagsResponse) error
	grpc.ServerStream
}

type referenceServiceListCommitTagsServer struct {
	grpc.ServerStream
}

func (x *referenceServiceListCommitTagsServer) Send(m *ListCommitTagsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ReferenceService_ServiceDesc is the grpc.ServiceDesc for ReferenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReferenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ReferenceService",
	HandlerType: (*ReferenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBranch",
			Handler:    _ReferenceService_CreateBranch_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _ReferenceService_DeleteBranch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBranches",
			Handler:       _ReferenceService_ListBranches_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListCommitTags",
			Handler:       _ReferenceService_ListCommitTags_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ref.proto",
}
