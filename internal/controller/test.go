package controller

import (
	"context"

	pb "gitee.com/haiyux/kratos-layout/api/testing"
	"gitee.com/haiyux/kratos-layout/internal/logic"
)

type TestingService struct {
	pb.UnimplementedTestingServer
}

func NewTestingService() *TestingService {
	return &TestingService{}
}

func (s *TestingService) ListTest(ctx context.Context, req *pb.ListTestRequest) (*pb.ListTestReply, error) {

	return &pb.ListTestReply{Tests: logic.Gettests()}, nil
}
