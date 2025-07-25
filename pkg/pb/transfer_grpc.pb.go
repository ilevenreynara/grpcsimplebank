// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/transfer.proto

package pb

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
	TransferService_CreateTransfer_FullMethodName           = "/transfer.TransferService/CreateTransfer"
	TransferService_GetTransferByID_FullMethodName          = "/transfer.TransferService/GetTransferByID"
	TransferService_ListTransfersByAccountID_FullMethodName = "/transfer.TransferService/ListTransfersByAccountID"
)

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	CreateTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	GetTransferByID(ctx context.Context, in *GetTransferByIDRequest, opts ...grpc.CallOption) (*Transfer, error)
	ListTransfersByAccountID(ctx context.Context, in *ListTransfersByAccountIDRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Transfer], error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) CreateTransfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferService_CreateTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) GetTransferByID(ctx context.Context, in *GetTransferByIDRequest, opts ...grpc.CallOption) (*Transfer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transfer)
	err := c.cc.Invoke(ctx, TransferService_GetTransferByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) ListTransfersByAccountID(ctx context.Context, in *ListTransfersByAccountIDRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Transfer], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TransferService_ServiceDesc.Streams[0], TransferService_ListTransfersByAccountID_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListTransfersByAccountIDRequest, Transfer]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransferService_ListTransfersByAccountIDClient = grpc.ServerStreamingClient[Transfer]

// TransferServiceServer is the server API for TransferService service.
// All implementations must embed UnimplementedTransferServiceServer
// for forward compatibility.
type TransferServiceServer interface {
	CreateTransfer(context.Context, *TransferRequest) (*TransferResponse, error)
	GetTransferByID(context.Context, *GetTransferByIDRequest) (*Transfer, error)
	ListTransfersByAccountID(*ListTransfersByAccountIDRequest, grpc.ServerStreamingServer[Transfer]) error
	mustEmbedUnimplementedTransferServiceServer()
}

// UnimplementedTransferServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferServiceServer struct{}

func (UnimplementedTransferServiceServer) CreateTransfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedTransferServiceServer) GetTransferByID(context.Context, *GetTransferByIDRequest) (*Transfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferByID not implemented")
}
func (UnimplementedTransferServiceServer) ListTransfersByAccountID(*ListTransfersByAccountIDRequest, grpc.ServerStreamingServer[Transfer]) error {
	return status.Errorf(codes.Unimplemented, "method ListTransfersByAccountID not implemented")
}
func (UnimplementedTransferServiceServer) mustEmbedUnimplementedTransferServiceServer() {}
func (UnimplementedTransferServiceServer) testEmbeddedByValue()                         {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransferServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_CreateTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).CreateTransfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_GetTransferByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).GetTransferByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_GetTransferByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).GetTransferByID(ctx, req.(*GetTransferByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_ListTransfersByAccountID_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTransfersByAccountIDRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransferServiceServer).ListTransfersByAccountID(m, &grpc.GenericServerStream[ListTransfersByAccountIDRequest, Transfer]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransferService_ListTransfersByAccountIDServer = grpc.ServerStreamingServer[Transfer]

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transfer.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransfer",
			Handler:    _TransferService_CreateTransfer_Handler,
		},
		{
			MethodName: "GetTransferByID",
			Handler:    _TransferService_GetTransferByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTransfersByAccountID",
			Handler:       _TransferService_ListTransfersByAccountID_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/transfer.proto",
}
