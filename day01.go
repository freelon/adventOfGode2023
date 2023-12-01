package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day1(input string) any {

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
