package solutions

import (
	"strconv"
	"strings"
)

func Day6Problem1(inputs []string) string {
	races := Day6ParseInputRows(inputs[0], inputs[1])

	product := 1
	for _, race := range races {
		product *= race.NumValidHoldTimes()
	}

	return strconv.Itoa(product)
}

type Day6Race struct {
	RaceTime       int
	RecordDistance int
}

func (r Day6Race) ValidHoldTimes() (holdTimes []int) {
	for i := 1; i < r.RaceTime; i++ {
		distanceTraveled := i * (r.RaceTime - i)
		if distanceTraveled > r.RecordDistance {
			holdTimes = append(holdTimes, i)
		}
	}
	return
}

func (r Day6Race) NumValidHoldTimes() int {
	for i := 1; i < r.RaceTime; i++ {
		distanceTraveled := i * (r.RaceTime - i)
		if distanceTraveled > r.RecordDistance {
			return r.RaceTime - i - i + 1
		}
	}
	return 0
}

func Day6ParseInputRows(timesRow string, distancesRow string) (races []Day6Race) {
	times := GetNumbersFromSpaceDelimited(timesRow)
	distances := GetNumbersFromSpaceDelimited(distancesRow)

	for index := range times {
		races = append(races, Day6Race{RaceTime: times[index], RecordDistance: distances[index]})
	}

	return
}

func GetNumbersFromSpaceDelimited(input string) (nums []int) {
	parts := strings.Split(input, " ")

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			nums = append(nums, num)
		}
	}

	return
}
