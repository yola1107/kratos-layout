package service

import (
	"context"

	v1 "kratos-layout/api/helloworld/v1"
	"kratos-layout/internal/biz"

	"github.com/yola1107/kratos/v2/transport/tcp"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHelloReq implements helloworld.GreeterServer.
func (s *GreeterService) SayHelloReq(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

// SayHello2Req implements helloworld.GreeterServer.
func (s *GreeterService) SayHello2Req(ctx context.Context, in *v1.Hello2Request) (*v1.Hello2Reply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.Hello2Reply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) SetCometChan(cl *tcp.ChanList, cs *tcp.Server) {}

func (s *GreeterService) IsLoopFunc(f string) (isLoop bool) { return false }
