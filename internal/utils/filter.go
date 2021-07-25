package utils

var hardcodeSlice = map[string]int{
	"apple": 1,
	"book":  2,
	"car":   3,
	"pie":   4,
}

func FilterByHardcodeSlice(input []string) []string {
	result := make([]string, 0, len(input))

	for _, v := range input {
		if _, ok := hardcodeSlice[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}
