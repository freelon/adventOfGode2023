package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day1Part1(input string) int {

	var numbers []int

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var digits []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		numberStr := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}

	var sum int
	for _, n := range numbers {
		sum = sum + n
	}

	return sum
}

func Day1Part2(input string) int {

	tokens := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var numbers []int

	for _, line := range strings.Split(input, "\n") {
		fmt.Println("line: ", line)
		if len(line) == 0 {
			continue
		}
		var digits []int

		l := line
	Outer:
		for i, _ := range l {
			sub := l[i:]
			for t, v := range tokens {
				if strings.HasPrefix(sub, t) {
					digits = append(digits, v)
					fmt.Println("..found ", t)
					continue Outer
				}
			}
		}

		numberStr := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
		fmt.Println(number)
	}

	var sum int
	for _, n := range numbers {
		sum = sum + n
	}

	// 55648 too low
	return sum
}
