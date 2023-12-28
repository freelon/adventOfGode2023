package day05

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	input = strings.TrimSpace(input)
	blocks := strings.Split(input, "\n\n")
	seeds := parseSeeds(blocks[0])
	var maps []elvesMap
	for i := 1; i < len(blocks); i++ {
		maps = append(maps, parseMap(blocks[i]))
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		v := seed
		for _, m := range maps {
			v = m.target(v)
		}
		if v < minLocation {
			minLocation = v
		}
	}

	return strconv.Itoa(minLocation)
}

type elvesMap struct {
	ranges []mapRange
}

type mapRange struct {
	sourceStart, destinationStart, length int
}

func (m elvesMap) target(source int) int {
	for _, r := range m.ranges {
		if source >= r.sourceStart && source < r.sourceStart+r.length {
			d := r.destinationStart - r.sourceStart
			return source + d
		}
	}
	return source
}

type interval struct {
	start, length int
}

func (m elvesMap) targetRange(source interval) (result []interval) {
	var sourceBorders = make(map[int]bool)
	for _, r := range m.ranges {
		sourceBorders[r.sourceStart] = true
		sourceBorders[r.sourceStart+r.length] = true
	}
	var borders = make([]int, 0)
	for sb := range sourceBorders {
		borders = append(borders, sb)
	}
	slices.Sort(borders)
	c := source
	for _, border := range borders {
		if c.start < border && border < c.start+c.length {
			l := border - c.start
			result = append(result, interval{c.start, l})
			c = interval{border, c.length - l}
		}
	}
	if c.length > 0 {
		result = append(result, c)
	}
	return
}

func parseMap(s string) elvesMap {
	lines := strings.Split(s, "\n")
	var ranges []mapRange
	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		targetStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		ranges = append(ranges, mapRange{
			sourceStart:      sourceStart,
			destinationStart: targetStart,
			length:           length,
		})
	}

	return elvesMap{ranges}
}

func parseSeeds(s string) (result []int) {
	rhs := strings.TrimPrefix(s, "seeds: ")
	for _, i := range strings.Split(rhs, " ") {
		v, _ := strconv.Atoi(i)
		result = append(result, v)
	}
	return
}

func Part2(input string) string {
	input = strings.TrimSpace(input)
	blocks := strings.Split(input, "\n\n")
	numbers := parseSeeds(blocks[0])
	var maps []elvesMap
	for i := 1; i < len(blocks); i++ {
		maps = append(maps, parseMap(blocks[i]))
	}

	intervals := make([]interval, 0)
	for i := 0; i < len(numbers); i += 2 {
		intervals = append(intervals, interval{numbers[i], numbers[i+1]})
	}

	for _, m := range maps {
		var next []interval
		for _, i := range intervals {
			for _, source := range m.targetRange(i) {
				target := interval{start: m.target(source.start), length: source.length}
				next = append(next, target)
			}
		}
		intervals = next
	}

	return strconv.Itoa(slices.MinFunc(intervals, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	}).start)
}
