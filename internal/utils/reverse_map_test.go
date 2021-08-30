package utils_test

import (
	"github.com/stretchr/testify/assert"
	. "ova-animal-api/internal/utils"
	"testing"
)

func TestNil(t *testing.T) {
	result, err := ReverseMap(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "reverse nil map", err.Error())
	assert.Nil(t, result)
}

func TestRegulargCase(t *testing.T) {
	result, err := ReverseMap(map[int]string{
		1: "a",
		2: "b",
		3: "c",
	})
	assert.Nil(t, err)
	assert.Equal(t, map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}, result)
}

func TestDuplicates(t *testing.T) {
	_, err := ReverseMap(map[int]string{
		1: "a",
		2: "b",
		3: "a",
	})

	assert.NotNil(t, err)
	assert.True(t,
		err.Error() == "key a already present in target map with value 1" ||
			err.Error() == "key a already present in target map with value 3",
	)
}
