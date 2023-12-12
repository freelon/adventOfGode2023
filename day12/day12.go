package day12

import (
	"fmt"
	"regexp"
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
	parts := strings.Split(line, " ")
	baseLine := []rune(parts[0])
	rhs := strings.Split(parts[1], ",")
	var numbers []int
	for _, number := range rhs {
		n, _ := strconv.Atoi(number)
		numbers = append(numbers, n)
	}
	var unknownIndices []int
	for i, r := range baseLine {
		if r == '?' {
			unknownIndices = append(unknownIndices, i)
		}
	}
	return foo(baseLine, unknownIndices, numbers)
}

func foo(line []rune, unknownIndices []int, numbers []int) int {
	if impossible(line, numbers) {
		return 0
	}

	x := string(line)
	x = x
	if len(unknownIndices) == 0 {
		if isValid(line, numbers) {
			//println(string(line))
			return 1
		} else {
			return 0
		}
	}
	result := 0
	unknownIndex := unknownIndices[0]
	line[unknownIndex] = '.'
	x = string(line)
	if hasValidPrefix(line, numbers) {
		result += foo(line, unknownIndices[1:], numbers)
	}
	line[unknownIndex] = '#'
	x = string(line)
	if hasValidPrefix(line, numbers) {
		result += foo(line, unknownIndices[1:], numbers)
	}
	line[unknownIndex] = '?'
	x = string(line)
	return result
}

func impossible(line []rune, numbers []int) bool {
	springs := 0
	unknown := 0

	for _, r := range line {
		if r == '#' {
			springs++
		}
		if r == '?' {
			unknown++
		}
	}

	needed := 0
	for _, number := range numbers {
		needed += number
	}

	if springs > needed {
		return true
	}
	if springs+unknown < needed {
		return true
	}
	return false
}

func hasValidPrefix(line []rune, numbers []int) bool {
	r := regexp.MustCompile(`#+`)
	stringLine := string(line)
	parts := strings.Split(stringLine, "?")

	// if the string is not fully determined do not consider #s right before the next ?
	s := parts[0]
	if len(parts) > 1 {
		for strings.HasSuffix(s, "#") {
			s = strings.TrimSuffix(s, "#")
		}
	}

	allString := r.FindAllString(s, -1)
	if len(allString) > len(numbers) {
		return false
	}
	for i := 0; i < len(allString); i++ {
		if len(allString[i]) != numbers[i] {
			return false
		}
	}
	return true
}

func isValid(line []rune, numbers []int) bool {
	r := regexp.MustCompile(`#+`)
	stringLine := string(line)

	allString := r.FindAllString(stringLine, -1)
	if len(allString) != len(numbers) {
		return false
	}
	for i := 0; i < len(allString); i++ {
		if len(allString[i]) != numbers[i] {
			return false
		}
	}
	return true
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		lhs := parts[0]
		rhs := parts[1]
		line = fmt.Sprintf("%s?%s?%s?%s?%s %s,%s,%s,%s,%s", lhs, lhs, lhs, lhs, lhs, rhs, rhs, rhs, rhs, rhs)
		println(line)
		sum += arrangements(line)
	}
	return strconv.Itoa(sum)
}
