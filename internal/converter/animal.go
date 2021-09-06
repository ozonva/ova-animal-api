package converter

import (
	"ova-animal-api/internal/domain"
	"ova-animal-api/pkg/ova-animal-api/github.com/ozonva/ova-animal-api/api"
)

func ConvertToDomain(in *api.Animal) *domain.Animal {
	return &domain.Animal{
		Id:     in.Id,
		UserId: in.UserId,
		Name:   in.Name,
		Type:   domain.AnimalType(in.Type),
	}
}

func ConvertToApi(in *domain.Animal) *api.Animal {
	return &api.Animal{
		Id:     in.Id,
		UserId: in.UserId,
		Name:   in.Name,
		Type:   api.Animal_Type(in.Type),
	}
}
