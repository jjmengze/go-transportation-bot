package service

import (
	"context"
	"go-transportation-bot/pkg/railway"
)

type railwayService struct {
	railwayRepo railway.RailwayRepository
}

func (r railwayService) GetCity(ctx context.Context) ([]railway.City, error) {
	return nil, nil
}

func NewRailwayService(railwayRepo railway.RailwayRepository) railway.RailwayService {
	return &railwayService{
		railwayRepo: railwayRepo,
	}
}
