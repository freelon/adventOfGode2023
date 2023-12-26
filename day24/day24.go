package day24

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input string) string {
	hail := parse(input)
	return strconv.Itoa(solve(hail, 200000000000000, 400000000000000))
}

func solve(hail []Hail, lBound int, uBound int) (count int) {
	c := 0
	for i, a := range hail {
		for j := i + 1; j < len(hail); j++ {
			b := hail[j]
			c++
			x, y, ok := intersect(a, b)
			if ok {
				if x < float64(lBound) || x > float64(uBound) || y < float64(lBound) || y > float64(uBound) {
					continue
				}
				count++
			}
		}
	}
	println(c)
	return
}

func intersect(a Hail, b Hail) (x float64, y float64, ok bool) {
	x1 := float64(a.pos.x)
	y1 := float64(a.pos.y)
	x2 := x1 + float64(a.vel.x)
	y2 := y1 + float64(a.vel.y)
	x3 := float64(b.pos.x)
	y3 := float64(b.pos.y)
	x4 := x3 + float64(b.vel.x)
	y4 := y3 + float64(b.vel.y)
	ta := (x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)
	ta /= (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	x, y = x1+ta*(x2-x1), y1+ta*(y2-y1)
	if math.IsInf(x, 0) || math.IsInf(y, 0) {
		ok = false
		return
	}

	tb := (x - x3) / float64(b.vel.x)

	if ta >= 0 && tb >= 0 {
		ok = true
	}
	return
}

func parse(input string) []Hail {
	var hail []Hail
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " @ ")
		lhs := strings.Split(parts[0], ", ")
		x, _ := strconv.Atoi(lhs[0])
		y, _ := strconv.Atoi(lhs[1])
		z, _ := strconv.Atoi(lhs[2])
		pos := V{x, y, z}
		rhs := strings.Split(parts[1], ", ")
		x, _ = strconv.Atoi(rhs[0])
		y, _ = strconv.Atoi(rhs[1])
		z, _ = strconv.Atoi(rhs[2])
		vel := V{x, y, z}
		hail = append(hail, Hail{pos, vel})
	}
	return hail
}

type Hail struct {
	pos, vel V
}

func (h Hail) String() string {
	return fmt.Sprintf("%s @ %s", h.pos, h.vel)
}

type V struct {
	x, y, z int
}

func (v V) String() string {
	return fmt.Sprintf("%d, %d, %d", v.x, v.y, v.z)
}

func Part2(input string) string {
	hail := parse(input)
	equationY(hail[0], hail[1])
	equationY(hail[0], hail[2])
	equationY(hail[0], hail[3])
	equationY(hail[0], hail[4])
	println()
	println()
	equationZ(hail[0], hail[1])
	equationZ(hail[0], hail[2])
	equationZ(hail[0], hail[3])
	equationZ(hail[0], hail[4])
	return ""
}

func equationY(a Hail, b Hail) {
	fmt.Printf("%3dX + %3dY + %3dA + %3dB = %3d\n",
		b.vel.y-a.vel.y,
		a.vel.x-b.vel.x,
		b.pos.y-a.pos.y,
		b.pos.x-a.pos.x,
		b.pos.x*b.vel.y-b.pos.y*b.vel.x-a.pos.x*a.vel.y+a.pos.y*a.vel.x)
}

func equationZ(a Hail, b Hail) {
	fmt.Printf("%3dX + %3dZ + %3dA + %3dB = %3d\n",
		b.vel.z-a.vel.z,
		a.vel.x-b.vel.x,
		b.pos.z-a.pos.z,
		b.pos.x-a.pos.x,
		b.pos.x*b.vel.z-b.pos.z*b.vel.x-a.pos.x*a.vel.z+a.pos.z*a.vel.x)
}
