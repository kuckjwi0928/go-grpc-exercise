// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: proto/board/board.proto

package board

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

// BoardClient is the client API for Board service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardClient interface {
	GetBoard(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*BoardResponse, error)
}

type boardClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardClient(cc grpc.ClientConnInterface) BoardClient {
	return &boardClient{cc}
}

func (c *boardClient) GetBoard(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*BoardResponse, error) {
	out := new(BoardResponse)
	err := c.cc.Invoke(ctx, "/Board/GetBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServer is the server API for Board service.
// All implementations must embed UnimplementedBoardServer
// for forward compatibility
type BoardServer interface {
	GetBoard(context.Context, *BoardRequest) (*BoardResponse, error)
	mustEmbedUnimplementedBoardServer()
}

// UnimplementedBoardServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServer struct {
}

func (UnimplementedBoardServer) GetBoard(context.Context, *BoardRequest) (*BoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedBoardServer) mustEmbedUnimplementedBoardServer() {}

// UnsafeBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServer will
// result in compilation errors.
type UnsafeBoardServer interface {
	mustEmbedUnimplementedBoardServer()
}

func RegisterBoardServer(s grpc.ServiceRegistrar, srv BoardServer) {
	s.RegisterService(&Board_ServiceDesc, srv)
}

func _Board_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Board/GetBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetBoard(ctx, req.(*BoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Board_ServiceDesc is the grpc.ServiceDesc for Board service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Board_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Board",
	HandlerType: (*BoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBoard",
			Handler:    _Board_GetBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/board/board.proto",
}
