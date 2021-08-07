package utils_test

import (
	"github.com/stretchr/testify/assert"
	. "ova-animal-api/internal/utils"
	"testing"
)

func TestZeroSize(t *testing.T) {
	result, err := BatchSplit([]int{}, 10)
	assert.Nil(t, err)
	assert.Empty(t, result, "Should be empty slice")
}

func TestNilInput(t *testing.T) {
	result, err := BatchSplit(nil, 10)
	assert.Nil(t, err)
	assert.Empty(t, result, "Should be empty slice")
}

func TestZeroSliceCount(t *testing.T) {
	_, err := BatchSplit([]int{1, 2, 3}, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "can't split slice to 0 pieces", err.Error())
}

func TestSinglePiece(t *testing.T) {
	result, err := BatchSplit([]int{1, 2}, 2)
	assert.Nil(t, err)
	assert.Equal(t, [][]int{{1, 2}}, result)
}

func TestSinglePieceLargerThenSource(t *testing.T) {
	result, err := BatchSplit([]int{1, 2}, 100)
	assert.Nil(t, err)
	assert.Equal(t, [][]int{{1, 2}}, result)
}

func TestDividesExactly(t *testing.T) {
	result, err := BatchSplit([]int{1, 2, 3, 4, 5, 6}, 2)
	assert.Nil(t, err)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, result)
}

func TestNotDividesExactly(t *testing.T) {
	result, err := BatchSplit([]int{1, 2, 3, 4, 5}, 2)
	assert.Nil(t, err)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, result)
}

func TestSourceAndResultAreIndependent(t *testing.T) {
	src := []int{1, 2, 3}
	result, err := BatchSplit(src, 2)
	assert.Nil(t, err)

	src[0] = 100

	assert.Equal(t, 1, result[0][0])
}
