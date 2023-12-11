package day11

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	var rows []string
	{
	}
	for _, l := range strings.Split(input, "\n") {
		rows = append(rows, l)
		if len(l) == strings.Count(l, ".") {
			rows = append(rows, l)
		}
	}
COL:
	for col := 0; ; col++ {
		if col == len(rows[0]) {
			break
		}

		for row := 0; row < len(rows); row++ {
			if rows[row][col] != '.' {
				continue COL
			}
		}
		// column has no galaxies
		for row := 0; row < len(rows); row++ {
			r := rows[row]
			rows[row] = r[0:col] + "." + r[col:]
		}
		col++
	}
	var galaxies []P
	for row := 0; row < len(rows); row++ {
		for col := 0; col < len(rows[row]); col++ {
			if rows[row][col] == '#' {
				galaxies = append(galaxies, P{x: col, y: row})
			}
		}
	}
	distanceSum := 0
	for i, a := range galaxies {
		for _, b := range galaxies[i:] {
			distanceSum += absD(a.x, b.x) + absD(a.y, b.y)
		}
	}
	return strconv.Itoa(distanceSum)
}

type P struct {
	x int
	y int
}

func absD(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func Part2(input string) string {
	return ""
}
