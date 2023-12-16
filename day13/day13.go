package day13

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	sum := 0
	for _, image := range strings.Split(input, "\n\n") {
		sum += points(image)
	}
	return strconv.Itoa(sum)
}

func points(image string) int {
	im := strings.Split(image, "\n")
	v := reflection(im)
	h := reflection(rotate(im))
	return 100*v + h
}

func rotate(im []string) (result []string) {
	result = make([]string, len(im[0]))
	for _, row := range im {
		for x, r := range row {
			y := len(im[0]) - 1 - x
			result[y] = result[y] + string(r)
		}
	}
	slices.Reverse(result)
	return
}

func reflection(split []string) int {
OUT:
	for i := 1; i < len(split); i++ {
		for h := 1; h < (len(split)+1)/2; h++ {
			if i-h < 0 || i+h-1 >= len(split) {
				break
			}
			u := split[i-h]
			l := split[i+h-1]
			if u != l {
				continue OUT
			}
		}
		return i
	}
	return 0
}

func Part2(_ string) string {
	return ""
}
