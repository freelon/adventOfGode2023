package day08

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	movements, desertMap := parse(input)
	count := 0
	node := "AAA"
	for {
		if node == "ZZZ" {
			break
		}
		m := movements[count%len(movements)]
		count++
		if m == 'L' {
			node = desertMap[node].l
		} else if m == 'R' {
			node = desertMap[node].r
		}
	}
	return strconv.Itoa(count)
}

func parse(input string) (movements []rune, desertMap map[string]Pair) {
	regex := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
	parts := strings.Split(input, "\n\n")
	for _, r := range parts[0] {
		movements = append(movements, r)
	}
	desertMap = make(map[string]Pair)
	for _, line := range strings.Split(parts[1], "\n") {
		matches := regex.FindStringSubmatch(line)
		p := Pair{
			l: matches[2],
			r: matches[3],
		}
		desertMap[matches[1]] = p
	}
	return
}

type Pair struct {
	l string
	r string
}

func Part2(input string) string {
	return ""
}
