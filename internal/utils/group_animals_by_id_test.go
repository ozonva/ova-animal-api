package utils_test

import (
	"github.com/stretchr/testify/assert"
	"ova-animal-api/internal/domain"
	. "ova-animal-api/internal/utils"
	"testing"
)

func TestGroupByIdNil(t *testing.T) {
	result, err := GroupById(nil)
	assert.Nil(t, err)
	assert.Equal(t, map[uint64]domain.Animal{}, result)
}

func TestGroupByIdEmpty(t *testing.T) {
	result, err := GroupById([]domain.Animal{})
	assert.Nil(t, err)
	assert.Equal(t, map[uint64]domain.Animal{}, result)
}

func TestGroupById(t *testing.T) {
	result, err := GroupById([]domain.Animal{animal1, animal2, animal3})
	assert.Nil(t, err)
	assert.Equal(t, map[uint64]domain.Animal{
		1: animal1,
		2: animal2,
		3: animal3,
	}, result)
}

func TestGroupByIdDuplicate(t *testing.T) {
	result, err := GroupById([]domain.Animal{animal1, animal2, animal1})
	assert.NotNil(t, err)
	assert.Equal(t, "duplicate key 1", err.Error())
	assert.Nil(t, result)
}
