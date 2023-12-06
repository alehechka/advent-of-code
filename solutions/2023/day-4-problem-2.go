package twentythree

import "strconv"

func Day4Problem2(inputs []string) string {
	games := make([]Day4Game, 0)

	cards := make(map[int]int)

	for _, input := range inputs {
		game := Day4ParseGameRow(input)
		games = append(games, game)

		winningNumbersCount := game.WinningNumbersCount()
		for i := 1; i <= winningNumbersCount; i++ {
			cards[game.GameId+i] += cards[game.GameId] + 1
		}
	}

	var totalCards int
	for _, game := range games {
		totalCards += 1 + cards[game.GameId]
	}

	return strconv.Itoa(totalCards)
}

func (g Day4Game) WinningNumbersCount() (count int) {
	for _, num := range g.CardNumbers {
		if _, exists := g.WinningNumbers[num]; exists {
			count++
		}
	}

	return
}
