package day08

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	movements, desertMap := parse(input)
	count := 0
	node := "AAA"
	for {
		if node == "ZZZ" {
			break
		}
		m := movements[count%len(movements)]
		count++
		if m == 'L' {
			node = desertMap[node].l
		} else if m == 'R' {
			node = desertMap[node].r
		}
	}
	return strconv.Itoa(count)
}

func parse(input string) (movements []rune, desertMap map[string]Pair) {
	regex := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
	parts := strings.Split(input, "\n\n")
	for _, r := range parts[0] {
		movements = append(movements, r)
	}
	desertMap = make(map[string]Pair)
	for _, line := range strings.Split(parts[1], "\n") {
		matches := regex.FindStringSubmatch(line)
		p := Pair{
			l: matches[2],
			r: matches[3],
		}
		desertMap[matches[1]] = p
	}
	return
}

type Pair struct {
	l string
	r string
}

func Part2(input string) string {
	movements, desertMap := parse(input)
	count := 0
	nodes := make([]string, 0)
	for key := range desertMap {
		if key[2] == 'A' {
			nodes = append(nodes, key)
		}
	}
	firstAtTarget := make(map[int]int)
	cycleTime := make(map[int]int)
	for {
		for n, node := range nodes {
			if node[2] == 'Z' {
				firstForNode, ok := firstAtTarget[n]
				if !ok {
					firstAtTarget[n] = count
				} else {
					// the node reaches his goal the 2nd time
					_, ok2 := cycleTime[n]
					if !ok2 {
						cycleTime[n] = count - firstForNode
					}
				}
			}
		}
		if len(cycleTime) == len(nodes) {
			break
		}
		m := movements[count%len(movements)]
		count++
		for i := 0; i < len(nodes); i++ {
			if m == 'L' {
				nodes[i] = desertMap[nodes[i]].l
			} else if m == 'R' {
				nodes[i] = desertMap[nodes[i]].r
			}
		}
	}
	fmt.Printf("After {} rounds all have cycled.\n")
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("Node %d: first @ %d, cycle %d\n", i, firstAtTarget[i], cycleTime[i])
	}
	var rest []int
	for i := 2; i < len(cycleTime); i++ {
		rest = append(rest, cycleTime[i])
	}
	lcm := LCM(cycleTime[0], cycleTime[1], rest)
	return strconv.Itoa(lcm)
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find the Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], make([]int, 0))
	}

	return result
}
