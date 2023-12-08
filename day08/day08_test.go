package day08

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 2, Part1(input))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
