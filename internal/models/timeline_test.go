package models

import (
	"testing"
)

func TestString(t *testing.T) {
	cases := []struct {
		input    Timeline
		expected string
	}{
		{Timeline{Id: 1, UserId: 2, Type: 3, From: Timestamp(123), To: Timestamp(200)}, "Id: 1; UserId: 2, Type: 3; From: 1970-01-01T03:02:03+03:00; To: 1970-01-01T03:03:20+03:00"},
	}

	for i := range cases {
		actual := cases[i].input.String()

		if actual != cases[i].expected {
			t.Fatalf("Not correct actual result; actual: %v; expected: %v", actual, cases[i].expected)
		}
	}
}
