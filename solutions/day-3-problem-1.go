package solutions

import (
	"regexp"
	"strconv"
)

func Day3Problem1(inputs []string) string {
	var total int

	for index, input := range inputs {
		above := ""
		if index > 0 {
			above = inputs[index-1]
		}
		below := ""
		if len(inputs) > index+1 {
			below = inputs[index+1]
		}
		_, rowTotal := Day3Row(input).ParseNumbersWithAdjacent(index, above, below)
		total += rowTotal
	}

	return strconv.Itoa(total)
}

type Day3Row string

func (r Day3Row) ParseNumbersWithAdjacent(rowIndex int, rowAbove string, rowBelow string) (nums []Day3Number, total int) {
	if r == "" {
		return
	}

	numbers := r.ParseNumbersFromRow(rowIndex)
	for _, number := range numbers {
		number.SetAdjacent(string(r), rowAbove, rowBelow)
		if number.HasAdjacentSpecial() {
			nums = append(nums, number)
			total += number.Number()
		}
	}

	return
}

type Day3Number struct {
	Raw        string
	Index      int
	Row        int
	CharsAbove string
	CharLeft   string
	CharRight  string
	CharsBelow string
}

func (n Day3Number) Number() int {
	i, _ := strconv.Atoi(n.Raw)
	return i
}

func (n Day3Number) Len() int {
	return len(n.Raw)
}

func (n Day3Number) HasAdjacentSpecial() bool {
	adjacent := n.CharsAbove + n.CharLeft + n.CharRight + n.CharsBelow

	reg, _ := regexp.Compile(`[^0-9\.]`)

	return reg.Match([]byte(adjacent))
}

func (n *Day3Number) SetAdjacent(current string, above string, below string) {
	adjacentLength := n.Len()
	adjacentStartIndex := n.Index
	if n.Index > 0 {
		adjacentLength++
		adjacentStartIndex--
		n.CharLeft = string(current[n.Index-1])
	}

	if n.Index+n.Len() < len(current) {
		adjacentLength++
		n.CharRight = string(current[n.Index+n.Len()])
	}

	adjacentEndIndex := adjacentStartIndex + adjacentLength
	if len(current) == len(above) {
		n.CharsAbove = above[adjacentStartIndex:adjacentEndIndex]
	}

	if len(current) == len(below) {
		n.CharsBelow = below[adjacentStartIndex:adjacentEndIndex]
	}
}

func (r Day3Row) ParseNumbersFromRow(rowIndex int) []Day3Number {
	nums := make([]Day3Number, 0)

	current := Day3Number{Row: rowIndex}
	for index := range r {
		ch := r[index]
		if isInt(ch) {
			if current.Raw == "" {
				current.Index = index
			}
			current.Raw += string(ch)
		} else if current.Raw != "" {
			nums = append(nums, current)
			current = Day3Number{Row: rowIndex}
		}
		if index == len(r)-1 && current.Raw != "" {
			nums = append(nums, current)
		}
	}

	return nums
}
