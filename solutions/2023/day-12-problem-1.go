package twentythree

import (
	"strconv"
	"strings"
)

func Day12Problem1(inputs []string) string {
	rows := Day12ParseInput(inputs)

	var sum int
	for _, row := range rows {
		sum += row.FindPossibilities()
	}

	return strconv.Itoa(sum)
}

type Day12Row struct {
	Springs      string
	StatusGroups []int
}

func (r Day12Row) FindPossibilities() (count int) {
	filled := Day12FillUnknownsRecursive(r.Springs)

	for _, fill := range filled {
		if r.IsFilledValid(fill) {
			// fmt.Printf("%s => %s, %v\n", r.Springs, fill, r.StatusGroups)
			count++
		}
	}

	return
}

func (r Day12Row) IsFilledValid(str string) bool {
	counts := make([]int, 0)

	var count int
	for index, ch := range str {
		if ch == '#' {
			count++
			if index == len(str)-1 {
				counts = append(counts, count)
			}
		} else if count == 0 {
			continue
		} else {
			counts = append(counts, count)
			count = 0
		}
	}

	// fmt.Printf("%s => %s | %v == %v\n", r.Springs, str, counts, r.StatusGroups)

	if len(counts) != len(r.StatusGroups) {
		return false
	}

	for index := range counts {
		if counts[index] != r.StatusGroups[index] {
			return false
		}
	}

	return true
}

func Day12FillUnknownsRecursive(str string) (filled []string) {
	var next []string
	if len(str) > 1 {
		next = Day12FillUnknownsRecursive(str[1:])
	} else if len(str) == 0 {
		return
	} else {
		next = append(next, "")
	}

	for _, n := range next {
		if str[0] == '?' {
			filled = append(filled, "."+n, "#"+n)
		} else {
			filled = append(filled, string(str[0])+n)
		}
	}

	return
}

func Day12ParseInput(inputs []string) (rows []Day12Row) {
	for _, input := range inputs {
		rows = append(rows, Day12ParseRow(input))
	}
	return
}

func Day12ParseRow(input string) (row Day12Row) {
	parts := strings.Split(input, " ")

	row.Springs = parts[0]

	nums := strings.Split(parts[1], ",")
	for _, num := range nums {
		if x, err := strconv.Atoi(num); err == nil {
			row.StatusGroups = append(row.StatusGroups, x)
		}
	}

	return
}
