package server

import (
	v1 "kratos-layout/api/helloworld/v1"
	"kratos-layout/internal/conf"
	"kratos-layout/internal/service"

	"github.com/yola1107/kratos/v2/log"
	"github.com/yola1107/kratos/v2/middleware/recovery"
	"github.com/yola1107/kratos/v2/transport/tcp"
)

// NewTCPServer new an TCP server.
func NewTCPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *tcp.Server {
	var opts = []tcp.ServerOption{
		tcp.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Tcp.Network != "" {
		opts = append(opts, tcp.Network(c.Tcp.Network))
	}
	if c.Tcp.Addr != "" {
		opts = append(opts, tcp.Address(c.Tcp.Addr))
	}
	if c.Tcp.Timeout != nil {
		opts = append(opts, tcp.Timeout(c.Tcp.Timeout.AsDuration()))
	}
	srv := tcp.NewServer(opts...)
	v1.RegisterGreeterTCPServer(srv, greeter)
	return srv
}
