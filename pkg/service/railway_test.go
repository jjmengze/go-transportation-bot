package service

import (
	"bytes"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"go-transportation-bot/pkg/railway"
	"testing"
)

type mockRailwayRepo struct {
}

func (m *mockRailwayRepo) GetCity(ctx context.Context, stationID string) (*railway.City, error) {
	//todo lose get all city mock function
	return nil, nil
}

func (m *mockRailwayRepo) PutAllCity(ctx context.Context, cities []*railway.City) error {
	//todo lose put all city mock function
	return nil
}

func (r *railwayService) GetStation() ([]*railway.Station, error) {
	//todo lose get station mock function
	return nil, nil
}

func (m *mockRailwayRepo) ScanCity(f func(e *colly.HTMLElement)) () {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}

	in := `<!DOCTYPE html>
<html>
<body>

<div class="line-inner">
	                <div class="line-head">
	                 			            <a href="#mainline" aria-controls="mainline" role="tab" data-toggle="tab" title="縣市">縣市</a>
			        			        			        			        			        			        </div>
	                <ul>	  
	                    					       <li><button type="button" class="btn tipCity active" title="常用" data-type="cityHot">常用</button></li>
					    		                
		                		                		                    					            <li><button type="button" class="btn tipCity" title="基隆市" data-type="city10017">基隆市</button></li>
					        					        					        					        					        	                    		                    					            <li><button type="button" class="btn tipCity" title="新北市" data-type="city65000">新北市</button></li>
	        					        	                    		                    		
</body>
</html>
`
	sel := ".line-inner li button[class='btn tipCity']"

	doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(in)))
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			f(colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})
}

func (m *mockRailwayRepo) ScanStationByCityID(cityID string, f func(e *colly.HTMLElement)) () {
	//todo lose scan station by city id mock function
}

func (r *mockRailwayRepo) GetAllCity(ctx context.Context) ([]*railway.City, error) {
	//todo lose get all city test case
	return nil, nil
}

func TestRailwayServiceScanCity(t *testing.T) {
	type fields struct {
		railwayRepo railway.RailwayRepository
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test Scan City",
			fields: fields{&mockRailwayRepo{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &railwayService{
				railwayRepo: tt.fields.railwayRepo,
			}
			r.ScanCity(context.Background())
		})
	}
}
