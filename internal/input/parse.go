package input

import (
	"os"
)

var newLine byte = 10

func getCount[T comparable](slice []T, match T) uint {
	var count uint
	for _, v := range slice {
		if v == match {
			count++
		}
	}
	return count
}

func GetRows(path string) [][]byte {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileLen := len(f)
	rows := make([][]byte, 0, getCount(f, newLine))

	var lastIdx int
	for idx := range fileLen {
		if f[idx] != newLine {
			continue
		}

		rows = append(rows, f[lastIdx:idx])
		lastIdx = idx + 1
	}

	if lastIdx < fileLen {
		rows = append(rows, f[lastIdx:])
	}

	return rows
}
