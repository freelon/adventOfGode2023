package day21

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestPart1(t *testing.T) {
	util.AssertEq(t, 1+2+4+8, nOnes(4))
	util.Assert(t, 16, solve1(input, 6))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
