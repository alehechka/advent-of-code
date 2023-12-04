package solutions

import (
	"strconv"
	"strings"
)

func Day4Problem1(inputs []string) string {
	games := make([]Day4Game, 0)

	var total int
	for _, input := range inputs {
		game := Day4ParseGameRow(input)
		games = append(games, game)
		total += game.Score()
	}

	return strconv.Itoa(total)
}

type Day4Game struct {
	GameId         int
	WinningNumbers map[int]bool
	CardNumbers    []int
}

func (g Day4Game) Score() (score int) {
	for _, num := range g.CardNumbers {
		if _, exists := g.WinningNumbers[num]; exists {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return
}

func Day4ParseGameRow(str string) (game Day4Game) {
	game.WinningNumbers = make(map[int]bool)
	game.CardNumbers = make([]int, 0)

	gameParts := strings.Split(str, ":")

	gameIDParts := strings.Split(gameParts[0], " ")
	game.GameId, _ = strconv.Atoi(gameIDParts[len(gameIDParts)-1])

	numberParts := strings.Split(gameParts[1], "|")

	winningNumbersParts := strings.Split(numberParts[0], " ")
	for _, winningNumber := range winningNumbersParts {
		if num, err := strconv.Atoi(winningNumber); err == nil {
			game.WinningNumbers[num] = true
		}
	}

	cardNumbersParts := strings.Split(numberParts[1], " ")
	for _, cardNumber := range cardNumbersParts {
		if num, err := strconv.Atoi(cardNumber); err == nil {
			game.CardNumbers = append(game.CardNumbers, num)
		}
	}

	return
}
