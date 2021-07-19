package utils

func RevertKeyValue(inputMap map[string]int) map[int]string {
	result := map[int]string{}

	for key, value := range inputMap {
		if _, ok := result[value]; !ok {
			result[value] = key
		}
	}
	return result
}
