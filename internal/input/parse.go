package input

import (
	"os"
)

var newLine byte = 10

func GetCount[T comparable](slice []T, match T) uint {
	var count uint
	for _, v := range slice {
		if v == match {
			count++
		}
	}
	return count
}
func GetFile(path string) []byte {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return f
}

func GetRows(path string) [][]byte {
	return Seperate(GetFile(path), newLine)
}

func Seperate(input []byte, seperator byte) [][]byte {
	fileLen := len(input)
	rows := make([][]byte, 0, GetCount(input, seperator))

	var lastIdx int
	for idx := range fileLen {
		if input[idx] != seperator {
			continue
		}

		rows = append(rows, input[lastIdx:idx])
		lastIdx = idx + 1
	}

	if lastIdx < fileLen {
		rows = append(rows, input[lastIdx:])
	}

	return rows
}
