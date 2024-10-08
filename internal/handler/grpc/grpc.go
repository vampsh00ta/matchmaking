package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"matchmaking/internal/app/service"
	"matchmaking/internal/handler/grpc/pb"
)

type serverAPI struct {
	pb.UnimplementedMatchmakingServer
	l     *zap.SugaredLogger
	Match service.Match
}

func New(matchmaking service.Match, l *zap.SugaredLogger) *serverAPI {
	return &serverAPI{Match: matchmaking, l: l}
}
func Register(s *grpc.Server, matchmaking *serverAPI) {
	pb.RegisterMatchmakingServer(s, matchmaking)

}

func (s *serverAPI) FindMatch(ctx context.Context, in *pb.FindMatchRequest) (*pb.FindMatchResponse, error) {
	methodName := "FindMatch"

	if in.TgID == 0 {

		s.l.Error(methodName, status.Error(codes.InvalidArgument, "tg_id is required"))
		return nil, status.Error(codes.InvalidArgument, "tg_id is required")
	}

	foundTgID, err := s.Match.Find(ctx, int(in.TgID))
	if err != nil {
		s.l.Error(methodName, zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.l.Info(methodName, zap.String("status", "ok"))

	return &pb.FindMatchResponse{TgID: int64(foundTgID)}, nil
}
func (s *serverAPI) MatchResult(ctx context.Context, in *pb.MatchResultRequest) (*pb.MatchResultResponse, error) {
	methodName := "MatchResult"

	if in.TgIDLoser == 0 || in.TgIDWinner == 0 {
		s.l.Error(methodName, status.Error(codes.InvalidArgument, "tg_id is required"))

		return nil, status.Error(codes.InvalidArgument, "tg_id is required")
	}

	if err := s.Match.Result(ctx, int(in.TgIDWinner), int(in.TgIDLoser)); err != nil {
		s.l.Error(methodName, zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.l.Info(methodName, zap.String("status", "ok"))

	return &pb.MatchResultResponse{Ok: true}, nil
}
