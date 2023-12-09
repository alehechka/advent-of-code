package twentythree

import (
	"strconv"
	"strings"
)

func Day9Problem1(inputs []string) string {
	var total int
	for _, input := range inputs {
		pattern := Day9Pattern(ParseInts(input))
		next := pattern.Next()
		total += next
	}

	return strconv.Itoa(total)
}

type Day9Pattern []int

func (p Day9Pattern) Next() int {
	nextPattern := make(Day9Pattern, 0)
	var numZeroes int

	for index := range p {
		if index+1 < len(p) {
			difference := p[index+1] - p[index]
			nextPattern = append(nextPattern, difference)
			if difference == 0 {
				numZeroes++
			}
		}
	}

	if numZeroes == len(p)-1 {
		return p[len(p)-1]
	}

	return p[len(p)-1] + nextPattern.Next()
}

func ParseInts(input string) (nums []int) {
	parts := strings.Split(input, " ")

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			nums = append(nums, num)
		}
	}

	return
}
