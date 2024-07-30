package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"matchmaking/internal/app/service"
	"matchmaking/internal/handler/grpc/pb"
)

type serverAPI struct {
	pb.UnimplementedMatchmakingServer
	matchmaking service.Matchmaking
}

func New(matchmaking service.Matchmaking) *serverAPI {
	return &serverAPI{matchmaking: matchmaking}
}
func Register(s *grpc.Server, matchmaking *serverAPI) {
	pb.RegisterMatchmakingServer(s, matchmaking)

}

func (s *serverAPI) FindMatch(ctx context.Context, in *pb.FindMatchRequest) (*pb.FindMatchResponse, error) {
	if in.TgID == 0 {
		return nil, status.Error(codes.InvalidArgument, "tg_id is required")
	}

	foundTgID, err := s.matchmaking.FindMatch(ctx, int(in.TgID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.FindMatchResponse{TgID: int64(foundTgID)}, nil
}
func (s *serverAPI) MatchResult(ctx context.Context, in *pb.MatchResultRequest) (*pb.MatchResultResponse, error) {
	if in.TgIDLoser == 0 || in.TgIDWinner == 0 {
		return nil, status.Error(codes.InvalidArgument, "tg_id is required")
	}

	if err := s.matchmaking.MatchResult(ctx, int(in.TgIDWinner), int(in.TgIDLoser)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MatchResultResponse{Ok: true}, nil
}
