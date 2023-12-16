package day16

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type direction int

func (d direction) String() string {
	return [...]string{"up",
		"down",
		"left",
		"right"}[d]
}

const (
	up direction = iota
	down
	left
	right
)

type Beam struct {
	x, y int
	d    direction
}

func (b *Beam) String() string {
	return fmt.Sprintf("(%d, %d -> %s)", b.x, b.y, b.d)
}

func (b *Beam) cp() Beam {
	return Beam{b.x, b.y, b.d}
}

func (b *Beam) down() {
	b.d = down
	b.y++
}

func (b *Beam) up() {
	b.d = up
	b.y--
}

func (b *Beam) left() {
	b.d = left
	b.x--
}

func (b *Beam) right() {
	b.d = right
	b.x++
}

func Part1(input string) string {
	energized := sendBeam(input, Beam{
		x: 0,
		y: 0,
		d: right,
	})
	return strconv.Itoa(energized)
}

func sendBeam(input string, start Beam) int {
	beams := []Beam{start}
	cave := parse(input)
	for len(beams) > 0 {
		for _, beam := range beams {
			cave[beam.y][beam.x].energizedCount++
			cave[beam.y][beam.x].seen[beam.d] = true
		}
		var newBeams []Beam
		for b := 0; b < len(beams); b++ {
			beam := beams[b]
			kind := cave[beam.y][beam.x].kind
			switch kind {
			case '.':
				{
					switch beam.d {
					case down:
						beam.down()
					case up:
						beam.up()
					case left:
						beam.left()
					case right:
						beam.right()
					}
				}
			case '/':
				{
					switch beam.d {
					case down:
						beam.left()
					case up:
						beam.right()
					case left:
						beam.down()
					case right:
						beam.up()
					}
				}
			case '\\':
				{
					switch beam.d {
					case down:
						beam.right()
					case up:
						beam.left()
					case left:
						beam.up()
					case right:
						beam.down()
					}
				}
			case '-':
				{
					switch beam.d {
					case down:
						cp := beam.cp()
						beam.left()
						cp.right()
						newBeams = append(newBeams, cp)
					case up:
						cp := beam.cp()
						beam.left()
						cp.right()
						newBeams = append(newBeams, cp)
					case left:
						beam.left()
					case right:
						beam.right()
					}
				}
			case '|':
				{
					switch beam.d {
					case down:
						beam.down()
					case up:
						beam.up()
					case left:
						cp := beam.cp()
						beam.down()
						cp.up()
						newBeams = append(newBeams, cp)
					case right:
						cp := beam.cp()
						beam.down()
						cp.up()
						newBeams = append(newBeams, cp)
					}
				}
			default:
				panic("unknown tile kind " + string(kind))
			}
			beams[b] = beam
		}
		beams = append(beams, newBeams...)
		beams = slices.DeleteFunc(beams, func(beam Beam) bool {
			if beam.x < 0 || beam.y < 0 {
				return true
			}
			if beam.y >= len(cave) || beam.x >= len(cave[beam.y]) {
				return true
			}
			if _, ok := cave[beam.y][beam.x].seen[beam.d]; ok {
				return true
			}
			return false
		})
	}
	energized := 0
	for _, row := range cave {
		for _, tile := range row {
			if tile.energizedCount > 0 {
				energized++
			}
		}
	}
	return energized
}

type Tile struct {
	kind           rune
	energizedCount int
	seen           map[direction]bool
}

func parse(input string) (cave [][]Tile) {
	for _, row := range strings.Split(input, "\n") {
		c := make([]Tile, 0)
		for _, kind := range row {
			c = append(c, Tile{kind, 0, make(map[direction]bool)})
		}
		cave = append(cave, c)
	}
	return
}

func Part2(input string) string {
	height := len(strings.Split(input, "\n"))
	width := len(strings.Split(input, "\n")[0])
	starts := make([]Beam, 0)
	for x := 0; x < width; x++ {
		starts = append(starts, Beam{x: x, y: 0, d: down})
		starts = append(starts, Beam{x: x, y: height - 1, d: up})
	}
	for y := 0; y < width; y++ {
		starts = append(starts, Beam{x: 0, y: y, d: right})
		starts = append(starts, Beam{x: width - 1, y: y, d: left})
	}
	mx := 0
	for _, start := range starts {
		e := sendBeam(input, start)
		if e > mx {
			mx = e
		}
	}
	return strconv.Itoa(mx)
}
