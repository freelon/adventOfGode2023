package day07

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	util.Assert(t, 6440, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}
