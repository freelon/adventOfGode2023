package day05

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
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

	nPairs := len(numbers) / 2
	minLocations := make([]int, nPairs)

	wg := sync.WaitGroup{}
	wg.Add(nPairs)

	for k := 0; k < nPairs; k++ {
		go func(k int) {
			defer wg.Done()
			i := 2 * k
			minLocation := math.MaxInt
			start := numbers[i]
			l := numbers[i+1]
			for j := start; j < start+l; j++ {
				v := j
				for _, m := range maps {
					v = m.target(v)
				}
				if v < minLocation {
					minLocation = v
				}
			}
			minLocations[k] = minLocation
		}(k)
	}

	wg.Wait()

	return strconv.Itoa(slices.Min(minLocations))
}
