package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1Problem2(inputs []string) string {
	nums := make([]int, 0)

	for _, input := range inputs {
		var firstNum, lastNum string
		for index := range input {
			backIndex := len(input) - 1 - index

			if firstNum == "" {
				if isInt(input[index]) {
					firstNum = string(input[index])
				} else if index+5 < len(input) {
					if n, ok := parseSpelledIntFromStart(input[index : index+5]); ok {
						firstNum = strconv.Itoa(n)
					}
				} else if index+4 < len(input) {
					if n, ok := parseSpelledIntFromStart(input[index : index+4]); ok {
						firstNum = strconv.Itoa(n)
					}
				} else if index+3 < len(input) {
					if n, ok := parseSpelledIntFromStart(input[index : index+3]); ok {
						firstNum = strconv.Itoa(n)
					}
				}
			}
			if lastNum == "" {
				if isInt(input[backIndex]) {
					lastNum = string(input[backIndex])
				} else if backIndex-5 >= 0 {
					if n, ok := parseSpelledIntFromBack(input[backIndex-5 : backIndex+1]); ok {
						lastNum = strconv.Itoa(n)
					}
				} else if backIndex-4 >= 0 {
					if n, ok := parseSpelledIntFromBack(input[backIndex-4 : backIndex+1]); ok {
						lastNum = strconv.Itoa(n)
					}
				} else if backIndex-3 >= 0 {
					if n, ok := parseSpelledIntFromBack(input[backIndex-3 : backIndex+1]); ok {
						lastNum = strconv.Itoa(n)
					}
				}
			}
			if firstNum != "" && lastNum != "" {
				break
			}
		}
		fmt.Printf("%s: %s+%s\n", input, firstNum, lastNum)
		num, err := strconv.Atoi(firstNum + lastNum)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return strconv.Itoa(sum(nums))
}

func parseSpelledIntFromStart(str string) (int, bool) {
	if strings.HasPrefix(str, "one") {
		return 1, true
	}

	if strings.HasPrefix(str, "two") {
		return 2, true
	}

	if strings.HasPrefix(str, "three") {
		return 3, true
	}

	if strings.HasPrefix(str, "four") {
		return 4, true
	}

	if strings.HasPrefix(str, "five") {
		return 5, true
	}

	if strings.HasPrefix(str, "six") {
		return 6, true
	}

	if strings.HasPrefix(str, "seven") {
		return 7, true
	}

	if strings.HasPrefix(str, "eight") {
		return 8, true
	}

	if strings.HasPrefix(str, "nine") {
		return 9, true
	}

	return 0, false
}

func parseSpelledIntFromBack(str string) (int, bool) {
	// fmt.Printf("Parsing from back: %s\n", str)
	if strings.HasSuffix(str, "one") {
		return 1, true
	}

	if strings.HasSuffix(str, "two") {
		return 2, true
	}

	if strings.HasSuffix(str, "three") {
		return 3, true
	}

	if strings.HasSuffix(str, "four") {
		return 4, true
	}

	if strings.HasSuffix(str, "five") {
		return 5, true
	}

	if strings.HasSuffix(str, "six") {
		return 6, true
	}

	if strings.HasSuffix(str, "seven") {
		return 7, true
	}

	if strings.HasSuffix(str, "eight") {
		return 8, true
	}

	if strings.HasSuffix(str, "nine") {
		return 9, true
	}

	return 0, false
}
