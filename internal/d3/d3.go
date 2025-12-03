package d3

import (
	"fmt"
	"math"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

// 0 in utf8
var ZERO byte = 48

func Part1() {
	banks := input.GetRows("./internal/d3/input.txt")
	var answer uint
	bankLen := len(banks[0])
	for _, bank := range banks {
		// primary, secondary
		var p byte
		var s byte
		for idx, bat := range bank {
			if bat > p && idx != bankLen-1 {
				p = bat
				s = 0
				continue
			}

			if bat > s {
				s = bat
				if bat == ZERO+9 {
					break
				}
			}
		}

		answer += uint((p-ZERO)*10 + s - ZERO)
	}

	fmt.Printf("D3 Part 1: %v\n", answer)
}

// 169129541760449 too high
func Part2() {
	banks := input.GetRows("./internal/d3/input.txt")
	var answer uint
	bankLen := len(banks[0])
	for _, bank := range banks {
		bats := make([]*byte, 12)
	bankLoop:
		for i, bat := range bank {
			offset := 0
			if i > bankLen-12 {
				offset = i + 12 - bankLen
			}

			logged := false
			for idx, b := range bats[offset:] {
				bOffset := idx + offset
				if logged {
					bats[idx+offset] = nil
					continue
				}
				if b != nil && *b >= bat {
					continue
				}
				bats[idx+offset] = &bat
				if bOffset == 11 && bat == ZERO+9 {
					break bankLoop
				}
				logged = true
			}
		}

		// aggregate
		for idx, b := range bats {
			answer += uint((*b - ZERO)) * uint(math.Pow10(11-idx))
		}
	}

	fmt.Printf("D3 Part 1: %v\n", answer)
}
