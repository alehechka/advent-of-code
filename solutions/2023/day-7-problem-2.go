package twentythree

import (
	"sort"
	"strconv"
)

func Day7Problem2(inputs []string) string {
	useJoker := true

	hands := Day7ParseHands(inputs, useJoker)

	SortHands(hands, useJoker)

	return strconv.Itoa(Day7Total(hands))
}

func SortHands(hands []Day7Hand, useJoker bool) {
	sort.Slice(hands, func(i, j int) bool {
		left := hands[i]
		right := hands[j]

		if left.Type == right.Type {
			for index := range left.Hand {
				leftStr := CardStrength(left.Hand[index], useJoker)
				rightStr := CardStrength(right.Hand[index], useJoker)

				if leftStr != rightStr {
					return leftStr < rightStr
				}
			}
		}

		return left.Type < right.Type
	})
}

func Day7Total(hands []Day7Hand) (total int) {
	for index, hand := range hands {
		total += (index + 1) * hand.Bet
	}
	return
}

func (h *Day7Hand) ParseTypeWithJoker() {
	cards := make(map[rune]int)

	for _, ch := range h.Hand {
		cards[ch] = cards[ch] + 1
	}

	jokerCount := cards['J']

	switch len(cards) {
	case 1: // only one card means all are the same and is a "five of a kind"
		h.Type = Day7FiveKind
	case 2: // two different cards means it can either be a "full house" or "four of a kind"
		for _, count := range cards {
			// full house
			if count == 2 || count == 3 {
				if jokerCount == 2 || jokerCount == 3 {
					h.Type = Day7FiveKind
					return
				}
				h.Type = Day7FullHouse
				return
			}
			// four of a kind
			if jokerCount == 4 || jokerCount == 1 {
				h.Type = Day7FiveKind
				return
			}
			h.Type = Day7FourKind
			return
		}
	case 3: // three different cards means it can either be a "three of a kind" or "two pair"
		h.Type = Day7TwoPair
		for _, count := range cards {
			if count == 3 {
				if jokerCount == 3 || jokerCount == 1 {
					h.Type = Day7FourKind
					return
				}
				h.Type = Day7ThreeKind
				return
			}
		}
		if jokerCount == 1 {
			h.Type = Day7FullHouse
			return
		}
		if jokerCount == 2 {
			h.Type = Day7FourKind
			return
		}
	case 4: // 4 different cards means it is a "one pair"
		if jokerCount == 1 || jokerCount == 2 {
			h.Type = Day7ThreeKind
			return
		}
		h.Type = Day7OnePair
	default: // default or 5 different cards means none match and is simply "high card"
		if jokerCount == 1 {
			h.Type = Day7OnePair
			return
		}
		h.Type = Day7HighCard
	}
}
