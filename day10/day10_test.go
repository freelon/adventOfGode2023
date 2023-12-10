package day10

import (
	"adventOfGode2023/util"
	"testing"
)

const bigExampleFor2 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

func TestDay2Part1(t *testing.T) {
	util.Assert(t, 4, Part1(".....\n.S-7.\n.|.|.\n.L-J.\n....."))
	util.Assert(t, 8, Part1("..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...\n"))
}

func TestDay2Part2(t *testing.T) {
	util.Assert(t, 1, Part2(".....\n.S-7.\n.|.|.\n.L-J.\n.....\n"))
	util.Assert(t, 4, Part2("...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n...........\n"))
	util.Assert(t, 8, Part2(bigExampleFor2))
}

// | is a vertical pipe connecting north and s.
// - is a horizontal pipe connecting e and west.
// L is a 90-degree bend connecting north and e.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting s and west.
// F is a 90-degree bend connecting s and e.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
