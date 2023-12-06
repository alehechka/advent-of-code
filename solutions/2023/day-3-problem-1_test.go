package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type day3Datum struct {
	line    string
	numbers []twentythree.Day3Number
}

var day3Data = []day3Datum{
	{
		line: "............................................411.....................363..134.........463.775..........................506...................",
		numbers: []twentythree.Day3Number{
			{Raw: "411", Index: 44, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "...*."},
			{Raw: "363", Index: 68, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "....*"},
			{Raw: "134", Index: 73, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "...=."},
			{Raw: "463", Index: 85, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "....*"},
			{Raw: "775", Index: 89, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "*...."},
			{Raw: "506", Index: 118, CharsAbove: "", CharLeft: ".", CharRight: ".", CharsBelow: "./..."},
		},
	},
	{
		line: "......429...836..$............../..960........*.............+..........*...=....381.....*........67......426.....=..../...304...............",
		numbers: []twentythree.Day3Number{
			{Raw: "429", Index: 6, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "...72"},
			{Raw: "836", Index: 12, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "960", Index: 35, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "381", Index: 80, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "3&../"},
			{Raw: "67", Index: 97, CharsAbove: "....", CharLeft: ".", CharRight: ".", CharsBelow: "696."},
			{Raw: "426", Index: 105, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "..131"},
			{Raw: "304", Index: 122, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "....="},
		},
	},
	{
		line: "........721........*...17................790...670.............$..........893.93&../...129$.651.696.......131*99.............=......446*781.",
		numbers: []twentythree.Day3Number{
			{Raw: "721", Index: 8, CharsAbove: "29...", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "17", Index: 23, CharsAbove: "....", CharLeft: ".", CharRight: ".", CharsBelow: "...."},
			{Raw: "790", Index: 41, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "670", Index: 47, CharsAbove: "*....", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "893", Index: 74, CharsAbove: "..=..", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "93", Index: 78, CharsAbove: "...3", CharLeft: ".", CharRight: "&", CharsBelow: "...."},
			{Raw: "129", Index: 87, CharsAbove: "..*..", CharLeft: ".", CharRight: "$", CharsBelow: "....."},
			{Raw: "651", Index: 92, CharsAbove: ".....", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "696", Index: 96, CharsAbove: "..67.", CharLeft: ".", CharRight: ".", CharsBelow: "....."},
			{Raw: "131", Index: 106, CharsAbove: "426..", CharLeft: ".", CharRight: "*", CharsBelow: "....."},
			{Raw: "99", Index: 110, CharsAbove: "....", CharLeft: "*", CharRight: ".", CharsBelow: "...."},
			{Raw: "446", Index: 132, CharsAbove: ".....", CharLeft: ".", CharRight: "*", CharsBelow: "....."},
			{Raw: "781", Index: 136, CharsAbove: ".....", CharLeft: "*", CharRight: ".", CharsBelow: "...*9"},
		},
	},
	{
		line: "1$........................................................................................................................................*9",
		numbers: []twentythree.Day3Number{
			{Raw: "1", Index: 0, CharsAbove: "..", CharLeft: "", CharRight: "$", CharsBelow: ""},
			{Raw: "9", Index: 139, CharsAbove: "1.", CharLeft: "*", CharRight: "", CharsBelow: ""},
		},
	},
}

func Test_Day3Number_ParseNumbersFromRow(t *testing.T) {
	for index, data := range day3Data {
		test_Day3Number_ParseNumbersFromRow(t, index, data)
	}
}

func test_Day3Number_ParseNumbersFromRow(t *testing.T, rowIndex int, data day3Datum) {
	nums := twentythree.Day3Row(data.line).ParseNumbersFromRow(rowIndex)

	assert.Equal(t, len(data.numbers), len(nums))
	fmt.Println(nums)

	for index, num := range data.numbers {
		assert.Equal(t, num.Index, nums[index].Index, "Index not equal")
		assert.Equal(t, num.Raw, nums[index].Raw, "Raw string not equal")
		assert.Equal(t, num.Number(), nums[index].Number(), "Parsed int not equal")
		assert.Equal(t, num.Len(), nums[index].Len(), "Raw length not equal")
	}
}

func Test_Day3Number_SetAdjacent(t *testing.T) {
	for index, data := range day3Data {
		above := ""
		if index > 0 {
			above = day3Data[index-1].line
		}
		below := ""
		if len(day3Data) > index+1 {
			below = day3Data[index+1].line
		}
		test_Day3Number_SetAdjacent(t, data, above, below)
	}
}

func test_Day3Number_SetAdjacent(t *testing.T, data day3Datum, above string, below string) {
	for _, expected := range data.numbers {
		actual := twentythree.Day3Number{Raw: expected.Raw, Index: expected.Index}
		actual.SetAdjacent(data.line, above, below)

		assert.Equal(t, expected.CharsAbove, actual.CharsAbove, fmt.Sprintf("CharsAbove not equal: [%d]%s", expected.Index, expected.Raw))
		assert.Equal(t, expected.CharLeft, actual.CharLeft, fmt.Sprintf("CharLeft not equal: [%d]%s", expected.Index, expected.Raw))
		assert.Equal(t, expected.CharRight, actual.CharRight, fmt.Sprintf("CharRight not equal: [%d]%s", expected.Index, expected.Raw))
		assert.Equal(t, expected.CharsBelow, actual.CharsBelow, fmt.Sprintf("CharsBelow not equal: [%d]%s", expected.Index, expected.Raw))
	}
}

func Test_Day3Number_HasAdjacentSpecial(t *testing.T) {
	assert.True(t, twentythree.Day3Number{CharsAbove: "..*."}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharLeft: "*"}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharRight: "*"}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharsBelow: "..*.."}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharsBelow: "..$.."}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharsBelow: "../.."}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharsBelow: "..=.."}.HasAdjacentSpecial())
	assert.True(t, twentythree.Day3Number{CharsBelow: "..+.."}.HasAdjacentSpecial())

	assert.False(t, twentythree.Day3Number{}.HasAdjacentSpecial())
	assert.False(t, twentythree.Day3Number{CharsAbove: "..34.."}.HasAdjacentSpecial())
	assert.False(t, twentythree.Day3Number{CharLeft: "."}.HasAdjacentSpecial())
	assert.False(t, twentythree.Day3Number{CharRight: "."}.HasAdjacentSpecial())
	assert.False(t, twentythree.Day3Number{CharsBelow: "5...6"}.HasAdjacentSpecial())
	assert.False(t, twentythree.Day3Number{CharsAbove: "..34..", CharLeft: ".", CharRight: ".", CharsBelow: "5...6"}.HasAdjacentSpecial())
}
