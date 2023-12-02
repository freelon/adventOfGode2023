package main

import (
	"strconv"
	"strings"
)

func Day2Part1(input string) string {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	input = strings.TrimSpace(input)
	goodGamesIdSum := 0
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ":")
		lhs, rhs := split[0], split[1]
		gameId, _ := strconv.Atoi(strings.Split(lhs, " ")[1])
		outtakes := strings.Split(strings.TrimSpace(rhs), "; ")
		goodGame := true
		for _, outtake := range outtakes {
			outtake := strings.TrimSpace(outtake)
			parts := strings.Split(outtake, ", ")
			for _, part := range parts {
				partSplit := strings.Split(part, " ")
				n, _ := strconv.Atoi(partSplit[0])
				color := partSplit[1]
				if maxCubes[color] < n {
					goodGame = false
				}
			}
		}
		if goodGame {
			goodGamesIdSum += gameId
		}
	}
	return strconv.Itoa(goodGamesIdSum)
}

func Day2Part2(input string) string {
	return ""
}
