package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gocolly/colly/v2"
	crawler "go-transportation-bot/pkg/crawler/railway"
	"go-transportation-bot/pkg/railway"
)

type railwayRepository struct {
	redis.UniversalClient
	crawler.RailwayCrawler
}

func (r *railwayRepository) PutAllCity(ctx context.Context, cities []railway.City) error {
	return nil
}

func NewRailwayRepository(client redis.UniversalClient, crawler crawler.RailwayCrawler) railway.RailwayRepository {
	return &railwayRepository{client, crawler}
}

func (r *railwayRepository) GetCity(ctx context.Context, stationID string) (*railway.City, error) {
	r.Get(ctx, stationID)
	return nil, nil
}

func (r *railwayRepository) ScanCityFunc(f func(e *colly.HTMLElement)) {
	r.ScanCity(f)
}
