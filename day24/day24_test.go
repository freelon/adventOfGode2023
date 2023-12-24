package day24

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `19, 13, 30 @ -2, 1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @ 1, -5, -3`

func TestPart1(t *testing.T) {

	util.AssertEq(t, 2, solve(parse(input), 7, 27))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
