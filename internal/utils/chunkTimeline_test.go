package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
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
			input:          []m.Timeline{newTimeline(1), newTimeline(2), newTimeline(3), newTimeline(4)},
			expected:       [][]m.Timeline{{newTimeline(1), newTimeline(2)}, {newTimeline(3), newTimeline(4)}},
			maxSizeOfChunk: 2,
		},
		{
			input:          []m.Timeline{newTimeline(1), newTimeline(2), newTimeline(3), newTimeline(4), newTimeline(5)},
			expected:       [][]m.Timeline{{newTimeline(1), newTimeline(2)}, {newTimeline(3), newTimeline(4)}, {newTimeline(5)}},
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
	return left.Id == right.Id &&
		left.UserId == right.UserId &&
		left.Type == right.Type &&
		left.From.Seconds == right.From.Seconds &&
		left.To.Seconds == right.To.Seconds
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

func newTimeline(id uint64) m.Timeline {
	return m.Timeline{
		Id:     id,
		UserId: uint64(2),
		Type:   uint64(3),
		From:   timestamp.Timestamp{Seconds: 100},
		To:     timestamp.Timestamp{Seconds: 200},
	}
}
