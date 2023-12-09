//go:build wireinject
// +build wireinject

// go:build wireinject
package main

import (
	"context"

	"subway_intro/internal"
	"subway_intro/internal/biz"
	"subway_intro/internal/data"
	"subway_intro/internal/rpc"
	"subway_intro/internal/server"
	"subway_intro/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(configPath string, ctx context.Context) (*kratos.App, func(), error) {
	panic(wire.Build(internal.ProviderSet, biz.ProviderSet, data.ProviderSet, service.ProviderSet, rpc.ProviderSet, server.ProviderSet, newApp))
}
