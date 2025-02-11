// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/yola1107/kratos/v2"
	"github.com/yola1107/kratos/v2/log"
	"kratos-layout/internal/biz"
	"kratos-layout/internal/conf"
	"kratos-layout/internal/data"
	"kratos-layout/internal/server"
	"kratos-layout/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	tcpServer := server.NewTCPServer(confServer, greeterService, logger)
	app := newApp(logger, grpcServer, httpServer, tcpServer)
	return app, func() {
		cleanup()
	}, nil
}
