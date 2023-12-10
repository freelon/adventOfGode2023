package day10

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	pipes, start := parse(input)
	// find two directions that are open to start
	openToStarts := make([]Pos, 0)
	if openToNorth(pipes[start.south()]) {
		openToStarts = append(openToStarts, start.south())
	}
	if openToSouth(pipes[start.north()]) {
		openToStarts = append(openToStarts, start.north())
	}
	if openToEast(pipes[start.west()]) {
		openToStarts = append(openToStarts, start.west())
	}
	if openToWest(pipes[start.east()]) {
		openToStarts = append(openToStarts, start.east())
	}
	fmt.Printf("search for path from %s to %s\n", openToStarts[0], openToStarts[1])
	sp := shortestPath(pipes, openToStarts[0], openToStarts[1])
	highestDistance := (len(sp) + 1) / 2
	return strconv.Itoa(highestDistance)
}

func shortestPath(pipes map[Pos]rune, from Pos, to Pos) []Pos {
	open := make([]Pos, 0)
	open = append(open, from)
	visited := make(map[Pos]bool)
	predecessor := make(map[Pos]Pos)
	for {
		current := open[0]
		open = open[1:]
		_, ok := visited[current]
		if ok {
			continue
		}
		visited[current] = true

		if current == to {
			break
		}

		connectedNeighbors := getConnectedNeighbors(pipes, current)

		for _, neighbor := range connectedNeighbors {
			_, ok := predecessor[neighbor]
			if !ok {
				predecessor[neighbor] = current
			}
			_, ok = visited[neighbor]
			if !ok {
				open = append(open, neighbor)
			}
		}
	}
	result := make([]Pos, 0)
	result = append(result, to)
	at := to
	for {
		if at == from {
			break
		}
		next, ok := predecessor[at]
		at = next
		if ok {
			result = slices.Insert(result, 0, at)
		} else {
			panic(fmt.Sprintf("missing predecessor for %s", at))
		}
	}
	return result
}

func getConnectedNeighbors(pipes map[Pos]rune, current Pos) []Pos {
	result := make([]Pos, 0)
	if openToNorth(pipes[current]) && openToSouth(pipes[current.north()]) {
		result = append(result, current.north())
	}
	if openToSouth(pipes[current]) && openToNorth(pipes[current.south()]) {
		result = append(result, current.south())
	}
	if openToWest(pipes[current]) && openToEast(pipes[current.west()]) {
		result = append(result, current.west())
	}
	if openToEast(pipes[current]) && openToWest(pipes[current.east()]) {
		result = append(result, current.east())
	}
	return result
}

type Pos struct {
	e int
	s int
}

func (p Pos) String() string {
	return fmt.Sprintf("e: %d, s: %d", p.e, p.s)
}

func (p Pos) north() Pos {
	return Pos{e: p.e, s: p.s - 1}
}

func (p Pos) south() Pos {
	return Pos{e: p.e, s: p.s + 1}
}

func (p Pos) east() Pos {
	return Pos{e: p.e + 1, s: p.s}
}

func (p Pos) west() Pos {
	return Pos{e: p.e - 1, s: p.s}
}

func openToNorth(r rune) bool {
	switch r {
	case '|', 'L', 'J':
		return true
	default:
		return false
	}
}

func openToSouth(r rune) bool {
	switch r {
	case '|', '7', 'F':
		return true
	default:
		return false
	}
}

func openToEast(r rune) bool {
	switch r {
	case '-', 'L', 'F':
		return true
	default:
		return false
	}
}

func openToWest(r rune) bool {
	switch r {
	case '-', 'J', '7':
		return true
	default:
		return false
	}
}

func parse(input string) (pipes map[Pos]rune, start Pos) {
	pipes = make(map[Pos]rune)
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			pipes[Pos{x, y}] = r
			if r == 'S' {
				start = Pos{x, y}
			}
		}
	}
	return
}

func Part2(input string) string {
	pipes, start := parse(input)
	// find two directions that are open to start
	openToStarts := make([]Pos, 0)
	if openToNorth(pipes[start.south()]) {
		openToStarts = append(openToStarts, start.south())
	}
	if openToSouth(pipes[start.north()]) {
		openToStarts = append(openToStarts, start.north())
	}
	if openToEast(pipes[start.west()]) {
		openToStarts = append(openToStarts, start.west())
	}
	if openToWest(pipes[start.east()]) {
		openToStarts = append(openToStarts, start.east())
	}
	fmt.Printf("search for path from %s to %s\n", openToStarts[0], openToStarts[1])
	sp := shortestPath(pipes, openToStarts[1], openToStarts[0])
	sp = append(sp, start)
	s := shoelace(sp)
	if s < 0 {
		s *= -1
	}
	inners := s - len(sp)/2 + 1 // pick's theorem wtf
	return strconv.Itoa(inners)
}

func shoelace(path []Pos) int {
	n := len(path)
	sumA := 0
	for i := 0; i < n; i++ {
		sumA += path[i].e * path[(i+1)%n].s
	}
	sumB := 0
	for i := 0; i < n; i++ {
		sumB += path[i].s * path[(i+1)%n].e
	}
	return (sumA - sumB) / 2
}
