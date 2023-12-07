package day07

import (
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	t     Type
	cards []Card
	bid   int
}

type Card int

const (
	C2 Card = iota
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	T
	J
	Q
	K
	A
)

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func Part1(input string) string {
	hands := parse(input)
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].t == hands[j].t {
			// sort by highest card
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
			panic("cards all equal")
		} else {
			return hands[i].t < hands[j].t
		}
	})
	totalWin := 0
	for i, hand := range hands {
		rank := i + 1
		totalWin += rank * hand.bid
	}
	return strconv.Itoa(totalWin)
}

func parse(s string) (result []Hand) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		c := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		cards := parseCards(c)
		t := getType(cards)
		hand := Hand{
			t:     t,
			cards: cards,
			bid:   bid,
		}
		result = append(result, hand)
	}
	return
}

func getType(cards []Card) Type {
	var counts map[Card]int = make(map[Card]int)
	for _, card := range cards {
		current := counts[card]
		counts[card] = current + 1
	}
	if hasExactCount(counts, 5) {
		return FiveOfAKind
	} else if hasExactCount(counts, 4) {
		return FourOfAKind
	} else if hasExactCount(counts, 3) && hasExactCount(counts, 2) {
		return FullHouse
	} else if hasExactCount(counts, 3) {
		return ThreeOfAKind
	} else if hasExactCount(counts, 2) {
		pairs := 0
		for _, count := range counts {
			if count == 2 {
				pairs++
			}
		}
		if pairs == 2 {
			return TwoPairs
		} else {
			return OnePair
		}
	} else {
		return HighCard
	}
}

func hasExactCount(counts map[Card]int, n int) bool {
	for _, count := range counts {
		if count == n {
			return true
		}
	}
	return false
}

func parseCards(c string) []Card {
	var cards []Card
	for _, card := range c {
		var ct Card
		switch card {
		case '2':
			ct = C2
		case '3':
			ct = C3
		case '4':
			ct = C4
		case '5':
			ct = C5
		case '6':
			ct = C6
		case '7':
			ct = C7
		case '8':
			ct = C8
		case '9':
			ct = C9
		case 'T':
			ct = T
		case 'J':
			ct = J
		case 'Q':
			ct = Q
		case 'K':
			ct = K
		case 'A':
			ct = A
		default:
			panic("bad card")
		}
		cards = append(cards, ct)
	}
	return cards
}

func Part2(input string) string {
	return ""
}
