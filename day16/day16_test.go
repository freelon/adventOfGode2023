package day16

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestPart1(t *testing.T) {
	util.Assert(t, 46, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 51, Part2(input))
}
