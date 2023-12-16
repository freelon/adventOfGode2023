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
	cave := parse(input)
	beams := []Beam{{x: 0, y: 0, d: right}}
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
	return strconv.Itoa(energized)
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

func Part2(_ string) string {
	return ""
}
