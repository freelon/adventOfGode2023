package day23

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	tiles := make(map[Pos]rune)
	var w, h int
	for y, line := range strings.Split(input, "\n") {
		h = y + 1
		for x, r := range line {
			tiles[Pos{x, y}] = r
			w = x + 1
		}
	}
	start := Pos{1, 0}
	goal := Pos{x: w - 2, y: h - 1}
	visited := make(map[Pos]bool)
	next := start
	result, found := longestPath(next, goal, visited, tiles, false)
	if !found {
		panic("#spiderman #nowayhome")
	}
	//printPathInMap(h, w, result, tiles)
	return strconv.Itoa(len(result))
}

func printPathInMap(h int, w int, result map[Pos]bool, tiles map[Pos]rune) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if _, ok := result[Pos{x, y}]; ok {
				fmt.Print("O")
			} else {
				fmt.Print(string(tiles[Pos{x, y}]))
			}
		}
		fmt.Println()
	}
}

func longestPath(current Pos, goal Pos, visited map[Pos]bool, tiles map[Pos]rune, ignoreSlopes bool) (visitedAlongPath map[Pos]bool, foundPath bool) {
	for {
		if current == goal {
			return visited, true
		}
		if _, contains := visited[current]; contains {
			panic("should not happen")
		}
		visited[current] = true
		currentTile, _ := tiles[current]
		var candidates []Pos
		if ignoreSlopes {
			candidates = current.neighbors()
		} else {
			switch currentTile {
			case '.':
				candidates = current.neighbors()
			case '<':
				candidates = append(candidates, current.left())
			case '>':
				candidates = append(candidates, current.right())
			case 'v':
				candidates = append(candidates, current.down())
			case '^':
				candidates = append(candidates, current.up())
			}
		}
		candidates = slices.DeleteFunc(candidates, func(candidate Pos) bool {
			tile, ok := tiles[candidate]
			if !ok {
				return true
			}
			if tile == '#' {
				return true
			}
			if _, contains := visited[candidate]; contains {
				return true
			}
			return false
		})
		if len(candidates) == 0 {
			return nil, false
		} else if len(candidates) == 1 {
			current = candidates[0]
		} else {
			for _, candidate := range candidates {
				newVisited := maps.Clone(visited)
				l, o := longestPath(candidate, goal, newVisited, tiles, ignoreSlopes)
				if o && len(l) > len(visitedAlongPath) {
					visitedAlongPath = l
					foundPath = true
				}
			}
			return
		}
	}
}

type Pos struct {
	x, y int
}

func (p Pos) neighbors() []Pos {
	return []Pos{p.left(), p.right(), p.up(), p.down()}
}

func (p Pos) left() Pos {
	return Pos{p.x - 1, p.y}
}
func (p Pos) right() Pos {
	return Pos{p.x + 1, p.y}
}
func (p Pos) up() Pos {
	return Pos{p.x, p.y - 1}
}
func (p Pos) down() Pos {
	return Pos{p.x, p.y + 1}
}

func Part2(input string) string {
	tiles := make(map[Pos]rune)
	var w, h int
	for y, line := range strings.Split(input, "\n") {
		h = y + 1
		for x, r := range line {
			tiles[Pos{x, y}] = r
			w = x + 1
		}
	}
	var junctions []Pos
	for k, v := range tiles {
		if v == '#' {
			continue
		}
		var candidates = slices.DeleteFunc(k.neighbors(), func(candidate Pos) bool {
			tile, ok := tiles[candidate]
			if !ok {
				return true
			}
			if tile == '#' {
				return true
			}
			return false
		})
		if len(candidates) > 2 {
			junctions = append(junctions, k)
		}
	}
	fmt.Printf("Found %d junctions.\n", len(junctions))
	slices.SortFunc(junctions, func(a, b Pos) int {
		if a.x == b.x {
			return b.y - a.y
		} else {
			return b.x - a.x
		}
	})

	start := Pos{1, 0}
	goal := Pos{x: w - 2, y: h - 1}
	poi := []Pos{start, goal}
	poi = append(poi, junctions...)
	am := make([][]int, len(poi))
	for i := 0; i < len(poi); i++ {
		am[i] = make([]int, len(poi))
	}
	// close all junctions

	for _, pos := range poi {
		tiles[pos] = '#'
	}

	for s := 0; s < len(poi); s++ {
		am[s][s] = -1
		start = poi[s]
		for e := s + 1; e < len(poi); e++ {
			am[s][e] = -1
			am[e][s] = -1

			end := poi[e]
			tiles[start] = '.'
			tiles[end] = '.'
			visited := make(map[Pos]bool)
			l, found := longestPath(start, end, visited, tiles, true)
			if found {
				am[s][e] = len(l)
				am[e][s] = len(l)
			}
			tiles[start] = '#'
			tiles[end] = '#'
		}
	}

	visited := make([]bool, len(poi))
	result, _, ok := longestPathGraph(0, 1, visited, am)
	if !ok {
		panic("no way home :(")
	}

	return strconv.Itoa(result)
}

func longestPathGraph(current int, goal int, visited []bool, am [][]int) (int, []int, bool) {
	if current == goal {
		return 0, []int{current}, true
	}
	if visited[current] {
		panic("should not happen")
	}
	visited[current] = true

	foundOne := false
	longest := 0
	pathOfLongest := make([]int, 0)
	for candidate := 0; candidate < len(visited); candidate++ {
		if current == candidate || visited[candidate] || am[current][candidate] < 0 {
			continue
		}
		l, p, f := longestPathGraph(candidate, goal, slices.Clone(visited), am)
		if f {
			foundOne = true
			d := am[current][candidate] + l
			if d > longest {
				longest = d
				pathOfLongest = p
			}
		}
	}
	if foundOne {
		pathOfLongest = append(pathOfLongest, current)
		return longest, pathOfLongest, true
	} else {
		return 0, make([]int, 0), false
	}
}
