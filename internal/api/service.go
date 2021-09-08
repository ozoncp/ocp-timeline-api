package api

import (
	"github.com/ozoncp/ocp-timeline-api/internal/broker"
	"github.com/ozoncp/ocp-timeline-api/internal/repo"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
)

type serviceOcpTimeline struct {
	desc.UnimplementedOcpTimelineApiServer
	repo     repo.Repo
	producer broker.Producer
}

func NewServiceOcpTimeline(repo repo.Repo, producer broker.Producer) desc.OcpTimelineApiServer {
	return &serviceOcpTimeline{
		repo:     repo,
		producer: producer,
	}
}
