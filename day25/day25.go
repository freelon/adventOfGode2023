package day25

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	links := make(map[string][]string)
	allNames := make(map[string]bool)
	for _, line := range strings.Split(input, "\n") {
		from, rhs, _ := strings.Cut(line, ": ")
		allNames[from] = true
		for _, to := range strings.Split(rhs, " ") {
			allNames[to] = true
			if _, ok := links[from]; !ok {
				links[from] = make([]string, 0)
			}
			links[from] = append(links[from], to)
		}
	}

	n := len(allNames)
	am := make([][]int, n)
	for i := 0; i < n; i++ {
		am[i] = make([]int, n)
	}
	ids := make(map[string]int)
	id := 0
	for k, _ := range allNames {
		ids[k] = id
		id++
	}
	for from, tos := range links {
		fromId := ids[from]
		for _, to := range tos {
			toId := ids[to]
			am[fromId][toId] = 1
			am[toId][fromId] = 1
		}
	}

	//Karger's algorithm (https://en.wikipedia.org/wiki/Karger%27s_algorithm)
	for attempt := 1; true; attempt++ {

		a, b, ok := threeCut(copyMatrix(am))
		if ok {
			fmt.Printf("Needed attempts: %d\n", attempt)
			return strconv.Itoa(a * b)
		}
	}
	panic("unreachable")
}

func threeCut(am [][]int) (sizeA int, sizeB int, match bool) {
	n := len(am)
	var nodes []int
	var representing = make([]int, n)
	for i := 0; i < n; i++ {
		nodes = append(nodes, i)
		representing[i] = 1
	}
	for len(nodes) > 2 {
		// pick a random edge
		nodeA := nodes[rand.Intn(len(nodes))]
		var edges []int
		for i := 0; i < n; i++ {
			if am[nodeA][i] > 0 {
				edges = append(edges, i)
			}
		}
		nodeB := edges[rand.Intn(len(edges))]

		// merge
		representing[nodeA] += representing[nodeB]
		representing[nodeB] = 0
		nodes = slices.DeleteFunc(nodes, func(i int) bool {
			return i == nodeB
		})

		for i := 0; i < n; i++ {
			am[nodeA][i] += am[nodeB][i]
			am[i][nodeA] += am[i][nodeB]
			am[nodeB][i] = 0
			am[i][nodeB] = 0
		}
		am[nodeA][nodeA] = 0
	}
	a := nodes[0]
	b := nodes[1]
	if am[a][b] == 3 {
		return representing[a], representing[b], true
	}
	return 0, 0, false
}

func copyMatrix(m [][]int) (r [][]int) {
	r = make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		r[i] = slices.Clone(m[i])
	}
	return
}
