package twentythree_test

import (
	twentythree "advent/solutions/2023"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Day12RowTest struct {
	Row          twentythree.Day12Row
	Arrangements int
}

func Test_FindPossibilities(t *testing.T) {
	rows := []Day12RowTest{
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "???.###", StatusGroups: []int{1, 1, 3}}},
		{Arrangements: 4, Row: twentythree.Day12Row{Springs: ".??..??...?##.", StatusGroups: []int{1, 1, 3}}},
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "?#?#?#?#?#?#?#?", StatusGroups: []int{1, 3, 1, 6}}},
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "????.#...#...", StatusGroups: []int{4, 1, 1}}},
		{Arrangements: 4, Row: twentythree.Day12Row{Springs: "????.######..#####.", StatusGroups: []int{1, 6, 5}}},
		{Arrangements: 10, Row: twentythree.Day12Row{Springs: "?###????????", StatusGroups: []int{3, 2, 1}}},
	}

	for index, row := range rows {
		assert.Equal(t, row.Arrangements, row.Row.FindPossibilities(), index)
	}
}

func Test_FindRecursivePossibilities(t *testing.T) {
	rows := []Day12RowTest{
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "???.###", StatusGroups: []int{1, 1, 3}}},
		{Arrangements: 4, Row: twentythree.Day12Row{Springs: ".??..??...?##.", StatusGroups: []int{1, 1, 3}}},
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "?#?#?#?#?#?#?#?", StatusGroups: []int{1, 3, 1, 6}}},
		{Arrangements: 1, Row: twentythree.Day12Row{Springs: "????.#...#...", StatusGroups: []int{4, 1, 1}}},
		{Arrangements: 4, Row: twentythree.Day12Row{Springs: "????.######..#####.", StatusGroups: []int{1, 6, 5}}},
		{Arrangements: 10, Row: twentythree.Day12Row{Springs: "?###????????", StatusGroups: []int{3, 2, 1}}},
	}

	for index, row := range rows {
		assert.Equal(t, row.Arrangements, row.Row.FindRecursivePossibilities(), index)
	}
}
