//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/yola1107/kratos/v2"
	"github.com/yola1107/kratos/v2/log"
	"github.com/yola1107/kratos-layout/internal/biz"
	"github.com/yola1107/kratos-layout/internal/conf"
	"github.com/yola1107/kratos-layout/internal/data"
	"github.com/yola1107/kratos-layout/internal/server"
	"github.com/yola1107/kratos-layout/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
