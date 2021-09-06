package api

import (
	"github.com/ozoncp/ocp-timeline-api/internal/repo"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
)

type serviceOcpTimeline struct {
	desc.UnimplementedOcpTimelineApiServer
	repo repo.Repo
}

func NewServiceOcpTimeline(repo repo.Repo) desc.OcpTimelineApiServer {
	return &serviceOcpTimeline{
		repo: repo,
	}
}
