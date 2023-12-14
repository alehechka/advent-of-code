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

var day14exampleLoaded1 = twentythree.Day14Platform{
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

func Test_CalculateNorthLoad(t *testing.T) {
	assert.Equal(t, 136, day14exampleLoaded1.CalculateNorthLoad())
}

func Test_LoadNorth(t *testing.T) {
	assert.Equal(t, day14exampleLoaded1, day14example1.LoadNorth())
}
