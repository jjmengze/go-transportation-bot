package railway

import (
	"context"
	"go-transportation-bot/apis/railway/grpc/v1beta1"
	"go-transportation-bot/pkg/railway"
	"google.golang.org/grpc"
)

type RailwayController struct {
	railwaySvc railway.RailwayService
	v1beta1.UnimplementedRailwayServer
}

func New(s *grpc.Server, railwaySvc railway.RailwayService) *RailwayController {
	c := &RailwayController{
		railwaySvc: railwaySvc,
	}
	v1beta1.RegisterRailwayServer(s, c)
	return c
}

func (c *RailwayController) GetCity(ctx context.Context, empty *v1beta1.Empty) (*v1beta1.CityResponse, error) {
	c.railwaySvc.ScanCity(ctx)
	return &v1beta1.CityResponse{
		City: []*v1beta1.City{
			{
				ID:   "100",
				Name: "Keelung",
			},
			{
				ID:   "101",
				Name: "taipei",
			},
			{
				ID:   "102",
				Name: "e.t.c.",
			},
		},
	}, nil
}
