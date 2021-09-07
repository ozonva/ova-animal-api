package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"math"
	"ova-animal-api/internal/converter"
	"ova-animal-api/internal/domain"
	"ova-animal-api/internal/repo/animal"
	"ova-animal-api/pkg/ova-animal-api/github.com/ozonva/ova-animal-api/api"
)

func NewHandler(repo animal.Repo) api.AnimalApiServer {
	return &AnimalApiServerImpl{
		repo: repo,
	}
}

type AnimalApiServerImpl struct {
	api.UnimplementedAnimalApiServer
	repo animal.Repo
}

func (this AnimalApiServerImpl) CreateEntity(_ context.Context, in *api.Animal) (*empty.Empty, error) {
	log.Info().Msgf("Entity %v created!", in)
	return &empty.Empty{}, this.repo.AddEntities([]domain.Animal{*converter.ConvertToDomain(in)})
}

func (this AnimalApiServerImpl) DescribeEntity(_ context.Context, in *api.IdRequest) (*api.Animal, error) {
	log.Info().Msgf("Entity %d described!", in.Id)
	entity, err := this.repo.DescribeEntity(in.Id)
	if err != nil {
		return nil, err
	}
	return converter.ConvertToApi(entity), err
}
func (this AnimalApiServerImpl) ListEntities(_ context.Context, in *empty.Empty) (*api.AnimalListResponse, error) {
	log.Info().Msgf("Entities listed!")
	entities, err := this.repo.ListEntities(math.MaxInt32, 0)
	if err != nil {
		return nil, err
	}

	resp := api.AnimalListResponse{}
	resp.Animal = make([]*api.Animal, 0)
	for _, animal := range entities {
		resp.Animal = append(resp.Animal, converter.ConvertToApi(&animal))
	}
	return &resp, nil
}
func (this AnimalApiServerImpl) RemoveEntity(_ context.Context, in *api.IdRequest) (*empty.Empty, error) {
	log.Info().Msgf("Entity %d removed!", in.Id)
	return &empty.Empty{}, this.repo.Delete(in.Id)
}
