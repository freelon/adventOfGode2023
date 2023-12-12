package day12

import (
	"adventOfGode2023/util"
	"strconv"
	"testing"
)

const input = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestPart1(t *testing.T) {
	util.Assert(t, 21, Part1(input))
}

func TestArrangements(t *testing.T) {
	util.Assert(t, 1, strconv.Itoa(arrangements("???.### 1,1,3")))
	util.Assert(t, 4, strconv.Itoa(arrangements(".??..??...?##. 1,1,3")))
	util.Assert(t, 1, strconv.Itoa(arrangements("?#?#?#?#?#?#?#? 1,3,1,6")))
	//.#.###.#.#####? 1,3,1,6
	util.Assert(t, 1, strconv.Itoa(arrangements("????.#...#... 4,1,1")))
	util.Assert(t, 4, strconv.Itoa(arrangements("????.######..#####. 1,6,5")))
	util.Assert(t, 10, strconv.Itoa(arrangements("?###???????? 3,2,1")))
}

func TestPart2(t *testing.T) {
	util.Assert(t, 525152, Part2(input))
}
