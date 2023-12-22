package day21

import (
	"adventOfGode2023/util"
	"strconv"
	"strings"
)

func Part1(input string) string {
	return solve1(input, 64)
}

func solve1(input string, wantedSteps int) string {
	garden := make(map[Pos]rune)
	var start Pos
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			if r == 'S' {
				r = '.'
				start = Pos{x, y}
			}
			garden[Pos{x, y}] = r
		}
	}
	count := count(start, garden, wantedSteps)
	return strconv.Itoa(count)
}

type Progress struct {
	p Pos
	d int
}

func count(start Pos, garden map[Pos]rune, wantedSteps int) (reachableCount int) {
	var queue util.Queue[Progress]
	queue.Enqueue(Progress{p: start, d: 0})
	visited := make(map[Pos]bool)
	for next, ok := queue.Dequeue(); ok; next, ok = queue.Dequeue() {
		target, ok := garden[next.p]
		if !ok {
			continue
		}
		if target == '#' {
			continue
		}
		if _, ok := visited[next.p]; ok {
			continue
		}
		if next.d > wantedSteps {
			continue
		}
		if next.d%2 == wantedSteps%2 {
			reachableCount++
		}
		visited[next.p] = true

		for _, neighbor := range next.p.neighbors() {
			queue.Enqueue(Progress{neighbor, next.d + 1})
		}
	}
	return
}

func contains(reachable uint64, c int) bool {
	var x uint64 = 1 << (c - 1)
	y := reachable & x
	return y > 0
}

func nOnes(steps int) (result uint64) {
	for i := 0; i < steps; i++ {
		result = result << 1
		result += 1
	}
	return result
}

type Pos struct {
	x, y int
}

func (p Pos) neighbors() []Pos {
	return []Pos{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
}

func Part2(_ string) string {
	return ""
}
