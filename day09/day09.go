package day09

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	result := 0
	for _, line := range strings.Split(input, "\n") {
		historyValues := numbers(line)
		diffs := make([][]int, 0)
		diffs = append(diffs, historyValues)
		for {
			lastLine := diffs[len(diffs)-1]
			allZero := true
			for i := 0; i < len(lastLine); i++ {
				if lastLine[i] != 0 {
					allZero = false
					break
				}
			}
			if !allZero {
				newLine := make([]int, 0)
				for i := 0; i < len(lastLine)-1; i++ {
					newLine = append(newLine, lastLine[i+1]-lastLine[i])
				}
				diffs = append(diffs, newLine)
			} else {
				lastLine = append(lastLine, 0)
				for d := len(diffs) - 2; d >= 0; d-- {
					newValue := diffs[d][len(diffs[d])-1] + diffs[d+1][len(diffs[d+1])-1]
					diffs[d] = append(diffs[d], newValue)
				}
				result += diffs[0][len(diffs[0])-1]
				break
			}
		}
	}
	return strconv.Itoa(result)
}

func numbers(line string) []int {
	var result []int
	for _, n := range strings.Split(line, " ") {
		in, _ := strconv.Atoi(n)
		result = append(result, in)
	}
	return result
}

func Part2(input string) string {
	return ""
}
