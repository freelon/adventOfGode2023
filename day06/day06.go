package day06

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	times := numbers(lines[0])
	distances := numbers(lines[1])
	result := 1
	for race := 0; race < len(times); race++ {
		time := times[race]
		record := distances[race]
		winningConfigurations := 0
		for t := 0; t < time+1; t++ {
			speed := t
			distance := (time - t) * speed
			if distance > record {
				winningConfigurations++
			}
		}
		result *= winningConfigurations
	}
	return strconv.Itoa(result)
}

func numbers(s string) []int {
	r := regexp.MustCompile(`\d+`)
	allString := r.FindAllString(s, -1)
	var result []int
	for _, s := range allString {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	time := number(lines[0])
	record := number(lines[1])
	winningConfigurations := 0
	for t := 0; t < time+1; t++ {
		speed := t
		distance := (time - t) * speed
		if distance > record {
			winningConfigurations++
		}
	}
	return strconv.Itoa(winningConfigurations)
}

func number(s string) int {
	r := regexp.MustCompile(`\d+`)
	allString := r.FindAllString(s, -1)
	var result string
	for _, s := range allString {
		result += s
	}
	i, _ := strconv.Atoi(result)
	return i
}
