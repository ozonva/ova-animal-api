package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"ova-animal-api/pkg/ova-animal-api/github.com/ozonva/ova-animal-api/api"
)

type AnimalApiServerImpl struct {
	api.UnimplementedAnimalApiServer
}

func (AnimalApiServerImpl) CreateEntity(_ context.Context, in *api.Animal) (*empty.Empty, error) {
	log.Info().Msgf("Entity %v created!", in)
	return &empty.Empty{}, nil
}
func (AnimalApiServerImpl) DescribeEntity(_ context.Context, in *api.IdRequest) (*api.Animal, error) {
	log.Info().Msgf("Entity %d described!", in.Id)
	return &api.Animal{}, nil
}
func (AnimalApiServerImpl) ListEntities(_ context.Context, in *empty.Empty) (*api.AnimalListResponse, error) {
	log.Info().Msgf("Entities listed!")
	return &api.AnimalListResponse{}, nil
}
func (AnimalApiServerImpl) RemoveEntity(_ context.Context, in *api.IdRequest) (*empty.Empty, error) {
	log.Info().Msgf("Entity %d removed!", in.Id)
	return &empty.Empty{}, nil
}
