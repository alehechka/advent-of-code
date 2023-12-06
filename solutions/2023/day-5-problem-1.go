package twentythree

import (
	"strconv"
	"strings"
)

func Day5Problem1(inputs []string) string {
	seeds := Day5ParseSeeds(inputs[0])

	categoryMap := Day5ParseCategoryMap(inputs[1:])

	var lowest int
	for _, seed := range seeds {
		dest := categoryMap.Traverse("seed", "location", seed)
		if lowest == 0 || dest < lowest {
			lowest = dest
		}
	}

	return strconv.Itoa(lowest)
}

func Day5ParseSeeds(input string) (seeds []int) {
	parts := strings.Split(input, " ")

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			seeds = append(seeds, num)
		}
	}
	return
}

func Day5ParseCategoryMap(inputs []string) Day5CategoryMap {
	categoryMap := make(Day5CategoryMap)

	var categoryStart, categoryEnd string
	for _, input := range inputs {
		parts := strings.Split(input, " ")
		if len(parts) <= 1 {
			continue
		}
		if strings.Contains(input, "map") {
			categoryParts := strings.Split(parts[0], "-")
			categoryStart = categoryParts[0]
			categoryEnd = categoryParts[2]
			if _, ok := categoryMap[categoryStart]; !ok {
				categoryMap[categoryStart] = Day5Category{
					Current:     categoryStart,
					Destination: categoryEnd,
					RangeMaps:   make(Day5RangeMaps, 0),
				}
			}
		} else {
			source, _ := strconv.Atoi(parts[1])
			dest, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[2])
			category := categoryMap[categoryStart]
			category.RangeMaps = append(categoryMap[categoryStart].RangeMaps, Day5RangeMap{
				Source:      source,
				Destination: dest,
				Range:       r,
			})
			categoryMap[categoryStart] = category
		}
	}

	return categoryMap
}

type Day5CategoryMap map[string]Day5Category

func (c Day5CategoryMap) Traverse(startCategory string, endCategory string, startIndex int) int {
	category, ok := c[startCategory]
	if !ok {
		return 0
	}

	nextCategory, dest := category.GetDestination(startIndex)
	if nextCategory == endCategory {
		return dest
	}

	return c.Traverse(nextCategory, endCategory, dest)
}

type Day5Category struct {
	Current     string
	Destination string
	RangeMaps   Day5RangeMaps
}

func (c Day5Category) GetDestination(start int) (category string, dest int) {
	return c.Destination, c.RangeMaps.GetDestination(start)
}

type Day5RangeMap struct {
	Source      int
	Destination int
	Range       int
}

func (r Day5RangeMap) GetDestination(start int) (dest int, ok bool) {
	if start >= r.Source && start <= r.Source+r.Range {
		return r.Destination + start - r.Source, true
	}
	return
}

type Day5RangeMaps []Day5RangeMap

func (r Day5RangeMaps) GetDestination(start int) (dest int) {
	for _, m := range r {
		if d, ok := m.GetDestination(start); ok {
			return d
		}
	}
	return start
}
