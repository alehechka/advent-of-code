package twentythree

import "strconv"

func Day2Problem2(inputs []string) string {
	powers := 0
	for _, input := range inputs {
		game := ParseDay2Line(input)

		powers += game.Power()
	}

	return strconv.Itoa(powers)
}

func (g Day2Game) MinimumRequired() (min Day2Round) {
	for _, round := range g.Rounds {
		if round.Reds > min.Reds {
			min.Reds = round.Reds
		}
		if round.Greens > min.Greens {
			min.Greens = round.Greens
		}
		if round.Blues > min.Blues {
			min.Blues = round.Blues
		}
	}

	return
}

func (g Day2Game) Power() int {
	min := g.MinimumRequired()

	return min.Reds * min.Greens * min.Blues
}
