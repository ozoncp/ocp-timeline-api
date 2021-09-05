package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	. "github.com/ozoncp/ocp-timeline-api/internal/broker"
	"github.com/ozoncp/ocp-timeline-api/internal/metrics"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/ozoncp/ocp-timeline-api/internal/utils"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (a *serviceOcpTimeline) CreateTimelineV1(context context.Context, req *desc.CreateTimelineV1Request) (*desc.CreateTimelineV1Response, error) {

	from, to, errFrom, errTo := convertTimeInTime(req.From, req.To)

	if errFrom != nil {
		return nil, errFrom
	}

	if errTo != nil {
		return nil, errTo
	}

	timeline := models.Timeline{
		UserId: req.UserId,
		Type:   req.Type,
		From:   models.Timestamp(from),
		To:     models.Timestamp(to),
	}

	if err := a.repo.AddEntities(context, &timeline); err != nil {
		log.Error().Err(err).Msg("error create timeline")

		return nil, err
	}

	log.Info().Msg("create timeline was successful")

	err := a.producer.Send(Create.String(), map[string]interface{}{"timeline": timeline})

	if err != nil {
		log.Error().Err(err).Msg("error produce in kafka")

		return nil, err
	}

	metrics.IncCreateCounter()
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
			From:   convertTimeInStr(time.Time(timeline.From)),
			To:     convertTimeInStr(time.Time(timeline.To)),
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
			From:   convertTimeInStr(time.Time(listTimelines[i].From)),
			To:     convertTimeInStr(time.Time(listTimelines[i].To)),
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

	err := a.producer.Send(Remove.String(), Message{Data: map[string]interface{}{"id": req.Id}})

	if err != nil {
		log.Error().Err(err).Msg("error produce in kafka")

		return nil, err
	}

	metrics.IncRemoveCounter()
	return &desc.RemoveTimelineV1Response{}, nil
}

func (a *serviceOcpTimeline) UpdateTimelineV1(ctx context.Context, req *desc.UpdateTimelineV1Request) (*desc.UpdateTimelineV1Response, error) {
	from, to, errFrom, errTo := convertTimeInTime(req.Timeline.From, req.Timeline.To)

	if errFrom != nil {

		log.Error().Err(errFrom).Msg("error format timeline from")
		return nil, errFrom
	}

	if errTo != nil {

		log.Error().Err(errTo).Msg("error format timeline to")
		return nil, errTo
	}

	temp := &models.Timeline{
		Id:     req.Timeline.Id,
		UserId: req.Timeline.UserId,
		Type:   req.Timeline.Type,
		From:   models.Timestamp(from),
		To:     models.Timestamp(to),
	}

	flag, err := a.repo.UpdateEntity(ctx, temp)

	if err != nil {
		log.Error().Err(err).Msg("error update timeline")
		return nil, err
	}

	log.Info().Msg("update timeline was successful")

	err = a.producer.Send(
		Update.String(),
		Message{Data: map[string]interface{}{"timeline": temp, "success": flag}},
	)

	if err != nil {
		log.Error().Err(err).Msg("error produce in kafka")

		return nil, err
	}

	metrics.IncUpdateCounter()
	return &desc.UpdateTimelineV1Response{Updated: flag}, nil
}

func (a *serviceOcpTimeline) MultiCreateTimelinesV1(ctx context.Context, req *desc.MultiCreateTimelinesV1Request) (*desc.MultiCreateTimelinesV1Response, error) {
	batchSize := 5
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateTimelinesV1")

	defer span.Finish()

	timelines := make([]models.Timeline, 0, len(req.Timelines))

	for _, t := range req.Timelines {

		from, to, errFrom, errTo := convertTimeInTime(t.From, t.To)

		if errFrom != nil {
			return nil, errFrom
		}

		if errTo != nil {
			return nil, errTo
		}

		timelines = append(timelines, models.Timeline{
			Type:   t.Type,
			UserId: t.UserId,
			From:   models.Timestamp(from),
			To:     models.Timestamp(to),
		})
	}

	batches := utils.ChunkTimeline(timelines, batchSize)
	for i := range batches {
		err := func() error {
			childSpan := tracer.StartSpan("Size of data bytes 12", opentracing.ChildOf(span.Context()))
			defer childSpan.Finish()

			_, err := a.repo.MultiCreateEntity(ctx, batches[i])

			return err
		}()

		if err != nil {
			return nil, err
		}
	}

	log.Info().Msg("multiple timelines add successful")

	return &desc.MultiCreateTimelinesV1Response{Added: true}, nil
}

func convertStrInTime(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, date)

	return t, err
}

func convertTimeInStr(date time.Time) string {

	str := date.Format(time.RFC3339)

	return str
}

func convertTimeInTime(from string, to string) (time.Time, time.Time, error, error) {
	fromTime, errFrom := convertStrInTime(from)

	toTime, errTo := convertStrInTime(to)

	return fromTime, toTime, errFrom, errTo
}
