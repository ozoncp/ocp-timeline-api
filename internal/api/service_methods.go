package api

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
import desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *serviceOcpTimeline) CreateTimelineV1(context context.Context, req *desc.CreateTimelineV1Request) (*desc.CreateTimelineV1Response, error) {
	log.Info().Msg("create timeline was successful")

	return &desc.CreateTimelineV1Response{}, nil
}

func (a *serviceOcpTimeline) GetTimelineV1(context context.Context, req *desc.GetTimelineV1Request) (*desc.GetTimelineV1Response, error) {
	log.Info().Msg("get timeline was successful")

	return &desc.GetTimelineV1Response{}, nil
}

func (a *serviceOcpTimeline) ListTimelineV1(context context.Context, req *desc.ListTimelineV1Request) (*desc.ListTimelineV1Response, error) {
	log.Info().Msg("get list timeline was successful")

	return &desc.ListTimelineV1Response{}, nil
}

func (a *serviceOcpTimeline) RemoveTimelineV1(context context.Context, req *desc.RemoveTimelineV1Request) (*desc.RemoveTimelineV1Response, error) {
	log.Info().Msg("remove timeline was successful")

	return &desc.RemoveTimelineV1Response{}, nil
}
