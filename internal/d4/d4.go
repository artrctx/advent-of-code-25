package d4

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/artrctx/advent-of-code-25/internal/input"
)

var (
	paper    byte     = '@'
	adjacent [][2]int = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
)

// func clamp(value, min, max int) int {
// 	return int(math.Min(math.Max(float64(value), float64(min)), float64(max)))
// }

func Part1() {
	rows := input.GetRows("./internal/d4/input.txt")

	var answer atomic.Int32
	var wg sync.WaitGroup
	rLen := len(rows)
	cLen := len(rows[0])
	for r := range rLen {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range cLen {
				if rows[r][c] != paper {
					continue
				}
				var pCount uint8
				for _, pos := range adjacent {
					newRowIdx, newColIdx := r+pos[1], c+pos[0]
					if newRowIdx < 0 || newColIdx < 0 || newRowIdx >= rLen || newColIdx >= cLen || rows[newRowIdx][newColIdx] != paper {
						continue
					}
					pCount++
					if pCount == 4 {
						break
					}
				}
				if pCount < 4 {
					answer.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("D4 Part 1: %v\n", answer.Load())
}

type RemovalSet struct {
	mu  sync.Mutex
	set [][2]int
}

func (rs *RemovalSet) Append(r, c int) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.set = append(rs.set, [2]int{r, c})
}

func Part2() {
	rows := input.GetRows("./internal/d4/input.txt")

	var answer int
	var wg sync.WaitGroup
	rLen, cLen := len(rows), len(rows[0])
	for {
		rs := RemovalSet{}
		for r := range rLen {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for c := range cLen {
					if rows[r][c] != paper {
						continue
					}
					var pCount uint8
					for _, pos := range adjacent {
						newRowIdx, newColIdx := r+pos[1], c+pos[0]
						if newRowIdx < 0 || newColIdx < 0 || newRowIdx >= rLen || newColIdx >= cLen || rows[newRowIdx][newColIdx] != paper {
							continue
						}
						pCount++
						if pCount == 4 {
							break
						}
					}
					if pCount < 4 {
						rs.Append(r, c)
					}
				}
			}()
		}
		wg.Wait()

		if len(rs.set) == 0 {
			break
		}
		answer += len(rs.set)
		for _, pos := range rs.set {
			rows[pos[0]][pos[1]] = '.'
		}
	}
	fmt.Printf("D4 Part 1: %v\n", answer)
}
