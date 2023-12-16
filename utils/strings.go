package utils

import "fmt"

func Indices(str string, subStr string) []int {
	indices := make([]int, 0)

	for i := 0; i < len(str); i++ {
		if i+len(subStr) <= len(str) && str[i:i+len(subStr)] == subStr {
			indices = append(indices, i)
		}
	}

	return indices
}

func CoordKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
