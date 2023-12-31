package day17

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	h := len(strings.Split(input, "\n"))
	w := len(strings.Split(input, "\n")[0])
	city := parse(input, w, h)
	start1 := Seen{1, 0, right, 1}
	start2 := Seen{0, 1, down, 1}
	open := &MyHeap[Seen]{HeapElement[Seen]{city[start1.x][start1.y], start1},
		HeapElement[Seen]{city[start2.x][start2.y], start2}}
	seen := map[Seen]bool{}
	seen[start1] = true
	seen[start2] = true
	heap.Init(open)
	for len(*open) > 0 {
		current := heap.Pop(open).(HeapElement[Seen])
		if current.value.x == w-1 && current.value.y == h-1 {
			return strconv.Itoa(current.key)
		}
		for _, reachable := range current.value.reachable() {
			// if within bounds ... add to heap
			if reachable.x < 0 || reachable.x >= w || reachable.y < 0 || reachable.y >= h {
				continue
			}
			if _, ok := seen[reachable]; ok {
				continue
			}
			seen[reachable] = true
			heap.Push(open, HeapElement[Seen]{current.key + city[reachable.x][reachable.y], reachable})
		}
	}
	panic("didn't find a path at all, wtf")
}

type Seen struct {
	x, y  int
	from  direction
	steps int
}

func (s Seen) reachable() []Seen {
	result := make([]Seen, 0)
	for _, d := range []direction{up, down, left, right} {
		if s.from.opposite() == d {
			continue
		}
		if s.from == d {
			if s.steps == 3 {
				continue
			} else {
				dx, dy := d.movement()
				result = append(result,
					Seen{
						x:     s.x + dx,
						y:     s.y + dy,
						from:  d,
						steps: s.steps + 1,
					})
				continue
			}
		}
		// turn
		dx, dy := d.movement()
		result = append(result, Seen{
			x:     s.x + dx,
			y:     s.y + dy,
			from:  d,
			steps: 1,
		})
	}
	return result
}

type direction int

func (d direction) String() string {
	return [...]string{"up",
		"down",
		"left",
		"right"}[d]
}

func (d direction) opposite() direction {
	switch d {
	case down:
		return up
	case up:
		return down
	case left:
		return right
	case right:
		return left
	default:
		panic("unreachable")
	}
}

func (d direction) movement() (dx int, dy int) {
	switch d {
	case down:
		return 0, 1
	case up:
		return 0, -1
	case left:
		return -1, 0
	case right:
		return 1, 0
	default:
		panic("unreachable")
	}
}

const (
	up direction = iota
	down
	left
	right
)

func parse(input string, w int, h int) [][]int {
	city := make([][]int, w)
	for x := 0; x < w; x++ {
		city[x] = make([]int, h)
	}
	for y, row := range strings.Split(input, "\n") {
		for x, s := range strings.Split(row, "") {
			i, _ := strconv.Atoi(s)
			city[x][y] = i
		}
	}
	return city
}

func Part2(input string) string {
	h := len(strings.Split(input, "\n"))
	w := len(strings.Split(input, "\n")[0])
	city := parse(input, w, h)
	start1 := SeenUltra{1, 0, right, 1}
	start2 := SeenUltra{0, 1, down, 1}
	open := &MyHeap[SeenUltra]{HeapElement[SeenUltra]{city[start1.x][start1.y], start1},
		HeapElement[SeenUltra]{city[start2.x][start2.y], start2}}
	seen := map[SeenUltra]bool{}
	seen[start1] = true
	seen[start2] = true
	heap.Init(open)
	for len(*open) > 0 {
		current := heap.Pop(open).(HeapElement[SeenUltra])
		if current.value.x == w-1 && current.value.y == h-1 {
			if current.value.steps >= 4 {
				// special requirement: the ultra crubicle has to move 4 steps in the same direction before it can stop
				return strconv.Itoa(current.key)
			} else {
				fmt.Println("so sad")
			}
		}
		for _, reachable := range current.value.reachable() {
			// if within bounds ... add to heap
			if reachable.x < 0 || reachable.x >= w || reachable.y < 0 || reachable.y >= h {
				continue
			}
			if _, ok := seen[reachable]; ok {
				continue
			}
			seen[reachable] = true
			heap.Push(open, HeapElement[SeenUltra]{current.key + city[reachable.x][reachable.y], reachable})
		}
	}
	panic("didn't find a path at all, wtf")
}

type HeapElement[T any] struct {
	key   int
	value T
}

// An MyHeap is a min-heap of where the keys are ints.
type MyHeap[T any] []HeapElement[T]

func (h MyHeap[T]) Len() int           { return len(h) }
func (h MyHeap[T]) Less(i, j int) bool { return h[i].key < h[j].key }
func (h MyHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(HeapElement[T]))
}

func (h *MyHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeenUltra struct {
	x, y  int
	from  direction
	steps int
}

func (s SeenUltra) reachable() []SeenUltra {
	result := make([]SeenUltra, 0)
	for _, d := range []direction{up, down, left, right} {
		if s.from.opposite() == d {
			continue
		}
		if s.from == d {
			if s.steps == 10 {
				continue
			} else {
				dx, dy := d.movement()
				result = append(result,
					SeenUltra{
						x:     s.x + dx,
						y:     s.y + dy,
						from:  d,
						steps: s.steps + 1,
					})
				continue
			}
		}
		// turn
		if s.steps < 4 {
			continue
		}
		dx, dy := d.movement()
		result = append(result, SeenUltra{
			x:     s.x + dx,
			y:     s.y + dy,
			from:  d,
			steps: 1,
		})
	}
	return result
}
