package api

import (
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
)

type serviceOcpTimeline struct {
	desc.UnimplementedOcpTimelineApiServer
}

func NewServiceOcpTimeline() desc.OcpTimelineApiServer {
	return &serviceOcpTimeline{}
}
