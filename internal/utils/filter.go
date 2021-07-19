package utils

var hardcodeSlice = []string{"apple", "book", "car", "pie"}

func FilterByHardcodeSlice(input []string) []string {
	result := make([]string, 0, len(input))

	for _, v := range input {
		if !contains(v) {
			result = append(result, v)
		}
	}

	return result
}

func contains(input string) bool {
	for _, v := range hardcodeSlice {
		if v == input {
			return true
		}
	}
	return false
}
