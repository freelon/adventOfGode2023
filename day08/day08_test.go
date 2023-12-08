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

const input2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestDay2Part2(t *testing.T) {
	util.Assert(t, 6, Part2(input2))
}
