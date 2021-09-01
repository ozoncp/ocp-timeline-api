package main

import (
	"github.com/ozoncp/ocp-timeline-api/internal/api"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	grpcPort = ":1025"
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

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
