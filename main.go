package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/pooulad/nextjs-golang-crud-app/bootstrap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	initTracer()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	app := fiber.New()
	app.Use(otelfiber.Middleware())

	bootstrap.InitializeApp(app)
}

func initTracer() {
	ctx := context.Background()
	client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint("localhost:8080"), otlptracehttp.WithInsecure())

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %e", err)
	}

	res, err := resource.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize resource: %e", err)
	}

	// Create the trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	// Set the propagator
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(propagator)
}
