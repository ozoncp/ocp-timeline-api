package utils

import "testing"

func TestInputEmptyMap_ReturnEmptyMap(t *testing.T) {
	input := make(map[string]int, 0)

	result := RevertKeyValue(input)

	if len(result) != 0 {
		t.Fatal("result map must be empty")
	}
}

func TestInputNotEmptyMap_ReturnCorrectResult(t *testing.T) {
	input := map[string]int{
		"str1": 1,
		"str2": 2,
		"str3": 3,
	}

	result := RevertKeyValue(input)

	if len(result) != len(input) {
		t.Fatalf("result map expected length: %v, but was %v", len(input), len(result))
	}

	if firstValue, ok := result[1]; !ok || firstValue != "str1" {
		t.Fatalf("result map doe not have required key: %v; or value not equal \"str1\" - it was %v", 1, firstValue)
	}
}

func TestInputNotEmptyMapWithRepeatedValues_ReturnCorrectResult(t *testing.T) {
	input := map[string]int{
		"str1": 1,
		"str2": 2,
		"str3": 3,
		"str4": 3,
		"str5": 3,
	}

	result := RevertKeyValue(input)

	if len(result) != 3 {
		t.Fatalf("result map expected length: %v, but was %v", 3, len(result))
	}

	if firstValue, ok := result[1]; !ok || firstValue != "str1" {
		t.Fatalf("result map doe not have required key: %v; or value not equal \"str1\" - it was %v", 1, firstValue)
	}
}
