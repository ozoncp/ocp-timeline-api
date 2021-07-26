package utils

import (
	"errors"
	m "github.com/ozoncp/ocp-timeline-api/internal/models"
)

func ConvertTimelineInMap(entities []m.Timeline) (map[uint64]m.Timeline, error) {
	result := make(map[uint64]m.Timeline, len(entities))

	for i := range entities {
		if _, ok := result[entities[i].Id]; !ok {
			result[entities[i].Id] = entities[i]
		} else {
			return nil, errors.New("repeated ids in input slice")
		}
	}

	return result, nil
}
