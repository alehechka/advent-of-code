package solutions

import (
	"strconv"
)

func Day1Problem1(inputs []string) string {
	nums := make([]int, 0)

	for _, input := range inputs {
		var firstNum, lastNum string
		for index := range input {
			backIndex := len(input) - 1 - index

			if firstNum == "" && isInt(input[index]) {
				firstNum = string(input[index])
			}
			if lastNum == "" && isInt(input[backIndex]) {
				lastNum = string(input[backIndex])
			}
			if firstNum != "" && lastNum != "" {
				break
			}
		}
		num, err := strconv.Atoi(firstNum + lastNum)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return strconv.Itoa(sum(nums))
}

func isInt(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
