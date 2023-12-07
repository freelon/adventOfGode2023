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
	var counts = make(map[Card]int)
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
	hands := parseForBest(input)
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

func parseForBest(s string) (result []Hand2) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		c := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		cards := parseCards2(c)
		t := getTypeBest(cards)
		hand := Hand2{
			t:     t,
			cards: cards,
			bid:   bid,
		}
		result = append(result, hand)
	}
	return
}

func getTypeBest(cards []Card2) Type {
	var realCounts = make(map[Card2]int)
	for _, card := range cards {
		current := realCounts[card]
		realCounts[card] = current + 1
	}
	if realCounts[NJ] == 0 {
		return type2OfCount(realCounts)
	}

	jokers := realCounts[NJ]
	realCounts[NJ] = 0
	var fakeHands []Type
	for _, cardToBoost := range []Card2{NC2, NC3, NC4, NC5, NC6, NC7, NC8, NC9, NT, NQ, NK, NA} {
		fakeCounts := realCounts
		fakeCounts[cardToBoost] += jokers
		fakeHands = append(fakeHands, type2OfCount(fakeCounts))
		fakeCounts[cardToBoost] -= jokers
	}
	sort.Slice(fakeHands, func(i, j int) bool {
		return fakeHands[i] < fakeHands[j]
	})
	return fakeHands[len(fakeHands)-1]
}

func type2OfCount(counts map[Card2]int) Type {
	if hasExactCount2(counts, 5) {
		return FiveOfAKind
	} else if hasExactCount2(counts, 4) {
		return FourOfAKind
	} else if hasExactCount2(counts, 3) && hasExactCount2(counts, 2) {
		return FullHouse
	} else if hasExactCount2(counts, 3) {
		return ThreeOfAKind
	} else if hasExactCount2(counts, 2) {
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

type Hand2 struct {
	t     Type
	cards []Card2
	bid   int
}

type Card2 int

const (
	NJ Card2 = iota
	NC2
	NC3
	NC4
	NC5
	NC6
	NC7
	NC8
	NC9
	NT
	NQ
	NK
	NA
)

func parseCards2(c string) []Card2 {
	var cards []Card2
	for _, card := range c {
		var ct Card2
		switch card {
		case '2':
			ct = NC2
		case '3':
			ct = NC3
		case '4':
			ct = NC4
		case '5':
			ct = NC5
		case '6':
			ct = NC6
		case '7':
			ct = NC7
		case '8':
			ct = NC8
		case '9':
			ct = NC9
		case 'T':
			ct = NT
		case 'J':
			ct = NJ
		case 'Q':
			ct = NQ
		case 'K':
			ct = NK
		case 'A':
			ct = NA
		default:
			panic("bad card")
		}
		cards = append(cards, ct)
	}
	return cards
}

func hasExactCount2(counts map[Card2]int, n int) bool {
	for _, count := range counts {
		if count == n {
			return true
		}
	}
	return false
}
