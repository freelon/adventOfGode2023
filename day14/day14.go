package day14

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	var rows [][]rune
	for _, row := range strings.Split(input, "\n") {
		rows = append(rows, []rune(row))
	}
	for x := 0; x < len(rows[0]); x++ {
		for y := 1; y < len(rows); y++ {
			if y == 0 {
				continue
			}
			if rows[y][x] == 'O' && rows[y-1][x] == '.' {
				rows[y][x] = '.'
				rows[y-1][x] = 'O'
				y -= 2
			}
		}
	}
	h := len(rows)
	weight := 0
	for x := 0; x < len(rows[0]); x++ {
		for y := 0; y < len(rows); y++ {
			if rows[y][x] == 'O' {
				weight += h - y
			}
		}
	}
	return strconv.Itoa(weight)
}

func Part2(input string) string {
	return ""
}
