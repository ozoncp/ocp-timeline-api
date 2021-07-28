package utils

import (
	m "github.com/ozoncp/ocp-timeline-api/internal/models"
	"testing"
)

func TestChunkTimelineAllCases(t *testing.T) {
	cases := []struct {
		input          []m.Timeline
		expected       [][]m.Timeline
		maxSizeOfChunk int
	}{
		{nil, [][]m.Timeline{{}}, 2},
		{[]m.Timeline{}, [][]m.Timeline{{}}, 2},
		{
			input:          []m.Timeline{fabricTimeline(1), fabricTimeline(2), fabricTimeline(3), fabricTimeline(4)},
			expected:       [][]m.Timeline{{fabricTimeline(1), fabricTimeline(2)}, {fabricTimeline(3), fabricTimeline(4)}},
			maxSizeOfChunk: 2,
		},
		{
			input:          []m.Timeline{fabricTimeline(1), fabricTimeline(2), fabricTimeline(3), fabricTimeline(4), fabricTimeline(5)},
			expected:       [][]m.Timeline{{fabricTimeline(1), fabricTimeline(2)}, {fabricTimeline(3), fabricTimeline(4)}, {fabricTimeline(5)}},
			maxSizeOfChunk: 2,
		},
	}

	for i := range cases {
		actual := ChunkTimeline(cases[i].input, cases[i].maxSizeOfChunk)

		if !equalTwoDimensionalTimeline(actual, cases[i].expected) {
			t.Fatalf("Not equal actual: %v and expected: %v", actual, cases[i].expected)
		}
	}
}

func equalTimeline(left *m.Timeline, right *m.Timeline) bool {
	return left.Id == right.Id && left.UserId == right.UserId && left.Type == right.Type && left.From == right.From && left.To == right.To
}

func equalTwoDimensionalTimeline(left [][]m.Timeline, right [][]m.Timeline) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {

		if len(left[i]) != len(right[i]) {
			return false
		}

		for j := range left[i] {
			if !equalTimeline(&left[i][j], &right[i][j]) {
				return false
			}
		}

	}
	return true
}

func fabricTimeline(id uint64) m.Timeline {
	return m.Timeline{
		Id:     id,
		UserId: uint64(2),
		Type:   uint64(3),
		From:   m.Timestamp(int64(200)),
		To:     m.Timestamp(int64(200)),
	}
}
