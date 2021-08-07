package utils

import (
	"fmt"
)

func BatchSplit(src []int, size uint) ([][]int, error) {
	if size == 0 {
		return nil, fmt.Errorf("can't split slice to 0 pieces")
	}
	if src == nil {
		// I think, it's go-way decision to return empty slice on nil input like `append` works
		return [][]int{}, nil
	}

	var (
		result     [][]int
		sliceCount uint
	)

	sliceCount = batchSize(uint(len(src)), size)
	result = make([][]int, sliceCount)

	for i := range result {
		startIndex := i * int(size)
		endIndex := min(startIndex+int(size), len(src))

		partOfSource := src[startIndex:endIndex]

		result[i] = make([]int, len(partOfSource))
		copy(result[i], partOfSource)
	}

	return result, nil
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func batchSize(srcLen uint, size uint) uint {
	if srcLen%size == 0 {
		return srcLen / size
	}
	return srcLen/size + 1
}
