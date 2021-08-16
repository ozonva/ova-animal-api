package utils

import (
	"fmt"
	"ova-animal-api/internal/domain"
)

func GroupById(entities []domain.Animal) (map[uint64]domain.Animal, error) {
	if entities == nil {
		return map[uint64]domain.Animal{}, nil
	}

	result := map[uint64]domain.Animal{}

	for _, animal := range entities {
		id := animal.Id

		if _, exists := result[id]; exists {
			return nil, fmt.Errorf("duplicate key %d", id)
		}
		result[id] = animal
	}

	return result, nil
}
