package utils_test

import (
	"github.com/stretchr/testify/assert"
	"ova-animal-api/internal/domain"
	. "ova-animal-api/internal/utils"
	"testing"
)

var animal1 = domain.Animal{
	Id:     1,
	Name:   "Барсик",
	Type:   domain.CAT,
	UserId: 100,
}
var animal2 = domain.Animal{
	Id:     2,
	Name:   "Пират",
	Type:   domain.DOG,
	UserId: 101,
}
var animal3 = domain.Animal{
	Id:     3,
	Name:   "Поликарп",
	Type:   domain.FISH,
	UserId: 100,
}
var animal4 = domain.Animal{
	Id:     4,
	Name:   "Кнопочка",
	Type:   domain.DOG,
	UserId: 101,
}
var animal5 = domain.Animal{
	Id:     1,
	Name:   "Том",
	Type:   domain.CAT,
	UserId: 100,
}
var animal6 = domain.Animal{
	Id:     2,
	Name:   "Джерри",
	Type:   domain.MOUSE,
	UserId: 101,
}

func TestZeroSizeAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{}, 10)
	assert.Empty(t, result, "Should be empty slice")
}

func TestNilInputAnimal(t *testing.T) {
	result := SplitToBulks(nil, 10)
	assert.Empty(t, result, "Should be empty slice")
}

func TestZeroSliceCountAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{
		animal1,
		animal2,
	}, 0)
	assert.Nil(t, result)
}

func TestSinglePieceAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{animal1, animal2}, 2)
	assert.Equal(t, [][]domain.Animal{{animal1, animal2}}, result)
}

func TestSinglePieceLargerThenSourceAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{animal1, animal2}, 100)
	assert.Equal(t, [][]domain.Animal{{animal1, animal2}}, result)
}

func TestDividesExactlyAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{animal1, animal2, animal3, animal4, animal5, animal6}, 2)
	assert.Equal(t, [][]domain.Animal{{animal1, animal2}, {animal3, animal4}, {animal5, animal6}}, result)
}

func TestNotDividesExactlyAnimal(t *testing.T) {
	result := SplitToBulks([]domain.Animal{animal1, animal2, animal3, animal4, animal5}, 2)
	assert.Equal(t, [][]domain.Animal{{animal1, animal2}, {animal3, animal4}, {animal5}}, result)
}

func TestSourceAndResultAreIndependentAnimal(t *testing.T) {
	src := []domain.Animal{animal1, animal2, animal3}
	result := SplitToBulks(src, 2)

	src[0] = animal4

	assert.Equal(t, animal1, result[0][0])
}
