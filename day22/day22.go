package day22

import (
	"adventOfGode2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Part1(input string) string {
	bricks := parse(input)
	start := time.Now()
	bricks = applyGravity(bricks)
	fmt.Println(time.Now().Sub(start))
	supports := supporting(bricks)
	supporterCount := make([]int, len(supports))
	for _, supported := range supports {
		for _, supportee := range supported {
			supporterCount[supportee]++
		}
	}
	disintegratable := 0
SUPPORTER:
	for s := 0; s < len(supports); s++ {
		supporting := supports[s]
		if len(supporting) == 0 {
			disintegratable++
			continue
		}
		for _, supportee := range supporting {
			if supporterCount[supportee] < 2 {
				continue SUPPORTER
			}
		}
		disintegratable++
	}
	return strconv.Itoa(disintegratable)
}

func supporting(bricks []Brick) map[int][]int {
	var supports = map[int][]int{} // list of bricks that one brick supports
	for i := 0; i < len(bricks); i++ {
		supports[i] = make([]int, 0)

		current := bricks[i]
		highestZ := max(current.to.z, current.from.z)
		var directlyAbove []C
		for x := min(current.from.x, current.to.x); x <= max(current.from.x, current.to.x); x++ {
			for y := min(current.from.y, current.to.y); y <= max(current.from.y, current.to.y); y++ {
				c := C{x, y, highestZ + 1}
				directlyAbove = append(directlyAbove, c)
			}
		}
		for j := 0; j < len(bricks); j++ {
			if bricks[j].containsAny(directlyAbove) {
				supports[i] = append(supports[i], j)
			}
		}
	}
	return supports
}

func applyGravity(bricks []Brick) []Brick {
	slices.SortFunc(bricks, func(a, b Brick) int {
		az := min(a.from.z, a.to.z)
		bz := min(b.from.z, b.to.z)
		return az - bz
	})
FALLING:
	for i := 0; i < len(bricks); i++ {
		current := bricks[i]
		lowestZ := min(current.to.z, current.from.z)
		if lowestZ <= 1 {
			continue
		}
		var hopefullyEmpty []C
		targetZ := lowestZ - 1
		for x := min(current.from.x, current.to.x); x <= max(current.from.x, current.to.x); x++ {
			for y := min(current.from.y, current.to.y); y <= max(current.from.y, current.to.y); y++ {
				c := C{x, y, targetZ}
				hopefullyEmpty = append(hopefullyEmpty, c)
			}
		}
		for j := i - 1; j >= 0; j-- {
			if bricks[j].highest() != targetZ {
				continue
			}
			if bricks[j].containsAny(hopefullyEmpty) {
				continue FALLING
			}
		}
		// no collision, let's change the world for a ~~better~~ sandier place
		current.to.z -= 1
		current.from.z -= 1
		bricks[i] = current
		// keep pointer at current brick to check if it can fall further
		i--
		// maintain sort by lowest
		if i > 0 {
			for bricks[i].lowest() < bricks[i-1].lowest() {
				bricks[i], bricks[i-1] = bricks[i-1], bricks[i]
				// keep pointer on current brick, again
				i--
			}
		}
	}

	return bricks
}

type Brick struct {
	from C
	to   C
}

func (b Brick) contains(c C) bool {
	return c.x >= min(b.from.x, b.to.x) &&
		c.x <= max(b.from.x, b.to.x) &&
		c.y >= min(b.from.y, b.to.y) &&
		c.y <= max(b.from.y, b.to.y) &&
		c.z >= min(b.from.z, b.to.z) &&
		c.z <= max(b.from.z, b.to.z)
}

func (b Brick) containsAny(cs []C) bool {
	for _, c := range cs {
		if b.contains(c) {
			return true
		}
	}
	return false
}

func (b Brick) highest() int {
	return max(b.from.z, b.to.z)
}

func (b Brick) lowest() int {
	return min(b.from.z, b.to.z)
}

type C struct {
	x, y, z int
}

func parseC(input string) C {
	split := strings.Split(input, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])
	return C{x, y, z}
}

func parse(input string) (result []Brick) {
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "~")
		result = append(result, Brick{from: parseC(split[0]), to: parseC(split[1])})
	}
	return
}

func Part2(input string) string {
	bricks := parse(input)
	bricks = applyGravity(bricks)
	supports := supporting(bricks)
	supporterCount := make([]int, len(supports))
	for _, supported := range supports {
		for _, supportee := range supported {
			supporterCount[supportee]++
		}
	}
	fallingSum := 0
	for supporter := 0; supporter < len(supports); supporter++ {
		hits := make([]int, len(supports))
		queue := util.Queue[int]{}
		falling := 0
		for _, supportee := range supports[supporter] {
			queue.Enqueue(supportee)
		}
		for supportee, ok := queue.Dequeue(); ok; supportee, ok = queue.Dequeue() {
			hits[supportee]++
			if hits[supportee] == supporterCount[supportee] {
				falling++
				for _, next := range supports[supportee] {
					queue.Enqueue(next)
				}
			}
		}
		fallingSum += falling
	}
	return strconv.Itoa(fallingSum)
}
