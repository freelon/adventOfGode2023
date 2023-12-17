package day17

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	h := len(strings.Split(input, "\n"))
	w := len(strings.Split(input, "\n")[0])
	city := parse(input, w, h)
	start1 := Seen{1, 0, []direction{right}}
	start2 := Seen{0, 1, []direction{down}}
	open := &MyHeap[Seen]{HeapElement[Seen]{city[start1.x][start1.y], start1},
		HeapElement[Seen]{city[start2.x][start2.y], start2}}
	seen := map[string]bool{}
	seen[start1.String()] = true
	seen[start2.String()] = true
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
			s := reachable.String()
			if _, ok := seen[s]; ok {
				continue
			}
			seen[s] = true
			heap.Push(open, HeapElement[Seen]{current.key + city[reachable.x][reachable.y], reachable})
		}
	}
	panic("didn't find a path at all, wtf")
}

type Seen struct {
	x, y int
	from []direction
}

func (s Seen) String() string {
	return fmt.Sprintf("%d,%d,%v", s.x, s.y, s.from)
}

func (s *Seen) equals(o *Seen) bool {
	if s.x != o.x || s.y != o.y || len(s.from) != len(o.from) {
		return false
	}
	for i := 0; i < len(s.from); i++ {
		if s.from[i] != o.from[i] {
			return false
		}
	}
	return true
}

func (s Seen) reachable() []Seen {
	result := []Seen{
		{
			x:    s.x + 1,
			y:    s.y,
			from: prepend(right, s.from),
		},
		{
			x:    s.x - 1,
			y:    s.y,
			from: prepend(left, s.from),
		},
		{
			x:    s.x,
			y:    s.y + 1,
			from: prepend(up, s.from),
		},
		{
			x:    s.x,
			y:    s.y - 1,
			from: prepend(down, s.from),
		},
	}
	// clean all where 4 times same direction
	// clean where last direction is opposite of second last
	result = slices.DeleteFunc(result, func(seen Seen) bool {
		if len(seen.from) > 3 && same(seen.from) {
			return true
		}
		if len(seen.from) > 1 && seen.from[0].opposite() == seen.from[1] {
			return true
		}
		return false
	})
	// trim cutList
	for i := 0; i < len(result); i++ {
		result[i].from = cutList(result[i].from, 3)
	}
	return result
}

func prepend[T any](item T, list []T) []T {
	return append([]T{item}, list...)
}

func cutList[T any](list []T, l int) []T {
	if len(list) > l {
		return list[:l]
	} else {
		return list
	}
}

func same(list []direction) bool {
	if len(list) == 0 {
		panic("empty list")
	}
	x := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] != x {
			return false
		}
	}
	return true
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
	start1 := SeenUltra{1, 0, []direction{right}}
	start2 := SeenUltra{0, 1, []direction{down}}
	open := &MyHeap[SeenUltra]{HeapElement[SeenUltra]{city[start1.x][start1.y], start1},
		HeapElement[SeenUltra]{city[start2.x][start2.y], start2}}
	seen := map[string]bool{}
	seen[start1.String()] = true
	seen[start2.String()] = true
	heap.Init(open)
	for len(*open) > 0 {
		current := heap.Pop(open).(HeapElement[SeenUltra])
		if current.value.x == w-1 && current.value.y == h-1 {
			if atLeast4Straight(current.value.from) {
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
			s := reachable.String()
			if _, ok := seen[s]; ok {
				continue
			}
			seen[s] = true
			heap.Push(open, HeapElement[SeenUltra]{current.key + city[reachable.x][reachable.y], reachable})
		}
	}
	panic("didn't find a path at all, wtf")
}

func atLeast4Straight(from []direction) bool {
	if len(from) < 4 {
		return false
	}
	last := from[0]
	for i := 0; i < 4; i++ {
		if from[i] != last {
			return false
		}
	}
	return true
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
	x, y int
	from []direction
}

func (s SeenUltra) String() string {
	return fmt.Sprintf("%d,%d,%v", s.x, s.y, s.from)
}

func (s *SeenUltra) equals(o *Seen) bool {
	if s.x != o.x || s.y != o.y || len(s.from) != len(o.from) {
		return false
	}
	for i := 0; i < len(s.from); i++ {
		if s.from[i] != o.from[i] {
			return false
		}
	}
	return true
}

func (s SeenUltra) reachable() []SeenUltra {
	result := []SeenUltra{
		{
			x:    s.x + 1,
			y:    s.y,
			from: prepend(right, s.from),
		},
		{
			x:    s.x - 1,
			y:    s.y,
			from: prepend(left, s.from),
		},
		{
			x:    s.x,
			y:    s.y - 1,
			from: prepend(up, s.from),
		},
		{
			x:    s.x,
			y:    s.y + 1,
			from: prepend(down, s.from),
		},
	}

	result = slices.DeleteFunc(result, func(seen SeenUltra) bool {
		// clean all where 10 times same direction
		if len(seen.from) > 10 && same(seen.from) {
			return true
		}
		// clean where last direction is opposite of second last
		if len(seen.from) > 1 && seen.from[0].opposite() == seen.from[1] {
			return true
		}
		// clean where it turned before doing 4 in the same direction
		if seen.from[0] != seen.from[1] {
			if len(seen.from) < 5 {
				return true
			}
			last := seen.from[1]
			for i := 1; i <= 4; i++ {
				if seen.from[i] != last {
					return true
				}
			}
		}
		return false
	})
	// trim cutList
	for i := 0; i < len(result); i++ {
		result[i].from = cutList(result[i].from, 10)
	}
	return result
}
