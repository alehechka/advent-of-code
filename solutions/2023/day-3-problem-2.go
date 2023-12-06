package twentythree

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day3Problem2(inputs []string) string {
	nums := make([]Day3Number, 0)
	for index, input := range inputs {
		above := ""
		if index > 0 {
			above = inputs[index-1]
		}
		below := ""
		if len(inputs) > index+1 {
			below = inputs[index+1]
		}
		currentNums, _ := Day3Row(input).ParseNumbersWithAdjacent(index, above, below)
		nums = append(nums, currentNums...)
	}

	gearMap := Day3NumbersToGearIndexMap(nums)

	return strconv.Itoa(GearIndexMapRatioTotal(gearMap))
}

func GearIndexMapRatioTotal(gearMap map[string][]int) (total int) {
	for _, nums := range gearMap {
		if len(nums) == 2 {
			total += nums[0] * nums[1]
		}
	}

	return
}

func Day3NumbersToGearIndexMap(nums []Day3Number) map[string][]int {
	gearMap := make(map[string][]int)

	for _, num := range nums {
		for _, gearIndex := range num.GearIndices() {
			gearKey := gearIndex.Key()
			if gearMap[gearKey] == nil {
				gearMap[gearKey] = make([]int, 0)
			}
			gearMap[gearKey] = append(gearMap[gearKey], gearIndex.Number)
		}
	}

	return gearMap
}

type Day3Index struct {
	Row    int
	Index  int
	Number int
}

func (i Day3Index) Key() string {
	return fmt.Sprintf("%d-%d", i.Row, i.Index)
}

func (n Day3Number) GearIndices() (indices []Day3Index) {
	num := n.Number()

	for _, gearIndex := range utils.Indices(n.CharsAbove, "*") {
		indices = append(indices, Day3Index{Row: n.Row - 1, Index: n.Index + gearIndex - 1, Number: num})
	}

	if strings.Contains(n.CharLeft, "*") {
		indices = append(indices, Day3Index{Row: n.Row, Index: n.Index - 1, Number: num})
	}

	if strings.Contains(n.CharRight, "*") {
		indices = append(indices, Day3Index{Row: n.Row, Index: n.Index + n.Len(), Number: num})
	}

	for _, gearIndex := range utils.Indices(n.CharsBelow, "*") {
		indices = append(indices, Day3Index{Row: n.Row + 1, Index: n.Index + gearIndex - 1, Number: num})
	}

	return
}
