package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-timeline-api/internal/api"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
	desc.RegisterOcpTimelineApiServer(s, api.NewServiceOcpTimeline())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
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

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
