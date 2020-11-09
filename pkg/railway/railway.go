package railway

import (
	"context"
	"go-transportation-bot/pkg/crawler/railway"
)

type Railway struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	ArriveTime string `json:"arrive_time"`
	TotalTime  string `json:"total_time"`
	Price      string `json:"price"`
}

type City struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type RailwayRepository interface {
	PutAllCity(ctx context.Context, cities []*City) error
	railway.RailwayCrawler
	GetCity(ctx context.Context, stationID string) (*City, error)
	//GetStation(ctx context.Context, d *City) error
	//GetStatus(ctx context.Context) error
}
type RailwayService interface {
	GetCity(ctx context.Context, cityID string) (*City, error)
	ScanCity(ctx context.Context)
}
