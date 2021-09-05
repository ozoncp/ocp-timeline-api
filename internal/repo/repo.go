package repo

import (
	"context"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
)

type Repo interface {
	AddEntities(ctx context.Context, entity *models.Timeline) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Timeline, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Timeline, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
	UpdateEntity(ctx context.Context, timeline *models.Timeline) (bool, error)
}
