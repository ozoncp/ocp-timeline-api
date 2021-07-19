package utils

import "testing"

func TestInputEmptySplit_ReturnEmptySlice(t *testing.T) {
	input := make([]string, 0)

	result := FilterByHardcodeSlice(input)

	if len(result) != 0 {
		t.Fatalf("Result filtering from empty split must be empty split")
	}
}

func TestInputWithOneHardcodeWord_ReturnCorrectResult(t *testing.T) {
	input := []string{"candle", "wonder", "car", "input"}

	result := FilterByHardcodeSlice(input)

	if len(result) != len(input)-1 {
		t.Fatalf("Filtering not correct working; Expexted length of filterd slice is: %v, but was %v ", len(input)-1, len(result))
	}
}

func TestInputWithoutHardcodeWord_ReturnCorrectResult(t *testing.T) {
	input := []string{"candle", "wonder", "HardcodeWord", "input"}

	result := FilterByHardcodeSlice(input)

	if len(result) != len(input) {
		t.Fatalf("After filtering must be equal length spices")
	}
}

func TestInputOnlyHardcodeWord_ReturnEmptySlice(t *testing.T) {
	input := []string{"apple", "book", "car", "pie"}

	result := FilterByHardcodeSlice(input)

	if len(result) != 0 {
		t.Fatalf("After filtering must be empty spice in result, but was %v", len(result))
	}
}
