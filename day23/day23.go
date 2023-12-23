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
	result, found := longestPath(next, goal, visited, tiles)
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

func longestPath(current Pos, goal Pos, visited map[Pos]bool, tiles map[Pos]rune) (visitedAlongPath map[Pos]bool, foundPath bool) {
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
				l, o := longestPath(candidate, goal, newVisited, tiles)
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

func Part2(_ string) string {
	return ""
}
