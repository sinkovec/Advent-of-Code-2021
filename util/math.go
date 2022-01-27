package util

func Sum(input []int) int {
	result := 0
	for _, i := range input {
		result += i
	}
	return result
}

func Max(input []int) int {
	result := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > result {
			result = input[i]
		}
	}
	return result
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
