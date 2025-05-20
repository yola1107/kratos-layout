package server

import (
	"github.com/yola1107/kratos/v2/log"
	"github.com/yola1107/kratos/v2/middleware/recovery"
	"github.com/yola1107/kratos/v2/transport/grpc"
	v1 "github.com/yola1107/kratos-layout/api/helloworld/v1"
	"github.com/yola1107/kratos-layout/internal/conf"
	"github.com/yola1107/kratos-layout/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
