package animal

import (
	"ova-animal-api/internal/domain"
)

//go:generate mockgen -source=animal_repo.go -destination=animal_repo_mock.go -package=animal
type Repo interface {
	AddEntities(entities []domain.Animal) error
	ListEntities(limit, offset uint64) ([]domain.Animal, error)
	DescribeEntity(entityId uint64) (*domain.Animal, error)
}

func New() (Repo, error) {
	r := repo{}
	return &r, nil
}

type repo struct {
}

func (this *repo) AddEntities(entities []domain.Animal) error {
	panic("not implemented")
}

func (this *repo) ListEntities(limit, offset uint64) ([]domain.Animal, error) {
	panic("not implemented")
}

func (this *repo) DescribeEntity(entityId uint64) (*domain.Animal, error) {
	panic("not implemented")
}
