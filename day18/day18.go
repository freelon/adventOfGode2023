package day18

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	points := []Pos{{0, 0}}
	walked := 0
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " ")
		d := parts[0]
		l, _ := strconv.Atoi(parts[1])
		last := points[len(points)-1]
		dx, dy := 0, 0
		switch d {
		case "U":
			dy = -l
		case "D":
			dy = l
		case "L":
			dx = -l
		case "R":
			dx = l
		}
		next := Pos{x: last.x + dx, y: last.y + dy}
		points = append(points, next)
		walked += l
	}

	// again picks formula to calculate number of blocks within the polygon (assuming the path doesn't cross itself)
	s := shoelace(points)
	inners := s - walked/2 + 1
	result := inners + walked
	return strconv.Itoa(result)
}

func Part2(input string) string {
	points := []Pos{{0, 0}}
	walked := 0
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " ")
		last := points[len(points)-1]
		d := parts[2][7]
		long, _ := strconv.ParseInt(parts[2][2:7], 16, 0)
		l := int(long)
		dx, dy := 0, 0
		switch d {
		case '3':
			dy = -l
		case '1':
			dy = l
		case '2':
			dx = -l
		case '0':
			dx = l
		}
		next := Pos{x: last.x + dx, y: last.y + dy}
		points = append(points, next)
		walked += l
	}

	// again picks formula to calculate number of blocks within the polygon (assuming the path doesn't cross itself)
	s := shoelace(points)
	inners := s - walked/2 + 1
	result := inners + walked
	return strconv.Itoa(result)

}

type Pos struct {
	x, y int
}

func shoelace(path []Pos) int {
	n := len(path)
	sumA := 0
	for i := 0; i < n; i++ {
		sumA += path[i].x * path[(i+1)%n].y
	}
	sumB := 0
	for i := 0; i < n; i++ {
		sumB += path[i].y * path[(i+1)%n].x
	}
	return (sumA - sumB) / 2
}
