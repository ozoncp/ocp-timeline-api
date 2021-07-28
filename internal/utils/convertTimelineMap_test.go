package utils

import (
	m "github.com/ozoncp/ocp-timeline-api/internal/models"
	"testing"
)

func TestConvertTimelineInMapAllScenarious(t *testing.T) {
	cases := []struct {
		input    []m.Timeline
		expected map[uint64]m.Timeline
		hasError bool
	}{
		{input: []m.Timeline{
			fabricTimeline(1),
			fabricTimeline(2),
			fabricTimeline(3)},
			expected: map[uint64]m.Timeline{
				1: fabricTimeline(1),
				2: fabricTimeline(2),
				3: fabricTimeline(3),
			},
			hasError: false},
		{nil, map[uint64]m.Timeline{}, false},
		{[]m.Timeline{}, map[uint64]m.Timeline{}, false},
		{input: []m.Timeline{
			fabricTimeline(1),
			fabricTimeline(2),
			fabricTimeline(2)},
			expected: nil,
			hasError: true},
	}

	for i := range cases {
		actual, err := ConvertTimelineInMap(cases[i].input)

		if err == nil && cases[i].hasError {
			t.Fatalf("Expected error on cases actual: %v; data: %v", actual, cases[i].input)
		}

		if err != nil && !cases[i].hasError {
			t.Fatalf("Not expected error on cases data: %v", cases[i].input)
		}

		if len(actual) != len(cases[i].expected) {
			t.Fatalf("Expected and actual length not equal. Actual: %v; Expected: %v", actual, cases[i].expected)
		}

		for k := range actual {
			if cases[i].expected[k] != actual[k] {
				t.Fatalf("Not correct value for key: %v actual[k]: %v, expected[k]: %v", k, actual[k], cases[i].expected[k])
			}
		}
	}

}
