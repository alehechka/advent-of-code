package twentythree

import (
	"strconv"
	"strings"
)

func Day7Problem1(inputs []string) string {
	useJoker := false

	hands := Day7ParseHands(inputs, useJoker)

	SortHands(hands, useJoker)

	return strconv.Itoa(Day7Total(hands))
}

type Day7HandType int

const (
	Day7HighCard  Day7HandType = iota
	Day7OnePair   Day7HandType = iota
	Day7TwoPair   Day7HandType = iota
	Day7ThreeKind Day7HandType = iota
	Day7FullHouse Day7HandType = iota
	Day7FourKind  Day7HandType = iota
	Day7FiveKind  Day7HandType = iota
)

type Day7CardStrength int

const (
	Day7Joker Day7CardStrength = iota
	Day7Two   Day7CardStrength = iota
	Day7Three Day7CardStrength = iota
	Day7Four  Day7CardStrength = iota
	Day7Five  Day7CardStrength = iota
	Day7Six   Day7CardStrength = iota
	Day7Seven Day7CardStrength = iota
	Day7Eight Day7CardStrength = iota
	Day7Nine  Day7CardStrength = iota
	Day7Ten   Day7CardStrength = iota
	Day7Jack  Day7CardStrength = iota
	Day7Queen Day7CardStrength = iota
	Day7King  Day7CardStrength = iota
	Day7Ace   Day7CardStrength = iota
)

func CardStrength(ch byte, useJoker bool) Day7CardStrength {
	switch ch {
	case 'A':
		return Day7Ace
	case 'K':
		return Day7King
	case 'Q':
		return Day7Queen
	case 'J':
		if useJoker {
			return Day7Joker
		}
		return Day7Jack
	case 'T':
		return Day7Ten
	case '9':
		return Day7Nine
	case '8':
		return Day7Eight
	case '7':
		return Day7Seven
	case '6':
		return Day7Six
	case '5':
		return Day7Five
	case '4':
		return Day7Four
	case '3':
		return Day7Three
	case '2':
		return Day7Two
	default:
		return Day7Two
	}
}

func Day7ParseHands(inputs []string, useJoker bool) (hands []Day7Hand) {
	for _, input := range inputs {
		hands = append(hands, Day7ParseHand(input, useJoker))
	}
	return
}

func Day7ParseHand(input string, useJoker bool) (hand Day7Hand) {
	parts := strings.Split(input, " ")

	hand.Hand = parts[0]
	hand.Bet, _ = strconv.Atoi(parts[1])

	if useJoker {
		hand.ParseTypeWithJoker()
	} else {
		hand.ParseType()
	}

	return
}

type Day7Hand struct {
	Hand string
	Bet  int
	Type Day7HandType
}

func (h *Day7Hand) ParseType() {
	cards := make(map[rune]int)

	for _, ch := range h.Hand {
		cards[ch] = cards[ch] + 1
	}

	switch len(cards) {
	case 1: // only one card means all are the same and is a "five of a kind"
		h.Type = Day7FiveKind
	case 2: // two different cards means it can either be a "full house" or "four of a kind"
		for _, count := range cards {
			if count == 2 || count == 3 {
				h.Type = Day7FullHouse
			} else {
				h.Type = Day7FourKind
			}
			break
		}
	case 3: // three different cards means it can either be a "three of a kind" or "two pair"
		h.Type = Day7TwoPair
		for _, count := range cards {
			if count == 3 {
				h.Type = Day7ThreeKind
				break
			}
		}
	case 4: // 4 different cards means it is a "one pair"
		h.Type = Day7OnePair
	default: // default or 5 different cards means none match and is simply "high card"
		h.Type = Day7HighCard
	}
}
