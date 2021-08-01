package flusher

import (
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/ozoncp/ocp-timeline-api/internal/repo"
)

type Flusher interface {
	Flush(entities []models.Timeline) []models.Timeline
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

func (f *flusher) Flush(entities []models.Timeline) []models.Timeline {

	f.chunkOperation(entities, func(chunk []models.Timeline) {
		err := f.entityRepo.AddEntities(chunk)
		if err != nil {
			panic("repo not work")
		}
	})

	return entities
}

func (f *flusher) chunkOperation(entities []models.Timeline, operationOnChunk func([]models.Timeline)) {
	length := len(entities)

	if length == 0 {
		return
	}

	if length <= f.chunkSize {
		operationOnChunk(entities)

		return
	}

	quotient, reminder := length/f.chunkSize, length%f.chunkSize

	for i := 0; i < quotient; i++ {
		operationOnChunk(entities[i*f.chunkSize : i*f.chunkSize+f.chunkSize])
	}

	if reminder != 0 {
		operationOnChunk(entities[quotient*f.chunkSize:])
	}
}
