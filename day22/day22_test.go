package day22

import (
	"adventOfGode2023/util"
	"testing"
)

const input = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func TestPart1(t *testing.T) {
	util.Assert(t, 5, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 7, Part2(input))
}
