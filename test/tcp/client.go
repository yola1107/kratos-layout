package main

import (
	"fmt"
	"time"

	pb "kratos-layout/api/helloworld/v1"

	"github.com/yola1107/kratos/v2/library/log/logrus"
	"github.com/yola1107/kratos/v2/log"
	"github.com/yola1107/kratos/v2/transport/tcp"
)

func main() {

	log.SetLogger(logrus.ShortColorLogger())

	c, err := tcp.NewTcpClient(&tcp.ClientConfig{
		Addr: "0.0.0.0:6000",
		PushHandlers: map[int32]tcp.PushMsgHandle{
			int32(pb.GameCommand_SayHelloRsp):  func(data []byte) { log.Infof("----- 1002 cb. data=%+v", data) },
			int32(pb.GameCommand_SayHello2Rsp): func(data []byte) { log.Infof("----- 1004 cb. data=%+v", data) },
		},
		RespHandlers: map[int32]tcp.RespMsgHandle{
			int32(pb.GameCommand_SayHelloReq):  func(data []byte, code int32) { log.Infof("1001 req. data=%+v code=%d", data, code) },
			int32(pb.GameCommand_SayHello2Req): func(data []byte, code int32) { log.Infof("1003 req. data=%+v code=%d", data, code) },
		},
		DisconnectFunc: func() { log.Infof("disconect.") },
		Token:          "",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// 向tcp服务器发请求
	i := 0
	for {
		req := pb.HelloRequest{Name: fmt.Sprintf("tcp_%d", i)}
		if err = c.Request(int32(pb.GameCommand_SayHelloReq), &req); err != nil {
			panic(err)
		}
		i++
		if i > 65535 {
			i = 0
		}
		time.Sleep(time.Second * 10)
	}
}
