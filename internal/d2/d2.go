package d2

import (
	"fmt"
	"strconv"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

var (
	comma byte = 44
	dash  byte = 45
)

func getIdsFromRange(ids []byte) (id1 int, id2 int) {
	sepIds := input.Seperate(ids, dash)
	id1, _ = strconv.Atoi(string(sepIds[0]))
	id2, _ = strconv.Atoi(string(sepIds[1]))
	return
}

func Part1() {
	rngs := input.Seperate(input.GetFile("./internal/d2/input.txt"), comma)

	var answer int
	for _, rng := range rngs {
		// dumb solution iterate each
		id1, id2 := getIdsFromRange(rng)
		for i := id1; i <= id2; i++ {
			iStr := strconv.Itoa(i)
			l := len(iStr)
			if l%2 == 1 || iStr[:l/2] != iStr[l/2:] {
				continue
			}

			answer += i
		}
	}

	fmt.Printf("D2 Part 1: %v\n", answer)
}

func Part2() {
	rngs := input.Seperate(input.GetFile("./internal/d2/input.txt"), comma)

	var answer int
	for _, rng := range rngs {
		// dumb solution iterate each
		id1, id2 := getIdsFromRange(rng)
		for i := id1; i <= id2; i++ {
			iStr := strconv.Itoa(i)
			l := len(iStr)

		outer:
			for step := range l / 2 {
				s := step + 1
				if l%s != 0 {
					continue
				}
				valid := true
				prev := iStr[:s]
			inner:
				for j := 1; j < l/s; j++ {
					if prev != iStr[j*s:j*s+s] {
						valid = false
						break inner
					}
				}
				if valid {
					answer += i
					break outer
				}
			}
		}
	}

	fmt.Printf("D2 Part 2: %v\n", answer)
}
