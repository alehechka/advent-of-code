package fifteen

import (
	"strconv"
	"strings"
)

func Day2Problem1(inputs []string) string {
	dimensions := Day2ParseDimensions(inputs)

	var total int
	for _, d := range dimensions {
		total += d.SurfaceArea() + d.SmallestArea()
	}

	return strconv.Itoa(total)
}

type Day2Dimensions struct {
	Width  int
	Height int
	Length int
}

func (d Day2Dimensions) SurfaceArea() int {
	return 2*d.Length*d.Width + 2*d.Width*d.Height + 2*d.Height*d.Length
}

func (d Day2Dimensions) SmallestArea() int {
	smallest := d.Width * d.Height

	if area := d.Height * d.Length; area < smallest {
		smallest = area
	}

	if area := d.Width * d.Length; area < smallest {
		smallest = area
	}

	return smallest
}

func Day2ParseDimensions(inputs []string) (dimensions []Day2Dimensions) {
	for _, input := range inputs {
		parts := strings.Split(input, "x")
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])

		dimensions = append(dimensions, Day2Dimensions{Length: l, Width: w, Height: h})
	}

	return
}
