package railway

import (
	"github.com/gocolly/colly/v2"
	"k8s.io/klog/v2"
)

type RailwayCrawle struct {
	baseURL string
}

type RailwayCrawler interface {
	ScanCity(f func(e *colly.HTMLElement))
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
func scanCity(api string, f func(e *colly.HTMLElement)) {
	var c = colly.NewCollector()

	c.OnHTML(".line-inner li button[class='btn tipCity']", f)

	c.OnRequest(func(r *colly.Request) {
		klog.Info("scan city with url:", api)
	})

	c.Post(api, nil)
}
