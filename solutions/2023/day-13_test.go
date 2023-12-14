package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day13patternEx1 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
}

var day13patternEx2 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

var day13pattern1 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"######.#.#.",
	".#..#...###",
	".####....##",
	".####..###.",
	"#....##.##.",
	"......#.###",
	"#.##.#.#.##",
	".####.#..#.",
	"..##....##.",
	"##..##.....",
	"##..##.....",
	"..##....##.",
	".####.#..#.",
	"#.##.#.#.##",
	"......#.###",
	"#....##..#.",
	".####..###.",
}

var day13pattern95 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"...#....#.##.",
	"##.##..###..#",
	"##.##..###..#",
	"...#....#.##.",
	".#.##.....##.",
	".##.#.#..####",
	"..#.##....##.",
	"#..#.#####..#",
	"#.####..#....",
	"..#####..####",
	".###.###.#...",
}

var day13pattern97 twentythree.Day13Pattern = twentythree.Day13Pattern{
	".###..###..",
	"#.#.##.#.##",
	".########..",
	"#..#..#..##",
	"#.#.##.#.##",
	"..######...",
	".###..###..",
	"##..##..###",
	"####.######",
}

var day13pattern98 twentythree.Day13Pattern = twentythree.Day13Pattern{
	".###..###..",
	"#.#.##.#.##",
	".########..",
	"#..#..#..##",
	"#.#.##.#.##",
	"..######...",
	".###..###..",
	"##..##..###",
	"####.######",
}

var day13pattern99 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"##....#.##.##",
	"#.....#...#..",
	"#.....##..#..",
	"##....#.##.##",
	"#.....#......",
	"#.#.....#.###",
	".#.#######.##",
}

func Test_FindHorizontalMirrorIndex(t *testing.T) {
	assert.Equal(t, 0, day13patternEx1.FindHorizontalMirrorIndex())
	assert.Equal(t, 4, day13patternEx2.FindHorizontalMirrorIndex())

	assert.Equal(t, 0, day13pattern1.FindHorizontalMirrorIndex())
	assert.Equal(t, 2, day13pattern95.FindHorizontalMirrorIndex())
	assert.Equal(t, 0, day13pattern97.FindHorizontalMirrorIndex())
	assert.Equal(t, 0, day13pattern98.FindHorizontalMirrorIndex())
	assert.Equal(t, 0, day13pattern99.FindHorizontalMirrorIndex())
}

func Test_FindVerticalMirrorIndex(t *testing.T) {
	assert.Equal(t, 5, day13patternEx1.FindVerticalMirrorIndex())
	assert.Equal(t, 0, day13patternEx2.FindVerticalMirrorIndex())

	assert.Equal(t, 3, day13pattern1.FindVerticalMirrorIndex())
	assert.Equal(t, 0, day13pattern95.FindVerticalMirrorIndex())
	assert.Equal(t, 10, day13pattern97.FindVerticalMirrorIndex())
	assert.Equal(t, 10, day13pattern98.FindVerticalMirrorIndex())
	assert.Equal(t, 12, day13pattern99.FindVerticalMirrorIndex())
}

func Test_FindPatternValue(t *testing.T) {
	assert.Equal(t, 5, day13patternEx1.FindPatternValue())
	assert.Equal(t, 400, day13patternEx2.FindPatternValue())

	assert.Equal(t, 3, day13pattern1.FindPatternValue())
	assert.Equal(t, 200, day13pattern95.FindPatternValue())
	assert.Equal(t, 10, day13pattern97.FindPatternValue())
	assert.Equal(t, 10, day13pattern98.FindPatternValue())
	assert.Equal(t, 12, day13pattern99.FindPatternValue())
}

func Test_FindSmudgedPatternValue(t *testing.T) {
	// assert.Equal(t, 300, day13patternEx1.FindSmudgedPatternValue())
	// assert.Equal(t, 100, day13patternEx2.FindSmudgedPatternValue())

	assert.Equal(t, 200, day13pattern95.FindSmudgedPatternValue())
	// assert.Equal(t, 200, day13pattern99.FindSmudgedPatternValue())
}
