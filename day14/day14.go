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
	var rows [][]rune
	for _, row := range strings.Split(input, "\n") {
		rows = append(rows, []rune(row))
	}
	var cache = make(map[string]int)
	MAX := 1000000000
	for i := 0; i < MAX; i++ {
		h := hash(rows)
		if v, ok := cache[h]; ok {
			cycle := i - v
			left := MAX - i
			cycles := left / cycle
			i += cycle * cycles
			cache = make(map[string]int) // lazyness
		} else {
			cache[h] = i
		}
		roll(rows)
		rows = rotate(rows)
		roll(rows)
		rows = rotate(rows)
		roll(rows)
		rows = rotate(rows)
		roll(rows)
		rows = rotate(rows)
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

func hash(rows [][]rune) string {
	result := ""
	for _, row := range rows {
		result += string(row)
	}
	return result
}

// rotate clockwise
func rotate(rows [][]rune) (result [][]rune) {
	result = make([][]rune, len(rows))
	for x := 0; x < len(rows); x++ {
		for y := len(rows) - 1; y >= 0; y-- {
			result[x] = append(result[x], rows[y][x])
		}
	}
	return
}

func roll(rows [][]rune) {
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
}
