package day01

import (
	"adventOfGode2023/util"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	util.Assert(t, 142, Part1(input))
}

func TestDay1Part2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	util.Assert(t, 281, Part2(input))
}
