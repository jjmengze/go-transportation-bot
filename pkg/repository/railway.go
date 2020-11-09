package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	CityId = "CITY_ID"
)

const (
	GetCityIdNotFound = iota
	//add any type you want
)

type RailwayRepositoryError struct {
	error
	errorType int
	message   string
}

func (e *RailwayRepositoryError) Error() string {
	return string(e.errorType)
}

func (r *railwayRepository) PutAllCity(ctx context.Context, cities []*railway.City) error {
	for i := 0; i < len(cities); i++ {
		data, err := json.Marshal(cities[i])
		if err != nil {
			klog.Warning("Marshal cities data Error: ", err)
			return err
		}
		err = r.HSet(ctx, CityId, cities[i].Name, data).Err()
		if err != nil {
			klog.Warning("Redis HSet Error: ", err)
			return err
		}
	}
	expiration := 24 * time.Hour * 10

	if err := r.Expire(ctx, CityId, expiration).Err(); err != nil {
		klog.Warning("Redis HSet Error: ", err)
		return err
	}
	return nil
}

func NewRailwayRepository(client redis.UniversalClient, crawler crawler.RailwayCrawler) railway.RailwayRepository {
	return &railwayRepository{client, crawler}
}

func (r *railwayRepository) GetCity(ctx context.Context, stationID string) (*railway.City, error) {
	city, err := r.HGet(ctx, CityId, stationID).Result()
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

func (r *railwayRepository) GetAllCity(ctx context.Context) ([]*railway.City, error) {
	if ok, err := r.Keys(ctx, CityId).Result(); err != nil {
		return nil, err
	} else if len(ok) == 0 {
		//	the key is not exits we should rescan city again
		return nil, &RailwayRepositoryError{
			errorType: GetCityIdNotFound,
			message:   fmt.Sprintf("Can't found the key %v in the cache server", CityId),
		}
	}

	allCities, err := r.HGetAll(ctx, CityId).Result()
	if err != nil {
		return nil, err
	}
	cities := make([]*railway.City, len(allCities))
	idx := 0
	for _, city := range allCities {
		c := &railway.City{}
		if err := json.Unmarshal(bytes.NewBufferString(city).Bytes(), c); err != nil {
			klog.Warning("Unmarshal city data Error: ", err)
			return nil, err
		}
		cities[idx] = c
		idx++
	}

	return cities, nil
}

func (r *railwayRepository) ScanCityFunc(f func(e *colly.HTMLElement)) {
	r.ScanCity(f)
}
