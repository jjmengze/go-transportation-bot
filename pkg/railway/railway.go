package railway

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go-transportation-bot/pkg/crawler/railway"
	"time"
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

type Station struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type InfoRequest struct {
	FromId      string
	ToId        string
	TrainNumber string
	FromTimes   *timestamp.Timestamp
	ToTimes     *timestamp.Timestamp
	Type        Type
}

type Info struct {
	FromId      string
	ToId        string
	TrainNumber string
	FromTimes   time.Time
	ToTimes     time.Time
	Type        Type
}

type Type int32

const (
	TIMES Type = iota
	STATION
	NUMBER
)

type RailwayRepository interface {
	PutAllCity(ctx context.Context, cities []*City) error
	railway.RailwayCrawler
	GetCity(ctx context.Context, stationID string) (*City, error)
	GetAllCity(ctx context.Context) ([]*City, error)
	//GetStation(ctx context.Context, d *City) error
	//GetStatus(ctx context.Context) error
}
type RailwayService interface {
	GetCity(ctx context.Context, cityID string) (*City, error)
	GetAllCity(ctx context.Context) ([]*City, error)
	ScanCity(ctx context.Context) ([]*City, error)
	GetStation(ctx context.Context, cityID string) ([]*Station, error)
	GetInfoByStation(ctx context.Context, infoRequest *InfoRequest) ([]*Info, error)
}
