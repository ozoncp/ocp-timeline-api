package utils

import "testing"

func TestInputEmptySlice_ReturnEmptySlices(t *testing.T) {
	emptySlice := make([]int, 0)

	result := ChunkSlice(emptySlice, 3)

	if len(result) != 1 {
		t.Fatal("result slice must have one empty slice")
	}

	if len(result[0]) != 0 {
		t.Fatal("Inner slice not empty")
	}
}

func TestInputWithRemainderEqualZero_ReturnCorrectChunkedSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maxSizeChunk := 2

	result := ChunkSlice(slice, maxSizeChunk)

	if len(result) != 5 {
		t.Fatal("result slice must have 5 inners slices")
	}

	for i, v := range result {
		if len(v) != maxSizeChunk {
			t.Fatalf("inner slice with index: %v must have legth is: %v", i, maxSizeChunk)
		}
	}
}

func TestInputWithRemainderNotEqualZero_ReturnCorrectChunkedSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	maxSizeChunk := 2

	result := ChunkSlice(slice, maxSizeChunk)

	if len(result) != 6 {
		t.Fatal("result slice must have 6 inners slices")
	}

	for i, v := range result {
		if len(v) > maxSizeChunk {
			t.Fatalf("inner slice with index: %v must have legth equal or less them: %v", i, maxSizeChunk)
		}
	}
}
