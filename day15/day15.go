package day15

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	parts := strings.Split(input, ",")
	sum := 0
	for _, part := range parts {
		sum += hash(part)
	}
	return strconv.Itoa(sum)
}

func hash(part string) (result int) {
	for _, r := range part {
		result += int(r)
		result *= 17
		result %= 256
	}
	return
}

func Part2(input string) string {
	boxes := installation(input)
	focusPower := 0
	for i, box := range boxes {
		for j, lens := range box {
			lensFocusPower := (i + 1) * (j + 1) * lens.focalLength
			focusPower += lensFocusPower
		}
	}
	return strconv.Itoa(focusPower)
}

func installation(input string) (boxes [256][]LabeledLens) {
	for _, part := range strings.Split(input, ",") {
		if i := strings.Index(part, "="); i > -1 {
			label := part[0:i]
			focalLength, _ := strconv.Atoi(part[i+1:])
			labelHash := hash(label)
			if j := indexOfLabel(boxes[labelHash], label); j > -1 {
				boxes[labelHash][j].focalLength = focalLength
			} else {
				boxes[labelHash] = append(boxes[labelHash], LabeledLens{
					label:       label,
					focalLength: focalLength,
				})
			}
		} else {
			label := part[:len(part)-1]
			labelHash := hash(label)
			boxes[labelHash] = slices.DeleteFunc(boxes[labelHash], func(lens LabeledLens) bool {
				return lens.label == label
			})
		}
	}
	return
}

func indexOfLabel(lenses []LabeledLens, label string) int {
	for i, labeledLens := range lenses {
		if labeledLens.label == label {
			return i
		}
	}
	return -1
}

type LabeledLens struct {
	label       string
	focalLength int
}
