package day13

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	sum := 0
	for _, image := range strings.Split(input, "\n\n") {
		p, _, _ := points(image, -1, -1)
		sum += p
	}
	return strconv.Itoa(sum)
}

func points(image string, iv int, ih int) (int, int, int) {
	im := strings.Split(image, "\n")
	v := reflection(im, iv)
	h := reflection(rotate(im), ih)
	return 100*v + h, v, h
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

func reflection(split []string, ignore int) int {
OUT:
	for i := 1; i < len(split); i++ {
		if i == ignore {
			continue
		}
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

func Part2(input string) string {
	sum := 0
	for _, image := range strings.Split(input, "\n\n") {
		_, ov, oh := points(image, -1, -1)
		for i := 0; i < len(image); i++ {
			var cleanedImage string
			if image[i] == '.' {
				cleanedImage = image[:i] + "#" + image[i+1:]
			} else if image[i] == '#' {
				cleanedImage = image[:i] + "." + image[i+1:]
			} else {
				continue // don't bother with \n
			}
			_, v, h := points(cleanedImage, ov, oh)
			over := false
			if v > 0 && ov != v {
				sum += 100 * v
				over = true
			}
			if h > 0 && oh != h {
				sum += h
				over = true
			}
			if over {
				break
			}
		}
	}
	return strconv.Itoa(sum)
}
