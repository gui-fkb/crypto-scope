package helper

func GenerateBoolArray(length int, value bool) []bool {
	result := make([]bool, length)
	for i := range result {
		result[i] = value
	}
	return result
}
