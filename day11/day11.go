package day11

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	var rows []string
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
	return p2(input, 1_000_000)
}

func p2(input string, expansionFactor int) string {
	var rows []string
	var emptyRows []int
	for row, l := range strings.Split(input, "\n") {
		rows = append(rows, l)
		if len(l) == strings.Count(l, ".") {
			emptyRows = append(emptyRows, row)
		}
	}
	var emptyCols []int
COL:
	for col := 0; col < len(rows[0]); col++ {
		for row := 0; row < len(rows); row++ {
			if rows[row][col] != '.' {
				continue COL
			}
		}
		emptyCols = append(emptyCols, col)
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
			for _, col := range emptyCols {
				if col > min(a.x, b.x) && col < max(a.x, b.x) {
					distanceSum = distanceSum - 1 + expansionFactor
				}
			}
			for _, row := range emptyRows {
				if row > min(a.y, b.y) && row < max(a.y, b.y) {
					distanceSum = distanceSum - 1 + expansionFactor
				}
			}
		}
	}
	return strconv.Itoa(distanceSum)
}
