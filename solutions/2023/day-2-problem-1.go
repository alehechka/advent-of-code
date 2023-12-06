package twentythree

import (
	"strconv"
	"strings"
)

func Day2Problem1(inputs []string) string {

	games := make([]Day2Game, 0)
	possibleGames := make([]int, 0)
	total := 0
	for _, input := range inputs {
		game := ParseDay2Line(input)
		games = append(games, game)

		if game.IsPossible() {
			possibleGames = append(possibleGames, game.ID)
			total += game.ID
		}
	}

	return strconv.Itoa(total)
}

type Day2Round struct {
	Reds   int
	Blues  int
	Greens int
}

func (r Day2Round) IsPossible() bool {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	if r.Reds > maxRed || r.Greens > maxGreen || r.Blues > maxBlue {
		return false
	}
	return true
}

type Day2Game struct {
	ID     int
	Rounds []Day2Round
}

func (g Day2Game) IsPossible() bool {
	for _, round := range g.Rounds {
		if !round.IsPossible() {
			return false
		}
	}

	return true
}

// ParseDay2Line parses the following line into a `Day2Game`
//
//	Game 1: 4 green, 2 blue; 1 red, 1 blue, 4 green; 3 green, 4 blue, 1 red; 7 green, 2 blue, 4 red; 3 red, 7 green; 3 red, 3 green
func ParseDay2Line(str string) (game Day2Game) {
	gameParts := strings.Split(str, ":")

	game.ID, _ = strconv.Atoi(strings.Split(gameParts[0], " ")[1])

	rounds := strings.Split(gameParts[1], ";")
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		var r Day2Round
		for _, color := range colors {
			parts := strings.Split(strings.TrimSpace(color), " ")
			num, _ := strconv.Atoi(parts[0])
			color := parts[1]
			switch color {
			case "red":
				r.Reds += num
			case "green":
				r.Greens += num
			case "blue":
				r.Blues += num
			}
		}
		game.Rounds = append(game.Rounds, r)
	}

	return
}
