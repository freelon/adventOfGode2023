package day25

import (
	"adventOfGode2023/util"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func Part1(input string) string {
	links := make(map[string][]string)
	edgeIds := make(map[string]int)
	edges := 0
	for _, line := range strings.Split(input, "\n") {
		from, rhs, _ := strings.Cut(line, ": ")
		for _, to := range strings.Split(rhs, " ") {
			if _, ok := links[from]; !ok {
				links[from] = make([]string, 0)
			}
			links[from] = append(links[from], to)
			if _, ok := links[to]; !ok {
				links[to] = make([]string, 0)
			}
			links[to] = append(links[to], from)
			edgeIds[from+to] = edges
			edgeIds[to+from] = edges
			edges++
		}
	}

	var allNodes []string
	for k := range links {
		allNodes = append(allNodes, k)
	}

	edgeCounts := make([]Count, edges)
	for k, v := range edgeIds {
		edgeCounts[v].from = k[:3]
		edgeCounts[v].to = k[3:]
	}

	runs := min(len(allNodes), 100)
	for run := 0; run < runs; run++ {
		start := allNodes[run]
		predecessors := bfs(links, start)
		for _, target := range allNodes {
			if start == target {
				continue
			}
			current := target
			for predecessor, ok := predecessors[current]; ok; predecessor, ok = predecessors[current] {
				edgeCounts[edgeIds[current+predecessor]].count++
				current = predecessor
			}
		}
	}

	slices.SortFunc(edgeCounts, func(a, b Count) int {
		return cmp.Compare(b.count, a.count)
	})

	top3 := edgeCounts[:3]
	for _, edge := range top3 {
		x := links[edge.from]
		x = slices.DeleteFunc(x, func(s string) bool {
			return s == edge.to
		})
		links[edge.from] = x

		y := links[edge.to]
		y = slices.DeleteFunc(y, func(s string) bool {
			return s == edge.from
		})
		links[edge.to] = y

	}

	predecessors := bfs(links, allNodes[0])
	countA := len(predecessors) + 1 // one for the start node
	countB := len(allNodes) - countA
	return fmt.Sprintf("%d", countA*countB)
}

type Count struct {
	from, to string
	count    int
}

func bfs(links map[string][]string, start string) (predecessors map[string]string) {
	predecessors = make(map[string]string)
	predecessors[start] = ""
	var queue util.Queue[string]
	queue.Enqueue(start)
	for current, ok := queue.Dequeue(); ok; current, ok = queue.Dequeue() {
		for _, next := range links[current] {
			if _, ok := predecessors[next]; ok {
				continue
			}
			predecessors[next] = current
			queue.Enqueue(next)
		}
	}
	delete(predecessors, start)
	return
}
