package day03

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 4361, Part1(input))
}

func TestSymbolCornerCases(t *testing.T) {
	input := "..*1"
	schematic := parseSchematic(input)
	parts := partNumbers(schematic)
	util.AssertEq(t, true, hasSymbolNeighbor(parts[0], schematic))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, 2286, Part2(input))
}
