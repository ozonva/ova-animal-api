package utils

import (
	"log"
	"ova-animal-api/internal/domain"
)

func SplitToBulks(src []domain.Animal, size uint) [][]domain.Animal {
	if size == 0 {
		log.Println("Can't split slice to 0 pieces!")
		return nil
	}
	if src == nil {
		return [][]domain.Animal{}
	}

	sliceCount := batchSize(uint(len(src)), size)
	result := make([][]domain.Animal, sliceCount)

	for i := range result {
		startIndex := i * int(size)
		endIndex := min(startIndex+int(size), len(src))

		partOfSource := src[startIndex:endIndex]

		result[i] = make([]domain.Animal, len(partOfSource))
		copy(result[i], partOfSource)
	}

	return result
}
