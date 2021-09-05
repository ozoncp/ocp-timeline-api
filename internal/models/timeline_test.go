package models

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	cases := []struct {
		input    Timeline
		expected string
	}{
		{Timeline{Id: 1, UserId: 2, Type: 3, From: Timestamp(time.Unix(100, 0)), To: Timestamp(time.Unix(200, 0))}, "Id: 1; UserId: 2, Type: 3; From: 1970-01-01T03:01:40+03:00; To: 1970-01-01T03:03:20+03:00"},
	}

	for i := range cases {
		actual := cases[i].input.String()

		if actual != cases[i].expected {
			t.Fatalf("Not correct actual result; actual: %v; expected: %v", actual, cases[i].expected)
		}
	}
}
