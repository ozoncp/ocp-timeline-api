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
			newTimeline(1),
			newTimeline(2),
			newTimeline(3)},
			expected: map[uint64]m.Timeline{
				1: newTimeline(1),
				2: newTimeline(2),
				3: newTimeline(3),
			},
			hasError: false},
		{nil, map[uint64]m.Timeline{}, false},
		{[]m.Timeline{}, map[uint64]m.Timeline{}, false},
		{input: []m.Timeline{
			newTimeline(1),
			newTimeline(2),
			newTimeline(2)},
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
			expected := cases[i].expected[k]
			actualCurrent := actual[k]

			if !equalTimeline(&expected, &actualCurrent) {
				t.Fatalf("Not correct value for key: %v actual[k]: %v, expected[k]: %v", k, actualCurrent, expected)
			}
		}
	}

}
