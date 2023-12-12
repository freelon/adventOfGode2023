package day12

import (
	"fmt"
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
	offset := 0
	return foo(baseLine, unknownIndices, numbers, nDamaged, neededDamaged, offset)
}

func foo(line []rune, unknownIndices []int, numbers []int, nDamaged int, neededDamaged int, offset int) int {
	//niceLine := string(line)
	//_ = niceLine
	if len(unknownIndices)+nDamaged < neededDamaged {
		return 0
	}
	if nDamaged > neededDamaged {
		return 0
	}
	if offset == len(line) {
		return 1
	}

	result := 0
	unknownIndex := unknownIndices[0]
	line[unknownIndex] = '.'
	if valid, validLength, numbersConsumed := prefix(line, numbers, offset); valid {
		result += foo(line, unknownIndices[1:], numbers[numbersConsumed:], nDamaged, neededDamaged, offset+validLength)
	}
	line[unknownIndex] = '#'
	if valid, validLength, numbersConsumed := prefix(line, numbers, offset); valid {
		result += foo(line, unknownIndices[1:], numbers[numbersConsumed:], nDamaged+1, neededDamaged, offset+validLength)
	}
	line[unknownIndex] = '?'
	return result
}

func prefix(line []rune, numbers []int, offset int) (valid bool, plusOffset int, numbersConsumed int) {
	foundNumbers := make([]int, 0)
	last := '.'
	x := 0
	currentExpectedNumberIndex := -1
	for i := offset; i < len(line); i++ {
		l := last
		c := line[i]
		last = c
		if l == '.' && c == '.' {
			continue
		} else if l == '.' && c == '#' {
			x = 1
			currentExpectedNumberIndex++
			if currentExpectedNumberIndex >= len(numbers) {
				return false, 0, 0
			}
			currentExpectedNumber := numbers[currentExpectedNumberIndex]
			for k := 0; k < currentExpectedNumber; k++ {
				expectDamageIndex := i + k
				if expectDamageIndex >= len(line) || line[expectDamageIndex] == '.' {
					return false, 0, 0
				}
			}
		} else if l == '#' && c == '#' {
			x++
		} else if l == '#' && c == '.' {
			foundNumbers = append(foundNumbers, x)
			plusOffset = i - offset
			x = 0
		} else if c == '?' {
			break
		}
	}
	if last == '#' { // stopped because end of string (no '?') and last not yet added
		foundNumbers = append(foundNumbers, x)
		plusOffset = len(line) - offset
	} else if last == '.' {
		plusOffset = len(line) - offset
	}
	if len(foundNumbers) > len(numbers) {
		return false, 0, 0
	}
	for j := 0; j < len(foundNumbers); j++ {
		if foundNumbers[j] != numbers[j] {
			return false, 0, 0
		}
		numbersConsumed++
	}
	valid = true
	return
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for k, line := range lines {
		parts := strings.Split(line, " ")
		lhs := parts[0]
		rhs := parts[1]
		line = fmt.Sprintf("%s?%s?%s?%s?%s %s,%s,%s,%s,%s", lhs, lhs, lhs, lhs, lhs, rhs, rhs, rhs, rhs, rhs)
		fmt.Printf("% 4d %s\n", k, line)
		sum += arrangements(line)
	}
	return strconv.Itoa(sum)
}
