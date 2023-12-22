package day21

import (
	"adventOfGode2023/util"
	"strings"
	"testing"
)

const input = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestPart1(t *testing.T) {
	util.Assert(t, 16, solve1(input, 6))
}

func TestBlank10(t *testing.T) {
	native := solve1(makeGrid(101), 10)
	util.Assert(t, native, solve2(makeGrid(11), 10))
}

func TestBlank20(t *testing.T) {
	native := solve1(makeGrid(101), 20)
	util.Assert(t, native, solve2(makeGrid(11), 20))
}

func TestBlank30(t *testing.T) {
	native := solve1(makeGrid(101), 30)
	util.Assert(t, native, solve2(makeGrid(11), 30))
}

func TestBlank40(t *testing.T) {
	native := solve1(makeGrid(101), 40)
	util.Assert(t, native, solve2(makeGrid(11), 40))
}

func TestBlank50(t *testing.T) {
	native := solve1(makeGrid(101), 50)
	util.Assert(t, native, solve2(makeGrid(11), 50))
}

func makeGrid(L int) (result string) {
	if L%2 == 0 {
		panic("need uneven side length")
	}
	for i := 0; i < L; i++ {
		s := strings.Repeat(".", L)
		if i == ((L - 1) / 2) {
			x := []rune(s)
			x[(L-1)/2] = 'S'
			s = string(x)
		}
		result = result + s + "\n"
	}
	result = strings.TrimSpace(result)
	return
}
