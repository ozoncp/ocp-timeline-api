package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-timeline-api/internal/api"
	"github.com/ozoncp/ocp-timeline-api/internal/broker"
	"github.com/ozoncp/ocp-timeline-api/internal/db"
	"github.com/ozoncp/ocp-timeline-api/internal/metrics"
	"github.com/ozoncp/ocp-timeline-api/internal/repo"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

const (
	grpcPort           = ":1025"
	grpcServerEndPoint = "localhost:1025"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpTimelineApiServer(s, api.NewServiceOcpTimeline(repo.NewRepoDb(db.Connect()), broker.NewProducer()))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	return nil
}

func runJSON(ctx context.Context) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpTimelineApiHandlerFromEndpoint(ctx, mux, grpcServerEndPoint, opts)

	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

func runMetrics() {
	metrics.RegisterMetrics()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe("192.168.1.7:9100", mux)
}

func initJaeger() (io.Closer, error) {
	cfgMetrics := &jaegerConfig.Configuration{
		ServiceName: "ocp-timeline-api",
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	tracer, closer, err := cfgMetrics.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))

	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	closer, err := initJaeger()

	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	defer cancel()

	go runMetrics()

	go runJSON(ctx)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
