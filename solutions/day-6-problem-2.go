package solutions

import (
	"strconv"
	"strings"
)

func Day6Problem2(inputs []string) string {
	race := Day6Race{RaceTime: GetNumberFromSpaceDelimited(inputs[0]), RecordDistance: GetNumberFromSpaceDelimited(inputs[1])}

	validHoldTimes := race.ValidHoldTimes()

	return strconv.Itoa(len(validHoldTimes))
}

func GetNumberFromSpaceDelimited(input string) int {
	parts := strings.Split(input, " ")

	var rawNum string
	for _, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			rawNum += part
		}
	}

	num, _ := strconv.Atoi(rawNum)
	return num
}
