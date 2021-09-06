package api

import (
	"context"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
import desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *serviceOcpTimeline) CreateTimelineV1(context context.Context, req *desc.CreateTimelineV1Request) (*desc.CreateTimelineV1Response, error) {

	timeline := models.Timeline{
		UserId: req.Timeline.UserId,
		Type:   req.Timeline.Type,
		From:   *req.Timeline.From,
		To:     *req.Timeline.To,
	}

	if err := a.repo.AddEntities(context, &timeline); err != nil {
		log.Error().Err(err).Msg("error create timeline")

		return nil, err
	}

	log.Info().Msg("create timeline was successful")

	return &desc.CreateTimelineV1Response{Id: timeline.Id}, nil
}

func (a *serviceOcpTimeline) GetTimelineV1(context context.Context, req *desc.GetTimelineV1Request) (*desc.GetTimelineV1Response, error) {

	timeline, err := a.repo.DescribeEntity(context, req.Id)

	if err != nil {
		log.Error().Err(err).Msg("error get timeline")

		return nil, err
	}

	log.Info().Msg("get timeline was successful")

	return &desc.GetTimelineV1Response{
		Timeline: &desc.Timeline{
			Id:     timeline.Id,
			UserId: timeline.UserId,
			Type:   timeline.Type,
			From:   &timeline.From,
			To:     &timeline.To,
		},
	}, nil
}

func (a *serviceOcpTimeline) ListTimelineV1(context context.Context, req *desc.ListTimelineV1Request) (*desc.ListTimelineV1Response, error) {
	listTimelines, err := a.repo.ListEntities(context, req.Limit, req.Offset)
	if err != nil {
		log.Error().Err(err).Msg("error get list timeline")

		return nil, err
	}
	log.Info().Msgf("found count timelines: %v", len(listTimelines))

	response := make([]*desc.Timeline, 0, len(listTimelines))

	for i := range listTimelines {
		response = append(response, &desc.Timeline{
			Id:     listTimelines[i].Id,
			UserId: listTimelines[i].UserId,
			Type:   listTimelines[i].Type,
			From:   &listTimelines[i].From,
			To:     &listTimelines[i].To,
		})
	}

	log.Info().Msg("get list timeline was successful")

	return &desc.ListTimelineV1Response{Total: uint64(len(listTimelines)), Timelines: response}, nil
}

func (a *serviceOcpTimeline) RemoveTimelineV1(context context.Context, req *desc.RemoveTimelineV1Request) (*desc.RemoveTimelineV1Response, error) {

	if err := a.repo.RemoveEntity(context, req.Id); err != nil {
		log.Error().Err(err).Msg("error remove timeline")

		return nil, err
	}

	log.Info().Msg("remove timeline was successful")
	return &desc.RemoveTimelineV1Response{}, nil
}

func (a *serviceOcpTimeline) UpdateTimelineV1(ctx context.Context, req *desc.UpdateTimelineV1Request) (*desc.UpdateTimelineV1Response, error) {

	temp := &models.Timeline{
		Id:     req.Timeline.Id,
		UserId: req.Timeline.UserId,
		Type:   req.Timeline.Type,
		From:   *req.Timeline.From,
		To:     *req.Timeline.To,
	}

	flag, err := a.repo.UpdateEntity(ctx, temp)

	if err != nil {
		log.Error().Err(err).Msg("error update timeline")
		return nil, err
	}

	return &desc.UpdateTimelineV1Response{Updated: flag}, nil
}
