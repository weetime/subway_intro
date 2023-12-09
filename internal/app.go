package internal

import (
	"context"
	"os"

	"subway_intro/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/eureka/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewConfig, NewLogger)

func NewConfig(path string) (*conf.Bootstrap, error) {

	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return nil, err
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, err
	}

	NewTrace(&bc) // todo 不知道放哪更合适
	return &bc, nil
}

func NewLogger(config *conf.Bootstrap) log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", config.App.Id,
		"service.name", config.App.Name,
		"service.version", config.App.Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

func NewRegistrar(c *conf.Bootstrap, ctx context.Context) (*eureka.Registry, error) {
	r, err := eureka.New(c.Eureka.Url, eureka.WithEurekaPath("eureka"))
	return r, err
}

// set trace provider
func NewTrace(c *conf.Bootstrap) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Trace.Endpoint)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(c.App.Name),
			attribute.String("env", c.App.Env),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
