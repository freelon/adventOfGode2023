package day15

import (
	"adventOfGode2023/util"
	"strconv"
	"testing"
)

const input = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestPart1(t *testing.T) {
	util.Assert(t, 52, strconv.Itoa(hash("HASH")))
	util.Assert(t, 1320, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 145, Part2(input))
}
