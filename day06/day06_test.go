package day06

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `Time:      7  15   30
Distance:  9  40  200`

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 288, Part1(input))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
