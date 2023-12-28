package day21

import (
	"adventOfGode2023/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	return solve1(input, 64)
}

func solve1(input string, remainingSteps int) string {
	garden, start := parseGarden(input)
	count := count(start, garden, remainingSteps)
	return strconv.Itoa(count)
}

func parseGarden(input string) (map[Pos]rune, Pos) {
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
	return garden, start
}

type Progress struct {
	p Pos
	d int
}

func count(start Pos, garden map[Pos]rune, remainingSteps int) (reachableCount int) {
	if remainingSteps < 0 {
		return 0
	}
	visited := markFields(start, garden, remainingSteps)
	for _, v := range visited {
		if v%2 == remainingSteps%2 {
			reachableCount++
		}
	}
	return
}

func markFields(start Pos, garden map[Pos]rune, remainingSteps int) map[Pos]int {
	var queue util.Queue[Progress]
	queue.Enqueue(Progress{p: start, d: 0})
	visited := make(map[Pos]int)
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
		if next.d > remainingSteps {
			continue
		}
		visited[next.p] = next.d

		for _, neighbor := range next.p.neighbors() {
			queue.Enqueue(Progress{neighbor, next.d + 1})
		}
	}
	return visited
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

func Part2(input string) string {
	return solve2(input, 26501365)
}

func solve2(input string, remainingSteps int) string {
	L := len(strings.Split(input, "\n")[0])
	l := (L - 1) / 2
	garden, startCenter := parseGarden(input)
	center := count(startCenter, garden, remainingSteps)

	startForWest := Pos{x: L - 1, y: startCenter.y}
	black := count(startForWest, garden, 2*L)
	white := count(startForWest, garden, 2*L+1)

	startForEast := Pos{x: 0, y: startCenter.y}
	startForNorth := Pos{x: startCenter.x, y: L - 1}
	startForSouth := Pos{x: startCenter.x, y: 0}
	startForNorthEast := Pos{x: 0, y: L - 1}
	startForNorthWest := Pos{x: L - 1, y: L - 1}
	startForSouthEast := Pos{x: 0, y: 0}
	startForSouthWest := Pos{x: L - 1, y: 0}

	N := max(0, (remainingSteps-(l+1))/L-1)
	nLast := count(startForNorth, garden, remainingSteps-(l+1)-(N+1)*L)
	nSecondLast := count(startForNorth, garden, remainingSteps-(l+1)-N*L)
	sLast := count(startForSouth, garden, remainingSteps-(l+1)-(N+1)*L)
	sSecondLast := count(startForSouth, garden, remainingSteps-(l+1)-N*L)
	wLast := count(startForWest, garden, remainingSteps-(l+1)-(N+1)*L)
	wSecondLast := count(startForWest, garden, remainingSteps-(l+1)-N*L)
	eLast := count(startForEast, garden, remainingSteps-(l+1)-(N+1)*L)
	eSecondLast := count(startForEast, garden, remainingSteps-(l+1)-N*L)
	xFirstCount := N/2 + N%2
	xSecondCount := N / 2

	blocks := max(0, N-1)
	blackCount := (blocks + blocks%2) * (blocks + blocks%2) / 4
	whiteCount := (blocks*blocks+blocks)/2 - blackCount

	neThirdLast := count(startForNorthEast, garden, remainingSteps-(N)*L-1)
	neSecondLast := count(startForNorthEast, garden, remainingSteps-(N+1)*L-1)
	neLast := count(startForNorthEast, garden, remainingSteps-(N+2)*L-1)
	nwThirdLast := count(startForNorthWest, garden, remainingSteps-(N)*L-1)
	nwSecondLast := count(startForNorthWest, garden, remainingSteps-(N+1)*L-1)
	nwLast := count(startForNorthWest, garden, remainingSteps-(N+2)*L-1)
	seThirdLast := count(startForSouthEast, garden, remainingSteps-(N)*L-1)
	seSecondLast := count(startForSouthEast, garden, remainingSteps-(N+1)*L-1)
	seLast := count(startForSouthEast, garden, remainingSteps-(N+2)*L-1)
	swThirdLast := count(startForSouthWest, garden, remainingSteps-(N)*L-1)
	swSecondLast := count(startForSouthWest, garden, remainingSteps-(N+1)*L-1)
	swLast := count(startForSouthWest, garden, remainingSteps-(N+2)*L-1)

	fmt.Printf("N = %d\n", N)
	fmt.Printf("blocks = %d\n", blocks)
	fmt.Printf("black = %d, white = %d\n", blackCount, whiteCount)

	// [nx][neLast]
	// ...
	// [n2][ne2][ne1]...[ne2][neSecondLast][neLast]
	// [n1][ne1][ne2]...[ne1][ne2][neSecondLast][neLast]
	// [S ][e1 ][e2 ]...[e1 ][e2 ][eSecondLast ][eLast ]

	all := center +
		4*black*xFirstCount +
		4*white*xSecondCount +
		wSecondLast +
		wLast +
		eSecondLast +
		eLast +
		nSecondLast +
		nLast +
		sSecondLast +
		sLast +
		4*black*blackCount +
		4*white*whiteCount +
		neThirdLast*N +
		neSecondLast*(N+1) +
		neLast*(N+2) +
		nwThirdLast*N +
		nwSecondLast*(N+1) +
		nwLast*(N+2) +
		seThirdLast*N +
		seSecondLast*(N+1) +
		seLast*(N+2) +
		swThirdLast*N +
		swSecondLast*(N+1) +
		swLast*(N+2)
	return strconv.Itoa(all)
}
