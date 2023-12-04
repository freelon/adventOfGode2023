package day04

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		score += winScore(line)
	}
	return strconv.Itoa(score)
}

func winScore(line string) int {
	gameWins := wins(line)
	if gameWins > 0 {
		return 1 << (gameWins - 1) // start with 1 for 1st win, double for each next win
	} else {
		return 0
	}
}

func wins(line string) int {
	scorePart := strings.Split(line, ": ")[1]
	scores := strings.Split(scorePart, " | ")
	winning := numbers(scores[0])
	youHave := numbers(scores[1])
	result := 0
	for _, number := range youHave {
		if slices.Contains(winning, number) {
			result += 1
		}
	}
	return result
}

func numbers(s string) (result []string) {
	for _, n := range strings.Split(s, " ") {
		if len(n) > 0 {
			result = append(result, n)
		}
	}
	return
}

func Part2(input string) string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	cards := 0
	for game := range lines {
		cards += count(game, lines)
	}
	return strconv.Itoa(cards)
}

var countCache = map[int]int{}

func count(game int, lines []string) (cards int) {
	v, ok := countCache[game]
	if ok {
		return v
	}
	cards = 1
	gameWins := wins(lines[game])
	for nextGame := game + 1; nextGame < game+gameWins+1; nextGame++ {
		cards += count(nextGame, lines)
	}
	countCache[game] = cards
	return
}
