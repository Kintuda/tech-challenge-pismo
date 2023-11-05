package tracer

import (
	"context"
	"os"

	"github.com/kintuda/tech-challenge-pismo/pkg/config"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer(cfg *config.ServerConfig) func(context.Context) error {
	log.Debug().Msg("starting tracing module")
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))

	if len(cfg.TracingInsecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(cfg.CollectorUrl),
		),
	)

	if err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", "tech-challenge-pismo"),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("Could not set resources")
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	log.Debug().Msg("finished tracing module initiation")

	return exporter.Shutdown
}
