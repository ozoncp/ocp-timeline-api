package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"testing"
)

func TestString(t *testing.T) {
	cases := []struct {
		input    Timeline
		expected string
	}{
		{Timeline{Id: 1, UserId: 2, Type: 3, From: timestamp.Timestamp{Seconds: 100}, To: timestamp.Timestamp{Seconds: 200}}, "Id: 1; UserId: 2, Type: 3; From: 1970-01-01T00:01:40Z; To: 1970-01-01T00:03:20Z"},
	}

	for i := range cases {
		actual := cases[i].input.String()

		if actual != cases[i].expected {
			t.Fatalf("Not correct actual result; actual: %v; expected: %v", actual, cases[i].expected)
		}
	}
}
