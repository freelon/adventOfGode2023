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
	util.Assert(t, 1030, p2(input, 10))
	util.Assert(t, 8410, p2(input, 100))
}
