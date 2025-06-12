package initialize

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer(serviceName string) func(context.Context) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://localhost:14268/api/traces"), // Jaeger UI 地址
	))
	if err != nil {
		log.Fatalf("failed to create jaeger exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
