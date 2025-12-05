package d5

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

func getIngredientIdsAndRange(rows [][]byte) (rng [][]byte, ids [][]byte) {
	for idx, row := range rows {
		if len(row) == 0 {
			rng, ids = rows[:idx], rows[idx+1:]
			break
		}
	}
	return
}

func prepIngridientRangeSet(rngs [][]byte) [][2]int {
	rngSet := make([][2]int, 0, len(rngs))
	slices.SortFunc(rngs, func(a []byte, b []byte) int {
		i1, i2 := input.Seperate(a, '-'), input.Seperate(b, '-')
		aNum, _ := strconv.Atoi(string(i1[0]))
		bNum, _ := strconv.Atoi(string(i2[0]))
		a2Num, _ := strconv.Atoi(string(i1[1]))
		b2Num, _ := strconv.Atoi(string(i2[1]))
		return cmp.Or(cmp.Compare(aNum, bNum), -cmp.Compare(a2Num, b2Num))
	})

	for _, v := range rngs {
		sep := input.Seperate(v, '-')
		start, _ := strconv.Atoi(string(sep[0]))
		end, _ := strconv.Atoi(string(sep[1]))

		tailOverlap := slices.IndexFunc(rngSet, func(r [2]int) bool {
			return r[1] >= start
		})

		if tailOverlap != -1 {
			if rngSet[tailOverlap][1] < end {
				rngSet[tailOverlap][1] = end
			}
			continue
		}

		rngSet = append(rngSet, [2]int{start, end})
	}

	return slices.Clip(rngSet)
}

func Part1() {
	rngs, ids := getIngredientIdsAndRange(input.GetRows("./internal/d5/input.txt"))
	rngSet := prepIngridientRangeSet(rngs)

	var answer int
	for _, id := range ids {
		idNum, _ := strconv.Atoi(string(id))

		low, high := 0, len(rngSet)-1
		for low <= high {
			// prevents overflow
			mid := low + (high-low)/2
			r := rngSet[mid]
			if idNum >= r[0] && idNum <= r[1] {
				answer++
				break
			}

			if r[1] < idNum {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	fmt.Printf("D5 Part 1: %v\n", answer)
}

// 350556583134447 too low
func Part2() {
	rngs, ids := getIngredientIdsAndRange(input.GetRows("./internal/d5/input.txt"))
	rngSet := prepIngridientRangeSet(rngs)
	var answer int
	loggedSet := make(map[[2]int]struct{}, len(rngSet))
	for _, id := range ids {
		idNum, _ := strconv.Atoi(string(id))
		low, high := 0, len(rngSet)-1
		for low <= high {
			// prevents overflow
			mid := low + (high-low)/2
			r := rngSet[mid]
			if idNum >= r[0] && idNum <= r[1] {
				if _, exists := loggedSet[r]; !exists {
					answer += (r[1] - r[0] + 1)
					loggedSet[r] = struct{}{}
				}
				break
			}

			if r[1] < idNum {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	fmt.Printf("D5 Part 1: %v\n", answer)
}
