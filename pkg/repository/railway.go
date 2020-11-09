package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gocolly/colly/v2"
	crawler "go-transportation-bot/pkg/crawler/railway"
	"go-transportation-bot/pkg/railway"
	"k8s.io/klog/v2"
	"time"
)

type railwayRepository struct {
	redis.UniversalClient
	crawler.RailwayCrawler
}

const (
	CITYID = "CITY_ID"
)

func (r *railwayRepository) PutAllCity(ctx context.Context, cities []*railway.City) error {
	for i := 0; i < len(cities); i++ {
		data, err := json.Marshal(cities[i])
		if err != nil {
			klog.Warning("Marshal cities data Error: ", err)
			return err
		}
		err = r.HSet(ctx, CITYID, cities[i].Name, data).Err()
		if err != nil {
			klog.Warning("Redis HSet Error: ", err)
			return err
		}
	}
	expiration := 24 * time.Hour * 10

	if err := r.Expire(ctx, CITYID, expiration).Err(); err != nil {
		klog.Warning("Redis HSet Error: ", err)
		return err
	}
	return nil
}

func NewRailwayRepository(client redis.UniversalClient, crawler crawler.RailwayCrawler) railway.RailwayRepository {
	return &railwayRepository{client, crawler}
}

func (r *railwayRepository) GetCity(ctx context.Context, stationID string) (*railway.City, error) {
	city, err := r.HGet(ctx, CITYID, stationID).Result()
	if err != nil {
		return nil, err
	}
	c := &railway.City{}
	if err := json.Unmarshal(bytes.NewBufferString(city).Bytes(), c); err != nil {
		klog.Warning("Unmarshal city data Error: ", err)
		return nil, err
	}

	return c, nil
}

func (r *railwayRepository) ScanCityFunc(f func(e *colly.HTMLElement)) {
	r.ScanCity(f)
}
