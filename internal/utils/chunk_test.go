package utils

import (
	"testing"
)

func TestChunkAllCases(t *testing.T) {
	cases := []struct {
		input          []int
		expected       [][]int
		maxSizeOfButch int
	}{
		{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}}, 2},
		{[]int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}}, 2},
		{[]int{}, [][]int{{}}, 2},
		{nil, [][]int{{}}, 2},
	}

	for i := range cases {
		actual := ChunkSlice(cases[i].input, cases[i].maxSizeOfButch)

		if !equalTwoDimensionalSlices(actual, cases[i].expected) {
			t.Fatalf("Expected and actual result not equal; actual: %v; expected: %v; maxSizeOfButch: %v", actual, cases[i].expected, cases[i].maxSizeOfButch)
		}
	}
}

func equalTwoDimensionalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSliceInt(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
