package utils

import (
	"testing"
)

func TestFilterByHardcodeSliceAllCases(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"candle", "wonder", "car", "input"}, []string{"candle", "wonder", "input"}},
		{[]string{"candle", "wonder", "HardcodeWord", "input"}, []string{"candle", "wonder", "HardcodeWord", "input"}},
		{[]string{"apple", "book", "car", "pie"}, []string{}},
		{nil, []string{}},
	}

	for i := range cases {
		actual := FilterByHardcodeSlice(cases[i].input)

		if !equalSliceString(actual, cases[i].expected) {
			t.Fatalf("Expected and actual result not equal; actual: %v; expected: %v;", actual, cases[i].expected)
		}

	}
}

func equalSliceString(a, b []string) bool {
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
