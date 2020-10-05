package railway

import "context"

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
}

type RailwayRepository interface {
	GetCity(ctx context.Context, stationID string) (*City, error)
	//GetStation(ctx context.Context, d *City) error
	//GetStatus(ctx context.Context) error
}
type RailwayService interface {
	GetCity(ctx context.Context) ([]City, error)
}
