package utils

func ChunkSlice(inputSlice []int, maxSizeOfChunk int) [][]int {
	var result [][]int

	length := len(inputSlice)

	if length <= maxSizeOfChunk {
		return append(result, inputSlice)
	}

	quotient, reminder := length/maxSizeOfChunk, length%maxSizeOfChunk

	for i := 0; i < quotient; i++ {
		result = append(result, inputSlice[i*maxSizeOfChunk:i*maxSizeOfChunk+maxSizeOfChunk])
	}

	if reminder != 0 {
		result = append(result, inputSlice[quotient*maxSizeOfChunk:])
	}

	return result
}
