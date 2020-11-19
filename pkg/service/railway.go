package service

import (
	"context"
	"github.com/gocolly/colly/v2"
	railwayCrawler "go-transportation-bot/pkg/crawler/railway"
	"go-transportation-bot/pkg/railway"
	"go-transportation-bot/pkg/repository"
	"k8s.io/klog/v2"
	"strings"
	"time"
)

type railwayService struct {
	railwayRepo railway.RailwayRepository
}

func (r railwayService) GetCity(ctx context.Context, cityId string) (*railway.City, error) {
	return r.railwayRepo.GetCity(ctx, cityId)
}

func NewRailwayService(railwayRepo railway.RailwayRepository) railway.RailwayService {
	return &railwayService{
		railwayRepo: railwayRepo,
	}
}

func (r *railwayService) GetAllCity(ctx context.Context) ([]*railway.City, error) {
	cities, err := r.railwayRepo.GetAllCity(ctx)
	if err != nil && err.Error() == string(repository.GetCityIdNotFound) {
		klog.Warning(err)
		cities = reScanCity(r.railwayRepo)
		putAllCity(ctx, r.railwayRepo, cities)
	} else if err != nil {
		klog.Warning(err)
		return nil, err
	}
	return cities, nil
}

func (r *railwayService) GetStation(ctx context.Context, cityID string) ([]*railway.Station, error) {
	return getStation(r.railwayRepo, cityID), nil
}

func (r *railwayService) ScanCity(ctx context.Context) ([]*railway.City, error) {
	cityList := reScanCity(r.railwayRepo)
	if err := putAllCity(ctx, r.railwayRepo, cityList); err != nil {
		klog.Warning("PutAllCity to cache server error :")
		return nil, err
	}
	return cityList, nil
}

func (r *railwayService) GetInfoByStation(ctx context.Context, infoRequest *railway.InfoRequest) ([]*railway.Info, error) {
	//cityList := reScanCity(r.railwayRepo)
	//if err := putAllCity(ctx, r.railwayRepo, cityList); err != nil {
	//	klog.Warning("PutAllCity to cache server error :")
	//	return nil, err
	//}
	trainInfoList := []*railway.Info{
		{
			FromId:      "HIJK",
			ToId:        "ABCD",
			TrainNumber: "1234",
			FromTimes:   time.Now(),
			ToTimes:     time.Now().Add(time.Hour),
			Type:        railway.TIMES,
		},
	}

	return trainInfoList, nil
}

func reScanCity(railwayCrawler railwayCrawler.RailwayCrawler) []*railway.City {
	cityList := make([]*railway.City, 0)
	railwayCrawler.ScanCity(func(e *colly.HTMLElement) {
		city := &railway.City{
			Name: e.Text,
			Id:   e.Attr("data-type"),
		}
		cityList = append(cityList, city)
	})
	return cityList
}

func putAllCity(ctx context.Context, railwayRepository railway.RailwayRepository, cities []*railway.City) error {
	return railwayRepository.PutAllCity(ctx, cities)
}

func getStation(railwayCrawler railwayCrawler.RailwayCrawler, cityID string) []*railway.Station {
	stationList := make([]*railway.Station, 0)
	railwayCrawler.ScanStationByCityID(cityID, func(e *colly.HTMLElement) {

		station := &railway.Station{
			Name: e.Text,
			Id:   strings.Split(e.Attr("title"), "-")[0],
		}
		stationList = append(stationList, station)
	})

	return stationList
}
