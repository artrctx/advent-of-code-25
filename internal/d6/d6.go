package d6

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

func getNumbersPart1(nums []byte) []int {
	newNums := make([]int, 0, len(nums))
	lastIdx := -1
	for idx, v := range nums {
		if v == ' ' {
			if lastIdx != -1 {
				n, _ := strconv.Atoi(string(nums[lastIdx:idx]))
				newNums = append(newNums, n)
				lastIdx = -1
			}
			continue
		}

		if lastIdx == -1 {
			lastIdx = idx
		}
	}

	if lastIdx != -1 {
		n, _ := strconv.Atoi(string(nums[lastIdx:]))
		newNums = append(newNums, n)
	}

	return slices.Clip(newNums)
}

func getOperators(ops []byte) []byte {
	newOps := make([]byte, 0, len(ops))
	for idx, v := range ops {
		if v == ' ' {
			continue
		}
		newOps = append(newOps, ops[idx])
	}
	return slices.Clip(newOps)
}

func Part1() {
	rows := input.GetRows("./internal/d6/input.txt")

	rLen := len(rows)
	ops := getOperators(rows[rLen-1])

	var nums [][]int
	for _, r := range rows[:rLen-1] {
		nums = append(nums, getNumbersPart1(r))
	}

	var answer int

	for idx, op := range ops {
		var val int
		if op == '*' {
			val = 1
		}

		for nIdx := range len(nums) {
			if op == '*' {
				val *= nums[nIdx][idx]
			} else {
				val += nums[nIdx][idx]
			}
		}

		answer += val
	}

	fmt.Printf("D6 Part 1: %v\n", answer)
}

func getNumbersPart2(input [][]byte) [][]int {
	numList := make([][]int, 0, len(input))

	var entrySet []int
	running := false
	for idx := range len(input[0]) {
		entry := make([]byte, 0, 4)
		for rIdx := range len(input) {
			v := input[rIdx][idx]
			if v == ' ' {
				continue
			}
			entry = append(entry, v)
		}

		if len(entry) != 0 {
			if !running {
				running = true
			}
			n, _ := strconv.Atoi(string(entry))
			entrySet = append(entrySet, n)
		} else if running {
			numList = append(numList, entrySet)
			entrySet = []int{}
			running = false
		}
	}

	if len(entrySet) > 0 {
		numList = append(numList, entrySet)
	}

	return slices.Clip(numList)
}

// 7858778299780 too low
func Part2() {
	rows := input.GetRows("./internal/d6/example.txt")

	rLen := len(rows)
	ops := getOperators(rows[rLen-1])

	nums := getNumbersPart2(rows[:rLen-1])
	var answer int

	for idx, ns := range nums {
		var val int
		if ops[idx] == '*' {
			val = 1
		}
		for _, v := range ns {
			if ops[idx] == '*' {
				val *= v
			} else {
				val += v
			}
		}
		answer += val
	}

	fmt.Printf("D6 Part 2: %v\n", answer)
}
