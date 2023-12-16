package day13

import (
	"adventOfGode2023/util"
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
	//util.Assert(t, 4, strconv.Itoa(reflection(strings.Split("#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#", "\n"))))
	util.Assert(t, 405, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
