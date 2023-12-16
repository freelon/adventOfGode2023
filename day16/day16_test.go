package day16

import (
	day16 "adventOfGode2023/dayTemplate"
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
	util.Assert(t, 46, day16.Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, nil, day16.Part2(input))
}
