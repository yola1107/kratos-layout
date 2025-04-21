package server

import (
	"github.com/yola1107/kratos/v2/log"
	"github.com/yola1107/kratos/v2/middleware/recovery"
	"github.com/yola1107/kratos/v2/transport/websocket"

	v1 "kratos-layout/api/helloworld/v1"
	"kratos-layout/internal/conf"
	"kratos-layout/internal/service"
)

// NewWebsocketServer new an Websocket server.
func NewWebsocketServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *websocket.Server {
	var opts = []websocket.ServerOption{
		websocket.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Tcp.Network != "" {
		opts = append(opts, websocket.Network(c.Websocket.Network))
	}
	if c.Tcp.Addr != "" {
		opts = append(opts, websocket.Address(c.Websocket.Addr))
	}
	if c.Tcp.Timeout != nil {
		opts = append(opts, websocket.Timeout(c.Websocket.Timeout.AsDuration()))
	}
	srv := websocket.NewServer(opts...)
	v1.RegisterGreeterWebsocketServer(srv, greeter)
	return srv
}
