package day11

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	util.Assert(t, 374, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
