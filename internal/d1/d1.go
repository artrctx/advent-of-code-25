package d1

import (
	"fmt"
	"strconv"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

var L byte = 76
var R byte = 82

func Part1() {
	rows := input.GetRows("./internal/d1/input.txt")

	var zeroCount int
	dial := 50
	for _, row := range rows {
		dir := row[0]
		val, _ := strconv.Atoi(string(row[1:]))

		if dir == L {
			dial -= val
		} else {
			dial += val
		}

		dial = dial % 100
		if dial < 0 {
			dial = 100 + dial
		}

		if dial == 0 {
			zeroCount++
		}
	}

	fmt.Printf("D1 Part 1: %v\n", zeroCount)
}

// 6671
func Part2() {
	rows := input.GetRows("./internal/d1/input.txt")

	var zeroCount int
	dial := 50
	for _, row := range rows {
		dir := row[0]
		val, _ := strconv.Atoi(string(row[1:]))

		prevDial := dial
		spin, remain := val/100, val%100
		if dir == L {
			dial = dial - remain
		} else {
			dial = dial + remain
		}

		if spin > 0 {
			zeroCount += spin
		}

		if prevDial != 0 && (dial > 100 || dial < 0) {
			zeroCount++
		}

		dial = dial % 100
		if dial < 0 {
			dial = 100 + dial
		}

		if dial == 0 {
			zeroCount++
		}
	}

	fmt.Printf("D1 Part 2: %v\n", zeroCount)
}
