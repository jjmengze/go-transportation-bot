package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-transportation-bot/pkg/railway"
)

type railwayRepository struct {
	redis.UniversalClient
}

func NewRailwayRepository(client redis.UniversalClient) railway.RailwayRepository {
	return &railwayRepository{client}
}

func (r *railwayRepository) GetCity(ctx context.Context, stationID string) (*railway.City, error) {
	r.Get(ctx, stationID)
	return nil, nil
}
