package utils

import (
	"errors"
	"fmt"
)

func ReverseMap(src map[int]string) (map[string]int, error) {
	if src == nil {
		return nil, errors.New("reverse nil map")
	}

	result := make(map[string]int)

	for k, v := range src {
		prevVal, alreadyExists := result[v]
		if alreadyExists {
			return nil, fmt.Errorf("key %s already present in target map with value %d", v, prevVal)
		}
		result[v] = k
	}

	return result, nil
}
