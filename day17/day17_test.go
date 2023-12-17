package day17

import (
	"adventOfGode2023/util"
	"strconv"
	"testing"
)

const input = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestPart1(t *testing.T) {
	util.Assert(t, 102, Part1(input))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 94, Part2(input))
	util.Assert(t, 71, Part2("111111111111\n999999999991\n999999999991\n999999999991\n999999999991"))
}

func TestReachable(t *testing.T) {
	s := SeenUltra{
		x:     0,
		y:     0,
		from:  down,
		steps: 4,
	}
	util.Assert(t, 3, strconv.Itoa(len(s.reachable())))
}
