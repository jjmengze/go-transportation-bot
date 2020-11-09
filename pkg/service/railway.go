package service

import (
	"context"
	"github.com/gocolly/colly/v2"
	"go-transportation-bot/pkg/railway"
	"k8s.io/klog/v2"
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

func (r *railwayService) ScanCity(ctx context.Context) {
	cityList := make([]*railway.City, 0)
	r.railwayRepo.ScanCity(func(e *colly.HTMLElement) {
		//fmt.Println(e.Attr("data-type"))
		//fmt.Println(e.Text)
		city := &railway.City{
			Name: e.Text,
			Id:   e.Attr("data-type"),
		}
		cityList = append(cityList, city)
	})
	if err := r.railwayRepo.PutAllCity(ctx, cityList); err != nil {
		klog.Warning("PutAllCity to cache server error :")
	}
}
