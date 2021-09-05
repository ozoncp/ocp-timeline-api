package flusher

import (
	"context"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/ozoncp/ocp-timeline-api/internal/repo"
	"github.com/ozoncp/ocp-timeline-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, entities []models.Timeline) []models.Timeline
}

func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (f *flusher) Flush(ctx context.Context, entities []models.Timeline) []models.Timeline {

	if len(entities) == 0 {
		return entities
	}

	chunks := utils.ChunkTimeline(entities, f.chunkSize)

	for i := range chunks {
		for j := range chunks[i] {
			err := f.entityRepo.AddEntities(ctx, &chunks[i][j])

			if err != nil {
				panic("repo not work")
			}
		}

	}

	return entities
}
