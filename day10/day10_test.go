package day10

import (
	"adventOfGode2023/util"
	"testing"
)

const input = ``

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 4, Part1(".....\n.S-7.\n.|.|.\n.L-J.\n....."))
	util.Assert(t, 8, Part1("..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...\n"))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, nil, Part2(input))
}

// | is a vertical pipe connecting north and s.
// - is a horizontal pipe connecting e and west.
// L is a 90-degree bend connecting north and e.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting s and west.
// F is a 90-degree bend connecting s and e.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
