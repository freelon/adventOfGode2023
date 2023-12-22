package day03

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) string {
	schematic := parseSchematic(input)
	partNumbers := partNumbers(schematic)

	sum := 0
	for _, partNumber := range partNumbers {
		if hasSymbolNeighbor(partNumber, schematic) {
			sum += partNumber.value
			//println(partNumber.value)
		}
	}
	return strconv.Itoa(sum)
}

func hasSymbolNeighbor(partNumber PartNumber, schematic Schematic) bool {
	for y := partNumber.y - 1; y < partNumber.y+2; y++ {
		for x := partNumber.x - 1; x < partNumber.x+partNumber.l+1; x++ {
			if y < 0 || y >= len(schematic) || x < 0 || x >= len(schematic[y]) {
				continue
			}
			r := schematic[y][x]
			if !unicode.IsDigit(r) && r != '.' {
				return true
			}
		}
	}
	return false
}

func partNumbers(schematic Schematic) []PartNumber {
	var result []PartNumber
	for y, line := range schematic {
		var xStart, l int
		var in = false
		var value = ""
		for x, r := range line {
			if in {
				if unicode.IsDigit(r) {
					l += 1
					value = fmt.Sprintf("%s%c", value, r)
				} else {
					v, _ := strconv.Atoi(value)
					p := PartNumber{
						value: v,
						x:     xStart,
						y:     y,
						l:     l,
					}
					result = append(result, p)
					in = false
				}
			} else {
				if unicode.IsDigit(r) {
					xStart = x
					in = true
					value = fmt.Sprintf("%c", r)
					l = 1
				} else {

				}
			}
		}
		if in {
			v, _ := strconv.Atoi(value)
			p := PartNumber{
				value: v,
				x:     xStart,
				y:     y,
				l:     l,
			}
			result = append(result, p)
			in = false
		}
	}
	return result
}

func parseSchematic(input string) Schematic {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var schematic Schematic
	for _, line := range lines {
		var schemaLine []rune
		for _, r := range line {
			schemaLine = append(schemaLine, r)
		}
		schematic = append(schematic, schemaLine)
	}
	return schematic
}

func Part2(input string) string {
	schematic := parseSchematic(input)
	partNumbers := partNumbers(schematic)
	stars := stars(schematic)
	sum := 0
	for _, star := range stars {
		var parts []PartNumber
		for _, partNumber := range partNumbers {
			if isAdjacent(star, partNumber) {
				parts = append(parts, partNumber)
			}
		}
		if len(parts) == 2 {
			sum += parts[0].value * parts[1].value
		}
	}
	return strconv.Itoa(sum)
}

func isAdjacent(point Point, number PartNumber) bool {
	col := point.x >= number.x-1 && point.x <= number.x+number.l
	row := point.y >= number.y-1 && point.y <= number.y+1
	return col && row
}

func stars(schematic Schematic) (stars []Point) {
	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			if schematic[y][x] == '*' {
				stars = append(stars, Point{x, y})
			}
		}
	}
	return
}

type Point struct {
	x, y int
}

type PartNumber struct {
	value   int
	x, y, l int
}

// Schematic  [y][x] -> rune
type Schematic = [][]rune
