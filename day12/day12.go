package day12

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += arrangements(line)
	}
	return strconv.Itoa(sum)
}

func arrangements(line string) (result int) {
	maps.DeleteFunc(cache, func(s string, i int) bool {
		return true
	})
	parts := strings.Split(line, " ")
	baseLine := []rune(parts[0])
	rhs := strings.Split(parts[1], ",")
	var numbers []int
	neededDamaged := 0
	for _, number := range rhs {
		n, _ := strconv.Atoi(number)
		numbers = append(numbers, n)
		neededDamaged += n
	}
	var unknownIndices []int
	for i, r := range baseLine {
		if r == '?' {
			unknownIndices = append(unknownIndices, i)
		}
	}
	nDamaged := 0
	for _, r := range baseLine {
		if r == '#' {
			nDamaged++
		}
	}
	return foo(baseLine, numbers)
}

var cache = make(map[string]int)

func foo(line []rune, numbers []int) (result int) {
	cacheKey := fmt.Sprintf("%s%d", string(line), len(numbers))
	if v, ok := cache[cacheKey]; ok {
		return v
	}
	defer func() { cache[cacheKey] = result }()
	niceLine := string(line)
	_ = niceLine

	if len(numbers) == 0 {
		if slices.Contains(line, '#') {
			result = 0
			return
		} else {
			result = 1
			return
		}
	}

	next := numbers[0]

	if len(line) < next {
		// doesn't fit
		result = 0
		return
	} else if len(line) == next {
		if !slices.Contains(line, '.') {
			result = foo(line[len(line):], numbers[1:])
		} else {
			result = 0
		}
		return
	}

	// from here we always have at least next+1 runes left

	if line[0] == '#' {
		// either a match or abort
		if !slices.Contains(line[:next], '.') && line[next] != '#' {
			result = foo(line[next+1:], numbers[1:])
		} else {
			result = 0
		}
	} else if line[0] == '.' {
		// advance 1
		result = foo(line[1:], numbers)
	} else if line[0] == '?' {
		// assume it's a #
		line[0] = '#'
		result += foo(line, numbers)
		line[0] = '?'
		// assume it's a . and advance 1
		result += foo(line[1:], numbers)
	}
	return
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		line := line
		parts := strings.Split(line, " ")
		lhs := parts[0]
		rhs := parts[1]
		line = fmt.Sprintf("%s?%s?%s?%s?%s %s,%s,%s,%s,%s", lhs, lhs, lhs, lhs, lhs, rhs, rhs, rhs, rhs, rhs)
		sum += arrangements(line)
	}
	return strconv.Itoa(sum)
}
