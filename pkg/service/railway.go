package service

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
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

func (r *railwayService) ScanCity() {
	r.railwayRepo.ScanCity(func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("data-type"))
		fmt.Println(e.Text)
	})
}
