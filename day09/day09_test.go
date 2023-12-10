package day09

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 114, Part1(input))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
