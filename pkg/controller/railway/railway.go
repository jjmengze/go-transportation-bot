package railway

import (
	"context"
	"go-transportation-bot/apis/railway/grpc/v1beta1"
	"go-transportation-bot/pkg/railway"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
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
	cityList, err := c.railwaySvc.GetAllCity(ctx)
	//cityList, err := c.railwaySvc.ScanCity(ctx)
	if err != nil {
		klog.Warning("Get City Error", err)
		return nil, err
	}
	respCityList := make([]*v1beta1.City, len(cityList))
	for i := 0; i < len(cityList); i++ {
		respCityList[i] = &v1beta1.City{
			ID:   cityList[i].Id,
			Name: cityList[i].Name,
		}
	}
	return &v1beta1.CityResponse{
		City: respCityList,
	}, nil
}

func (c *RailwayController) GetStationByCityID(ctx context.Context, cityRequest *v1beta1.City) (*v1beta1.StationResponse, error) {
	stationList, err := c.railwaySvc.GetStation(ctx, cityRequest.ID)
	if err != nil {
		klog.Warning("Get City Error", err)
		return nil, err
	}
	respStationList := make([]*v1beta1.Station, len(stationList))
	for i := 0; i < len(stationList); i++ {
		respStationList[i] = &v1beta1.Station{
			Id:   stationList[i].Id,
			Name: stationList[i].Name,
		}
	}
	return &v1beta1.StationResponse{
		Station: respStationList,
	}, nil
}
