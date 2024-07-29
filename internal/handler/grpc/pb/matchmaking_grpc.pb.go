// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: pb/matchmaking.proto

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

// MatchmakingClient is the client API for Matchmaking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchmakingClient interface {
	FindMatch(ctx context.Context, in *FindMatchRequest, opts ...grpc.CallOption) (*FindMatchResponse, error)
	MatchResult(ctx context.Context, in *MatchResultRequest, opts ...grpc.CallOption) (*MatchResultResponse, error)
}

type matchmakingClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchmakingClient(cc grpc.ClientConnInterface) MatchmakingClient {
	return &matchmakingClient{cc}
}

func (c *matchmakingClient) FindMatch(ctx context.Context, in *FindMatchRequest, opts ...grpc.CallOption) (*FindMatchResponse, error) {
	out := new(FindMatchResponse)
	err := c.cc.Invoke(ctx, "/matchmaking.Matchmaking/FindMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchmakingClient) MatchResult(ctx context.Context, in *MatchResultRequest, opts ...grpc.CallOption) (*MatchResultResponse, error) {
	out := new(MatchResultResponse)
	err := c.cc.Invoke(ctx, "/matchmaking.Matchmaking/MatchResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchmakingServer is the server API for Matchmaking service.
// All implementations must embed UnimplementedMatchmakingServer
// for forward compatibility
type MatchmakingServer interface {
	FindMatch(context.Context, *FindMatchRequest) (*FindMatchResponse, error)
	MatchResult(context.Context, *MatchResultRequest) (*MatchResultResponse, error)
	mustEmbedUnimplementedMatchmakingServer()
}

// UnimplementedMatchmakingServer must be embedded to have forward compatible implementations.
type UnimplementedMatchmakingServer struct {
}

func (UnimplementedMatchmakingServer) FindMatch(context.Context, *FindMatchRequest) (*FindMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMatch not implemented")
}
func (UnimplementedMatchmakingServer) MatchResult(context.Context, *MatchResultRequest) (*MatchResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchResult not implemented")
}
func (UnimplementedMatchmakingServer) mustEmbedUnimplementedMatchmakingServer() {}

// UnsafeMatchmakingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchmakingServer will
// result in compilation errors.
type UnsafeMatchmakingServer interface {
	mustEmbedUnimplementedMatchmakingServer()
}

func RegisterMatchmakingServer(s grpc.ServiceRegistrar, srv MatchmakingServer) {
	s.RegisterService(&Matchmaking_ServiceDesc, srv)
}

func _Matchmaking_FindMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchmakingServer).FindMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matchmaking.Matchmaking/FindMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchmakingServer).FindMatch(ctx, req.(*FindMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matchmaking_MatchResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchmakingServer).MatchResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matchmaking.Matchmaking/MatchResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchmakingServer).MatchResult(ctx, req.(*MatchResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Matchmaking_ServiceDesc is the grpc.ServiceDesc for Matchmaking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Matchmaking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "matchmaking.Matchmaking",
	HandlerType: (*MatchmakingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindMatch",
			Handler:    _Matchmaking_FindMatch_Handler,
		},
		{
			MethodName: "MatchResult",
			Handler:    _Matchmaking_MatchResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/matchmaking.proto",
}
