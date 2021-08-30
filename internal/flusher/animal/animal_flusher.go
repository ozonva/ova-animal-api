package animal

import (
	"ova-animal-api/internal/domain"
	"ova-animal-api/internal/repo/animal"
)

//go:generate mockgen -source=animal_flusher.go -destination=animal_flusher_mock.go -package=animal
type Flusher interface {
	Flush(entities []domain.Animal) []domain.Animal
}

func New(repo animal.Repo) Flusher {
	f := flusher{repo}
	return &f
}

type flusher struct {
	repo animal.Repo
}

// Flush persists entities and return not persistable (for example, all entities have no
// mandatory fields will be returned and all other saved)
func (this *flusher) Flush(entities []domain.Animal) []domain.Animal {

	// Если какой-то элемент вызывает ошибку вставки, действуем методом дихотомии.
	// Это хорошо сработает, когда в большой коллекции мало "плохих" элементов.
	// В противном случае лучше перебирать по-одному
	if err := this.repo.AddEntities(entities); err != nil {

		if len(entities) <= 1 {
			return entities
		}

		medianLen := len(entities) / 2
		left := entities[0:medianLen]
		right := entities[medianLen:]

		var notPersisted []domain.Animal
		notPersisted = append(notPersisted, this.Flush(left)...)
		notPersisted = append(notPersisted, this.Flush(right)...)
		return notPersisted
	}

	return []domain.Animal{}
}
