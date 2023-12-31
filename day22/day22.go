package day22

import (
	"adventOfGode2023/util"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	bricks := parse(input)
	bricks = applyGravity(bricks)
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
		return a.from.z - b.from.z
	})
FALLING:
	for i := 0; i < len(bricks); i++ {
		current := bricks[i]
		lowestZ := min(current.to.z, current.from.z)
		if lowestZ <= 1 {
			continue
		}
		target := current
		target.from.z--
		target.to.z--
		for j := i - 1; j >= 0; j-- {
			if bricks[j].to.z != target.from.z {
				continue
			}
			if bricks[j].overlaps(target) {
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
			for bricks[i].from.z < bricks[i-1].from.z {
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
	return c.x >= b.from.x &&
		c.x <= b.to.x &&
		c.y >= b.from.y &&
		c.y <= b.to.y &&
		c.z >= b.from.z &&
		c.z <= b.to.z
}

func (b Brick) containsAny(cs []C) bool {
	for _, c := range cs {
		if b.contains(c) {
			return true
		}
	}
	return false
}

func (b Brick) overlaps(other Brick) bool {
	overlapZ := intervalOverlaps(b.from.z, b.to.z, other.from.z, other.to.z)
	overlapY := intervalOverlaps(b.from.y, b.to.y, other.from.y, other.to.y)
	overlapX := intervalOverlaps(b.from.x, b.to.x, other.from.x, other.to.x)

	return overlapX && overlapY && overlapZ
}

// intervalOverlaps interval borders are inclusive
func intervalOverlaps(aStart int, aEnd int, bStart int, bEnd int) bool {
	if aStart <= bStart && bStart <= aEnd {
		return true
	}
	if aStart <= bEnd && bEnd <= aEnd {
		return true
	}
	if bStart <= aStart && aStart <= bEnd {
		return true
	}
	if bStart <= aEnd && aEnd <= bEnd {
		return true
	}
	return false
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
