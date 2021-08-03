package repo

import "github.com/ozoncp/ocp-timeline-api/internal/models"

type Repo interface {
	AddEntities(entities []models.Timeline) error
	ListEntities(limit, offset uint64) ([]models.Timeline, error)
	DescribeEntity(entityId uint64) (*models.Timeline, error)
}
