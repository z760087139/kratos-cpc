package service

import (
	"context"
	"kratos-project/app/helloworld/internal/biz"

	v1 "kratos-project/api/helloworld/v1"
)

// HelloWorldService is a greeter service.
type HelloWorldService struct {
	v1.UnsafeGreeterServer

	uc *biz.GreeterUsecase
}

// NewUserService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *HelloWorldService {
	return &HelloWorldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *HelloWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
