package railway

import (
	"github.com/gocolly/colly/v2"
	"k8s.io/klog/v2"
	"net/http"
)

type RailwayCrawle struct {
	baseURL string
}

type RailwayCrawler interface {
	ScanCity(f func(e *colly.HTMLElement))
	ScanStationByCityID(cityID string, f func(e *colly.HTMLElement))
}

const (
	URL = "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"
)

func NewRailwayCrawle(scanUrl string) RailwayCrawler {
	return &RailwayCrawle{
		baseURL: scanUrl,
	}
}

func (rc *RailwayCrawle) ScanCity(f func(e *colly.HTMLElement)) () {
	scanCity(rc.baseURL, f)
}

//todo:需要修改成英文
//實際到台鐵爬city的過程
//注意本操作是同步操作！
func scanCity(api string, f func(e *colly.HTMLElement)) error {
	var c = colly.NewCollector()

	c.OnHTML(".line-inner li button[class='btn tipCity']", f)

	c.OnRequest(func(r *colly.Request) {
		klog.Info("scan city with url:", api)
	})

	err := c.Request(http.MethodGet, api, nil, nil, nil)
	return err
}

func (rc *RailwayCrawle) ScanStationByCityID(cityID string, f func(e *colly.HTMLElement)) () {
	scanStationByCityID(rc.baseURL, cityID, f)
}

//todo:需要修改成英文
//實際到台鐵爬站名與編號的過程
//注意本操作是同步操作！
func scanStationByCityID(api string, cityID string, f func(e *colly.HTMLElement)) error {
	querySegment := "#" + cityID + " button"
	var c = colly.NewCollector()

	c.OnHTML(querySegment, f)

	c.OnRequest(func(r *colly.Request) {
		klog.Info("scan station with url:", api)
	})

	err := c.Request(http.MethodGet, api, nil, nil, nil)
	return err
}
