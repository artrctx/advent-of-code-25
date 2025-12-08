package d8

import (
	"fmt"
	"math"
	"strconv"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

type point struct {
	x int
	y int
	z int
}

func (p *point) dis(np point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-np.x), 2) + math.Pow(float64(p.y-np.y), 2) + math.Pow(float64(p.z-np.z), 2))
}

func pointFromBytes(bs []byte) point {
	pnts := input.Seperate(bs, ',')
	x, _ := strconv.Atoi(string(pnts[0]))
	y, _ := strconv.Atoi(string(pnts[1]))
	z, _ := strconv.Atoi(string(pnts[2]))
	return point{x, y, z}
}

func Part1() {
	var answer int
	_ = input.GetRows("./internal/d8/example.txt")

	fmt.Printf("D8 Part 1: %v\n", answer)
}

func Part2() {
	var answer int

	fmt.Printf("D8 Part 2: %v\n", answer)
}
