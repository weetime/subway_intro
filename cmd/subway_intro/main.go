package main

import (
	"context"
	"flag"

	"subway_intro/internal/conf"
	"github.com/go-kratos/kratos/v2"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../config/config.dev.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	var ctx = context.Background()
	app, cleanup, err := initApp(flagconf, ctx)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(c *conf.Bootstrap, logger log.Logger, hs *http.Server, gs *grpc.Server, ctx context.Context) *kratos.App {
	return kratos.New(
		kratos.Context(ctx),
		kratos.ID(c.App.Id),
		kratos.Name(c.App.Name),
		kratos.Version(c.App.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		// kratos.Registrar(r),
	)
}
