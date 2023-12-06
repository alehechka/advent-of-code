package fifteen

import "strconv"

func Day2Problem2(inputs []string) string {
	dimensions := Day2ParseDimensions(inputs)

	var total int
	for _, d := range dimensions {
		total += d.ShortestPerimeter() + d.Volume()
	}

	return strconv.Itoa(total)
}

func (d Day2Dimensions) Perimeter() int {
	return d.Length*2 + d.Width*2
}

func (d Day2Dimensions) ShortestPerimeter() int {
	a, b := d.ShortestSides()

	return a*2 + b*2
}

func (d Day2Dimensions) ShortestSides() (int, int) {
	if d.Height >= d.Length && d.Height >= d.Width {
		return d.Length, d.Width
	}

	if d.Length >= d.Height && d.Length >= d.Width {
		return d.Height, d.Width
	}

	return d.Height, d.Length
}

func (d Day2Dimensions) Volume() int {
	return d.Length * d.Width * d.Height
}
