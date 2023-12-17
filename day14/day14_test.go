package day14

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPart1(t *testing.T) {
	util.Assert(t, 136, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 64, Part2(input))
}
