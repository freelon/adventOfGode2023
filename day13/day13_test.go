package day13

import (
	"adventOfGode2023/util"
	"strconv"
	"strings"
	"testing"
)

const input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestPart1(t *testing.T) {
	util.Assert(t, 4, strconv.Itoa(reflection(strings.Split("#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#", "\n"), -1)))
	util.Assert(t, 405, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 400, Part2(input))
}
