package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day14example1 = twentythree.Day14Platform{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

var day14exampleNorthLoaded1 = twentythree.Day14Platform{
	"OOOO.#.O..",
	"OO..#....#",
	"OO..O##..O",
	"O..#.OO...",
	"........#.",
	"..#....#.#",
	"..O..#.O.O",
	"..O.......",
	"#....###..",
	"#....#....",
}

var day14example1Cycle1 = twentythree.Day14Platform{
	".....#....",
	"....#...O#",
	"...OO##...",
	".OO#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#....",
	"......OOOO",
	"#...O###..",
	"#..OO#....",
}

var day14example1Cycle2 = twentythree.Day14Platform{
	".....#....",
	"....#...O#",
	".....##...",
	"..O#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#...O",
	".......OOO",
	"#..OO###..",
	"#.OOO#...O",
}

var day14example1Cycle3 = twentythree.Day14Platform{
	".....#....",
	"....#...O#",
	".....##...",
	"..O#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#...O",
	".......OOO",
	"#...O###.O",
	"#.OOO#...O",
}

func Test_CalculateNorthLoad(t *testing.T) {
	assert.Equal(t, 136, day14exampleNorthLoaded1.CalculateNorthLoad())
}

func Test_LoadNorth(t *testing.T) {
	assert.Equal(t, day14exampleNorthLoaded1, day14example1.LoadNorth())
}

func Test_SpinCycle(t *testing.T) {
	cycle1 := day14example1.SpinCycle()
	assert.Equal(t, day14example1Cycle1, cycle1)

	cycle2 := cycle1.SpinCycle()
	assert.Equal(t, day14example1Cycle2, cycle2)

	cycle3 := cycle2.SpinCycle()
	assert.Equal(t, day14example1Cycle3, cycle3)
}
