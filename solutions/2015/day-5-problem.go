package fifteen

import (
	"strconv"
)

func Day5Problem1(inputs []string) string {

	var count int
	for _, input := range inputs {
		if Day5Problem1IsStringNice(input) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Day5Problem2(inputs []string) string {
	var count int
	for _, input := range inputs {
		if Day5Problem2IsStringNice(input) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Day5Problem1IsStringNice(str string) bool {
	var vowels int
	var hasDouble bool
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}

		if i > 0 {
			switch str[i-1 : i+1] {
			case "ab", "cd", "pq", "xy":
				return false
			}

			if !hasDouble && str[i-1] == str[i] {
				hasDouble = true
			}
		}
	}

	return vowels >= 3 && hasDouble
}

func Day5Problem2IsStringNice(str string) bool {
	pairs := make(map[string]map[int]bool)

	var hasRepeat bool
	for i := 1; i < len(str); i++ {
		if !hasRepeat && i > 1 && str[i-2] == str[i] {
			hasRepeat = true
		}

		key := string(str[i-1]) + string(str[i])

		indices := pairs[key]
		if indices == nil {
			indices = make(map[int]bool)
		}
		indices[i] = true
		pairs[key] = indices
	}

	if !hasRepeat {
		return false
	}

	for _, indices := range pairs {
		if len(indices) < 2 {
			continue
		}
		for index := range indices {
			_, hasPrev := indices[index-1]
			_, hasNext := indices[index+1]
			if !hasPrev && !hasNext {
				return true
			}
		}
	}

	return false
}
