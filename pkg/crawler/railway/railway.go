package railway

import (
	"github.com/gocolly/colly/v2"
	"k8s.io/klog/v2"
)

type RailwayCrawle struct {
}

type RailwayCrawler interface {
	ScanCity(f func(e *colly.HTMLElement))
}

const (
	URL = "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"
)

func NewRailwayCrawler() RailwayCrawler {
	return &RailwayCrawle{}
}

func (rc *RailwayCrawle) ScanCity(f func(e *colly.HTMLElement)) () {
	var c = colly.NewCollector()

	c.OnHTML(".line-inner li button[class='btn tipCity']", f)

	c.OnRequest(func(r *colly.Request) {
		klog.Info("scam city with url:", r.URL)
	})

	c.Post(URL, nil)
}
