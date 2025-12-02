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
	// id1, id2 := ids[0], ids[1]
	id1, _ = strconv.Atoi(string(sepIds[0]))
	id2, _ = strconv.Atoi(string(sepIds[1]))
	return
}

func Part1() {
	rngs := input.Seperate(input.GetFile("./internal/d2/example.txt"), comma)

	var answer int
	for _, rng := range rngs {
		sepIds := input.Seperate(rng, dash)
		id1, id2 := sepIds[0], sepIds[1]
		id1Num, _ := strconv.Atoi(string(id1))
		id2Num, _ := strconv.Atoi(string(id2))
		// dumb solution iterate each
		for i := id1Num; i <= id2Num; i++ {
			// mid := len(i)
		}
	}

	fmt.Printf("D2 Part 1: %v\n", answer)
}
