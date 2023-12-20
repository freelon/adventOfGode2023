package day20

import (
	"adventOfGode2023/util"
	"testing"
)

const inputA = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func TestPart1a(t *testing.T) {
	util.Assert(t, 32000000, Part1(inputA))
}
func TestPart1b(t *testing.T) {
	util.Assert(t, 11687500, Part1("broadcaster -> a\n%a -> inv, con\n&inv -> b\n%b -> con\n&con -> output"))
}

//func TestPart2(t *testing.T) {
//	util.Assert(t, nil, Part2(input))
//}
