package utils

import m "github.com/ozoncp/ocp-timeline-api/internal/models"

func ChunkTimeline(entities []m.Timeline, maxSizeOfChunk int) [][]m.Timeline {
	var result [][]m.Timeline

	length := len(entities)

	if length <= maxSizeOfChunk {
		return append(result, entities)
	}

	quotient, reminder := length/maxSizeOfChunk, length%maxSizeOfChunk

	for i := 0; i < quotient; i++ {
		result = append(result, entities[i*maxSizeOfChunk:i*maxSizeOfChunk+maxSizeOfChunk])
	}

	if reminder != 0 {
		result = append(result, entities[quotient*maxSizeOfChunk:])
	}

	return result
}
