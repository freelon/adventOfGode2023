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
		scorePart := strings.Split(line, ": ")[1]
		scores := strings.Split(scorePart, " | ")
		winning := numbers(scores[0])
		youHave := numbers(scores[1])
		gameScore := 0
		for _, number := range youHave {
			if slices.Contains(winning, number) {
				if gameScore == 0 {
					gameScore = 1
				} else {
					gameScore *= 2
				}
			}
		}
		score += gameScore
	}
	return strconv.Itoa(score)
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
	return ""
}
