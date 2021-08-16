package domain_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	. "ova-animal-api/internal/domain"
	"testing"
)

func TestAnimalString(t *testing.T) {
	animal := Animal{
		Id:     1,
		UserId: 100,
		Type:   CAT,
		Name:   "Леопольд",
	}
	assert.Equal(t, "1 100 CAT Леопольд", animal.String())
}

func TestFishSay(t *testing.T) {
	fish := Animal{1, 100, "Иннокентий", FISH}

	defer func() {
		if message := recover(); message != nil {
			log.Println("Fish really do not speak!")
		}
	}()

	fish.Say()

	panic("This code should not perform because of fish can't say anything!")
}

func TestCatSay(t *testing.T) {
	cat := Animal{1, 100, "Леопольд", CAT}
	cat.Say()
}

func TestRename(t *testing.T) {
	cat := Animal{1, 100, "Леопольд", CAT}
	cat.Rename("Том")
	assert.Equal(t, "Том", cat.Name)
}

func TestIncorrectRename(t *testing.T) {
	cat := Animal{1, 100, "Леопольд", CAT}
	cat.IncorrectRename("Том")
	assert.Equal(t, "Леопольд", cat.Name)
}
