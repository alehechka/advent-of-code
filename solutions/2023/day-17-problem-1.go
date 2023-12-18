package twentythree

import (
	"advent/utils"
	"errors"
	"fmt"
	"strconv"
)

func Day17Problem1(inputs []string) string {
	blocks := Day17ParseBlocks(inputs)

	distance := blocks.Traverse(0, 0)

	return strconv.Itoa(distance)
}

type Day17Block struct {
	HeatLoss      int
	Visited       bool
	Distance      int
	Direction     rune
	X             int
	Y             int
	StraightCount int
}

// Day17Blocks represents all blocks in a grid where the first key is the 'x' coordinate, the second key is the 'y' coordinate and the final value is the heat loss at that location.
type Day17Blocks map[string]Day17Block

type Day17 struct {
	Blocks           Day17Blocks
	VerticalLength   int
	HorizontalLength int
}

func (d Day17) Traverse(x, y int) int {
	currentKey := utils.CoordKey(x, y)
	current := d.Blocks[currentKey]
	endKey := utils.CoordKey(d.HorizontalLength-1, d.VerticalLength-1)
	fmt.Printf("Traversing d[%d][%d]: %#v\n", x, y, current)

	if d.Blocks[endKey].Visited {
		return d.Blocks[endKey].Distance
	}

	left, leftErr := d.GetLeftBlock(current.Direction, x, y)
	right, rightErr := d.GetRightBlock(current.Direction, x, y)
	straight, straightErr := d.GetStraightBlock(current.Direction, x, y, current.StraightCount)

	unvisited := make([]Day17Block, 0)

	if leftErr == nil {
		if current.Distance+left.HeatLoss < left.Distance {
			left.Distance = current.Distance + left.HeatLoss
			left.Direction = Left
			left.StraightCount = 0
			d.Blocks[utils.CoordKey(left.X, left.Y)] = left
		}
		unvisited = append(unvisited, left)
	}

	if rightErr == nil {
		if current.Distance+right.HeatLoss < right.Distance {
			right.Distance = current.Distance + right.HeatLoss
			right.Direction = Right
			right.StraightCount = 0
			d.Blocks[utils.CoordKey(right.X, right.Y)] = right
		}
		unvisited = append(unvisited, right)
	}

	if straightErr == nil {
		if current.Distance+straight.HeatLoss < straight.Distance {
			straight.Distance = current.Distance + straight.HeatLoss
			straight.Direction = current.Direction
			straight.StraightCount = current.StraightCount + 1
			d.Blocks[utils.CoordKey(straight.X, straight.Y)] = straight
		}
		unvisited = append(unvisited, straight)
	}

	current.Visited = true
	d.Blocks[currentKey] = current
	if d.Blocks[endKey].Visited {
		return d.Blocks[endKey].Distance
	}
	delete(d.Blocks, currentKey)

	if nextBlock, err := Day17MinBlock(unvisited); err == nil {
		return d.Traverse(nextBlock.X, nextBlock.Y)
		// fmt.Printf("Next: %#v", nextBlock)
	}

	return 0
}

func Day17MinBlock(blocks []Day17Block) (Day17Block, error) {
	if len(blocks) == 0 {
		return Day17Block{}, errors.New("no blocks provided")
	}

	min := blocks[0]

	for i := 1; i < len(blocks); i++ {
		if blocks[i].Distance < min.Distance {
			min = blocks[i]
		}
	}

	return min, nil
}

func (d Day17) GetLeftBlock(direction rune, x, y int) (block Day17Block, err error) {
	switch direction {
	case Up:
		x--
	case Down:
		x++
	case Right:
		y--
	case Left:
		y++
	}

	block, err = d.GetBlock(x, y)
	// fmt.Printf("Left d[%d][%d]: %#v, %v\n", x, y, block, err)
	return
}

func (d Day17) GetRightBlock(direction rune, x, y int) (block Day17Block, err error) {
	switch direction {
	case Up:
		x++
	case Down:
		x--
	case Right:
		y++
	case Left:
		y--
	}

	block, err = d.GetBlock(x, y)
	// fmt.Printf("Right d[%d][%d]: %#v, %v\n", x, y, block, err)
	return
}

func (d Day17) GetStraightBlock(direction rune, x, y int, straightCount int) (block Day17Block, err error) {
	if straightCount >= 2 {
		return Day17Block{}, errors.New("straight limit reached or exceeded")
	}

	switch direction {
	case Up:
		y--
	case Down:
		y++
	case Right:
		x++
	case Left:
		x--
	}

	block, err = d.GetBlock(x, y)
	// fmt.Printf("Straight d[%d][%d]: %#v, %v\n", x, y, block, err)
	return
}

func (d Day17) GetBlock(x, y int) (Day17Block, error) {
	if x < 0 || x >= d.HorizontalLength || y < 0 || y >= d.VerticalLength {
		return Day17Block{}, errors.New("block is out of bounds")
	}

	if block, ok := d.Blocks[utils.CoordKey(x, y)]; ok && !block.Visited {
		return block, nil
	} else {
		return Day17Block{}, errors.New("left block has been visited and removed")
	}
}

func Day17ParseBlocks(inputs []string) Day17 {
	blocks := make(Day17Blocks)

	for y, row := range inputs {
		for x, ch := range row {
			num, _ := strconv.Atoi(string(ch))
			distance := utils.MaxInt
			var direction rune
			if x == 0 && y == 0 {
				distance = 0
				direction = Right
			}

			blocks[utils.CoordKey(x, y)] = Day17Block{
				HeatLoss:  num,
				Visited:   false,
				Distance:  distance,
				Direction: direction,
				X:         x,
				Y:         y,
			}
		}
	}

	return Day17{
		Blocks:           blocks,
		VerticalLength:   len(inputs),
		HorizontalLength: len(inputs[0]),
	}
}
