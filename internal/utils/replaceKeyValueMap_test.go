package utils

import (
	"testing"
)

func TestRevertKeyValueAllCases(t *testing.T) {
	cases := []struct {
		input    map[string]int
		expected map[int]string
	}{
		{map[string]int{}, map[int]string{}},
		{nil, map[int]string{}},
		{map[string]int{"str1": 1, "str2": 2, "str3": 3}, map[int]string{1: "str1", 2: "str2", 3: "str3"}},
		{map[string]int{"str1": 1, "str2": 2, "str3": 3, "str4": 3}, map[int]string{1: "str1", 2: "str2", 3: "str3"}},
	}

	for i := range cases {
		actual := RevertKeyValue(cases[i].input)

		if len(actual) != len(cases[i].expected) {
			t.Fatalf("Not correct length of actual and expected; actual: %v; expected: %v", actual, cases[i].expected)
		}

		for k, v := range actual {
			if cases[i].expected[k] != v {
				t.Fatalf("Not correct revevert key with values of actual and expected; actual: %v; expected: %v", actual, cases[i].expected)
			}
		}
	}
}
