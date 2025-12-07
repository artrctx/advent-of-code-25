package d7

import (
	"fmt"
	"slices"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

func Part1() {
	rows := input.GetRows("./internal/d7/input.txt")
	var answer uint

	rLen := len(rows[0])

	beams := rows[0]
	for idx := 2; idx < rLen; idx += 2 {
		for bIdx, b := range beams {
			if b != 'S' || rows[idx][bIdx] == '.' {
				continue
			}

			beams[bIdx] = '.'
			if bIdx > 0 {
				beams[bIdx-1] = 'S'
			}

			if bIdx < rLen-1 {
				beams[bIdx+1] = 'S'
			}

			answer++

		}
	}

	fmt.Printf("D7 Part 1: %v\n", answer)
}

func Part2() {
	rows := input.GetRows("./internal/d7/input.txt")

	rLen := len(rows[0])
	sIdx := slices.Index(rows[0], 'S')
	beams := slices.Replace(make([]uint, rLen), sIdx, sIdx+1, 1)
	fmt.Println(beams)
	for idx := 2; idx < rLen; idx += 2 {
		for bIdx, b := range beams {
			if b == 0 || rows[idx][bIdx] == '.' {
				continue
			}
			if idx == 3 {
				fmt.Print()
			}

			beams[bIdx] = 0
			beams[bIdx+1] += b
			beams[bIdx-1] += b
		}
	}

	var answer uint
	for _, v := range beams {
		answer += v
	}

	fmt.Printf("D7 Part 2: %v\n", answer)
}
