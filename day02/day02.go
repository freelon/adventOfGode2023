package day02

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
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

func Part2(input string) string {
	input = strings.TrimSpace(input)
	powerSum := 0
	for _, line := range strings.Split(input, "\n") {
		maxCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		split := strings.Split(line, ":")
		_, rhs := split[0], split[1]
		outtakes := strings.Split(strings.TrimSpace(rhs), "; ")
		for _, outtake := range outtakes {
			outtake := strings.TrimSpace(outtake)
			parts := strings.Split(outtake, ", ")
			for _, part := range parts {
				partSplit := strings.Split(part, " ")
				n, _ := strconv.Atoi(partSplit[0])
				color := partSplit[1]
				if n > maxCubes[color] {
					maxCubes[color] = n
				}
			}
		}
		powerSum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
	}
	return strconv.Itoa(powerSum)
}
