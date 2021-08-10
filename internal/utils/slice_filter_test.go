package utils_test

import (
	"github.com/stretchr/testify/assert"
	. "ova-animal-api/internal/utils"
	"testing"
)

func TestSliceFilterNil(t *testing.T) {
	filtered := SliceFilter(nil)
	assert.Equal(t, []string{}, filtered)
}

func TestSliceFilterEmpty(t *testing.T) {
	filtered := SliceFilter([]string{})
	assert.Equal(t, []string{}, filtered)
}

func TestSliceFilterAll(t *testing.T) {
	filtered := SliceFilter([]string{"cat", "dog", "rabbit"})
	assert.Equal(t, []string{}, filtered)
}

func TestSliceFilterNone(t *testing.T) {
	filtered := SliceFilter([]string{"shark", "lynx", "eagle"})
	assert.Equal(t, []string{"shark", "lynx", "eagle"}, filtered)
}

func TestSliceFilter(t *testing.T) {
	filtered := SliceFilter([]string{"cat", "", "shark", "dog", "lynx", "cat", "eagle"})
	assert.Equal(t, []string{"", "shark", "lynx", "eagle"}, filtered)
}
