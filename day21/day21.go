package day21

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	return solve1(input, 64)
}

func solve1(input string, wantedSteps int) string {
	CLEANER := nOnes(wantedSteps)
	garden := make(map[Pos]Tile)
	var start = Pos{0, 0}
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			if r == 'S' {
				r = '.'
				start = Pos{x, y}
			}
			garden[Pos{x, y}] = Tile{kind: r}
		}
	}
	var queue Queue
	for _, neighbor := range start.neighbors() {
		queue.enqueue(Message{neighbor, 1})
	}
	for message, ok := queue.dequeue(); ok; message, ok = queue.dequeue() {
		target, ok := garden[message.destination]
		if !ok {
			continue
		}
		if target.kind == '#' {
			continue
		}
		//fmt.Println(message)
		before := target.reachable
		after := before | message.increasedReachable
		after = after & CLEANER
		if after == before {
			continue
		}
		target.reachable = after
		garden[message.destination] = target
		increasedReachable := after << 1
		for _, neighbor := range message.destination.neighbors() {
			queue.enqueue(Message{neighbor, increasedReachable})
		}
	}
	count := 0
	for _, v := range garden {
		if contains(v.reachable, wantedSteps) {
			count++
		}
	}
	//printGarden(garden)
	return strconv.Itoa(count)
}

func printGarden(garden map[Pos]Tile) {
	for c := 1; c <= 6; c++ {
		fmt.Printf("Distance %d:\n", c)
		for y := 0; y < 11; y++ {
			for x := 0; x < 11; x++ {
				v := garden[Pos{x, y}]
				//fmt.Printf("%2d %2d ,%06b", x, y, v.reachable)
				if contains(v.reachable, c) {
					fmt.Print("O")
				} else {
					fmt.Print(string(v.kind))
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
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

type Tile struct {
	reachable uint64
	kind      rune
}

type Message struct {
	destination        Pos
	increasedReachable uint64
}

func (m Message) String() string {
	return fmt.Sprintf("%v - %08b", m.destination, m.increasedReachable)
}

func Part2(_ string) string {
	return ""
}

type Queue struct {
	messages []Message
}

func (q *Queue) enqueue(m Message) {
	q.messages = append(q.messages, m)
}

func (q *Queue) dequeue() (message Message, ok bool) {
	if len(q.messages) > 0 {
		m := q.messages[0]
		q.messages = q.messages[1:]
		return m, true
	} else {
		return Message{}, false
	}
}
