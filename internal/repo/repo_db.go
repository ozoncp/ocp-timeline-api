package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/ozoncp/ocp-timeline-api/internal/utils"
	"github.com/rs/zerolog/log"
	"time"
)

var timelineTableName = "timeline"

type repoDb struct {
	db *sqlx.DB
}

func (r *repoDb) UpdateEntity(ctx context.Context, timeline *models.Timeline) (bool, error) {
	var query = sq.Update(timelineTableName).
		Set("user_id", &timeline.UserId).
		Set("\"type_id\"", &timeline.Type).
		Set("\"from\"", timeline.From.AsTime()).
		Set("\"to\"", timeline.To.AsTime()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	traceQuerySql(query)

	exec, err := query.ExecContext(ctx)

	if err != nil {
		return false, err
	}

	if _, err := exec.RowsAffected(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *repoDb) RemoveEntity(ctx context.Context, entityId uint64) error {

	query := sq.Delete(timelineTableName).
		Where(sq.Eq{"timeline_id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	traceQuerySql(query)

	if _, err := query.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repoDb) AddEntities(ctx context.Context, entity *models.Timeline) error {
	fromTime := entity.From.AsTime()
	toTime := entity.To.AsTime()

	query := sq.Insert(timelineTableName).
		Columns("user_id", "\"type_id\"", "\"from\"", "\"to\"").
		Values(
			entity.UserId,
			entity.Type,
			fromTime,
			toTime).
		Suffix("RETURNING \"timeline_id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	traceQuerySql(query)

	err := query.QueryRowContext(ctx).Scan(&entity.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repoDb) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Timeline, error) {
	query := sq.Select("timeline_id", "user_id", "\"type_id\"", "\"from\"", "\"to\"").
		From(timelineTableName).
		OrderBy("timeline_id").Limit(limit).Offset(offset).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	traceQuerySql(query)

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var timelines []models.Timeline

	for rows.Next() {
		var from time.Time
		var to time.Time
		var timeline models.Timeline
		if err = rows.Scan(&timeline.Id,
			&timeline.UserId,
			&timeline.Type,
			&from,
			&to); err != nil {
			return nil, err
		}

		timeline.From = utils.ConvertTimeInTimeStamp(from)
		timeline.To = utils.ConvertTimeInTimeStamp(to)
		timelines = append(timelines, timeline)
	}

	return timelines, nil
}

func (r *repoDb) DescribeEntity(ctx context.Context, entityId uint64) (*models.Timeline, error) {
	query := sq.Select("timeline_id", "user_id", "\"type_id\"", "\"from\"", "\"to\"").
		From(timelineTableName).
		Where(sq.Eq{"timeline_id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	traceQuerySql(query)

	var timeline models.Timeline

	if err := query.QueryRowContext(ctx).
		Scan(&timeline.Id,
			&timeline.UserId,
			&timeline.Type,
			&timeline.From,
			&timeline.To); err != nil {
		return nil, err
	}

	return &timeline, nil
}

func NewRepoDb(db *sqlx.DB) Repo {
	return &repoDb{
		db: db,
	}
}

func traceQuerySql(builder traceSqlStatement) {
	sql, _, _ := builder.ToSql()

	log.Trace().Msg(sql)
}

type traceSqlStatement interface {
	ToSql() (string, []interface{}, error)
}
