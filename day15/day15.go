package day15

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	parts := strings.Split(input, ",")
	sum := 0
	for _, part := range parts {
		sum += hash(part)
	}
	return strconv.Itoa(sum)
}

func hash(part string) (result int) {
	for _, r := range part {
		result += int(r)
		result *= 17
		result %= 256
	}
	return
}

func Part2(input string) string {
	return ""
}
