//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-project/app/policy/internal/biz"
	"kratos-project/app/policy/internal/conf"
	"kratos-project/app/policy/internal/data"
	"kratos-project/app/policy/internal/server"
	"kratos-project/app/policy/internal/service"
)

// wireApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp))
}
