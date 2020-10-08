package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

const URL = "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"

func main() {

	//body := map[string]string{
	//	"_csrf":          "d6cecf75-c5c9-4027-8ab9e-36e8f1dc89a1",
	//	"startStation":   "7380-四腳亭",
	//	"endStation":     "1000-臺北",
	//	"transfer":       "ONE",
	//	"rideDate":       t.Format("2006/01/02"),
	//	"startOrEndTime": "true",
	//	"startTime":      "00:00",
	//	"endTime":        "23:59",
	//	"trainTypeList":  "ALL",
	//	"query":          "查詢",
	//}

	//c.OnHTML(".itinerary-controls", func(e *colly.HTMLElement) {
	//	//e.Request.Visit(e.Attr("href"))
	//	e.ForEach(".trip-column", func(i int, element *colly.HTMLElement) {
	//		//fmt.Printf("index : %d ,vale %s \n", i, element.Text)
	//		//fmt.Println(element.ChildText(".train-number"))
	//		data := element.ChildTexts("td")
	//		for _, v := range data {
	//			fmt.Printf("idx: %v ", v)
	//		}
	//		fmt.Println()
	//	})
	//})
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	//c.Post(URL, body)
	scanCity()
}
func scanCity() {
	//var t = time.Now()
	var c = colly.NewCollector()
	c.OnHTML(".line-inner li button[class='btn tipCity']", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("title"))
		fmt.Println(e.Attr("data-type"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Post(URL, nil)
}
