package main

import "testing"

func TestDay1Part1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	Assert(t, 142, Day1Part1(input))
}

func TestDay1Part2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	Assert(t, 281, Day1Part2(input))
}
