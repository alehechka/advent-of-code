package solutions_test

import (
	"advent/solutions"
	"testing"

	"github.com/stretchr/testify/assert"
)

type day2Datum struct {
	line       string
	game       solutions.Day2Game
	isPossible bool
}

var data = []day2Datum{
	{
		line: "Game 1: 4 green, 2 blue; 1 red, 1 blue, 4 green; 3 green, 4 blue, 1 red; 7 green, 2 blue, 4 red; 3 red, 7 green; 3 red, 3 green",
		game: solutions.Day2Game{ID: 1, Rounds: []solutions.Day2Round{
			{Greens: 4, Blues: 2},
			{Reds: 1, Blues: 1, Greens: 4},
			{Greens: 3, Blues: 4, Reds: 1},
			{Greens: 7, Blues: 2, Reds: 4},
			{Reds: 3, Greens: 7},
			{Reds: 3, Greens: 3},
		}},
		isPossible: true,
	},
	{
		line: "Game 18: 6 red, 4 green, 7 blue; 2 red, 3 green, 12 blue; 3 red, 6 blue, 6 green; 9 red, 10 blue; 6 green, 4 blue, 2 red; 12 red, 12 blue, 9 green",
		game: solutions.Day2Game{ID: 18, Rounds: []solutions.Day2Round{
			{Reds: 6, Greens: 4, Blues: 7},
			{Reds: 2, Greens: 3, Blues: 12},
			{Reds: 3, Blues: 6, Greens: 6},
			{Reds: 9, Blues: 10},
			{Greens: 6, Blues: 4, Reds: 2},
			{Reds: 12, Blues: 12, Greens: 9},
		}},
		isPossible: true,
	},
	{
		line: "Game 41: 5 red, 14 blue, 3 green; 3 red, 3 blue, 7 green; 19 blue, 15 green, 6 red",
		game: solutions.Day2Game{ID: 41, Rounds: []solutions.Day2Round{
			{Reds: 5, Blues: 14, Greens: 3},
			{Reds: 3, Blues: 3, Greens: 7},
			{Blues: 19, Greens: 15, Reds: 6},
		}},
		isPossible: false,
	},
}

func Test_ParseDay2Line(t *testing.T) {

	for _, d := range data {
		test_ParseDay2Line(t, d)
	}
}

func test_ParseDay2Line(t *testing.T, expected day2Datum) {
	actual := solutions.ParseDay2Line(expected.line)

	assert.Equal(t, expected.game.ID, actual.ID)
	assert.Equal(t, len(expected.game.Rounds), len(actual.Rounds))
	assert.Equal(t, expected.game.Rounds, actual.Rounds)
	assert.Equal(t, expected.isPossible, actual.IsPossible())
}
