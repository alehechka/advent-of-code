package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day13pattern1 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
}

var day13pattern2 twentythree.Day13Pattern = twentythree.Day13Pattern{
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

var day13pattern3 twentythree.Day13Pattern = twentythree.Day13Pattern{
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

func Test_FindHorizontalMirrorIndex(t *testing.T) {
	assert.Equal(t, 0, day13pattern1.FindHorizontalMirrorIndex())
	assert.Equal(t, 4, day13pattern2.FindHorizontalMirrorIndex())
	assert.Equal(t, 0, day13pattern3.FindHorizontalMirrorIndex())
}

func Test_FindVerticalMirrorIndex(t *testing.T) {
	assert.Equal(t, 5, day13pattern1.FindVerticalMirrorIndex())
	assert.Equal(t, 0, day13pattern2.FindVerticalMirrorIndex())
	assert.Equal(t, 3, day13pattern3.FindVerticalMirrorIndex())
}
