package railway

import (
	"bytes"
	"github.com/gocolly/colly/v2"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const TestUrl = "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"
const (
	mockTestCityID = "<!DOCTYPE html>\n<html>\n<body>\n\n<div class=\"line-inner\">\n\t                <div class=\"line-head\">\n\t                 \t\t\t            <a href=\"#mainline\" aria-controls=\"mainline\" role=\"tab\" data-toggle=\"tab\" title=\"縣市\">縣市</a>\n\t\t\t        \t\t\t        \t\t\t        \t\t\t        \t\t\t        \t\t\t        </div>\n\t                <ul>\t  \n\t                    \t\t\t\t\t       <li><button type=\"button\" class=\"btn tipCity active\" title=\"常用\" data-type=\"cityHot\">常用</button></li>\n\t\t\t\t\t    \t\t                \n\t\t                \t\t                \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"基隆市\" data-type=\"city10017\">基隆市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"新北市\" data-type=\"city65000\">新北市</button></li>\n\t        \t\t\t\t\t        \t                    \t\t                    \t\t\n</body>\n</html>\n"
)

func TestNewRailwayCrawle(t *testing.T) {
	type args struct {
		scanUrl string
	}
	tests := []struct {
		name string
		args args
		want RailwayCrawler
	}{
		{
			name: "Check base URL",
			args: args{TestUrl},
			want: &RailwayCrawle{TestUrl},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRailwayCrawle(tt.args.scanUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRailwayCrawler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRailwayCrawleScanCity(t *testing.T) {
	const endPoint = "/city"

	testData := map[string]string{
		"city10017": "基隆市",
		"city65000": "新北市",
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("Request method should be Get")
		}

		if r.RequestURI != endPoint {
			t.Error("Request URI should be " + endPoint)
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(mockTestCityID))

	}))
	defer testServer.Close()
	var crawler RailwayCrawler
	crawler = &RailwayCrawle{
		testServer.URL + endPoint,
	}
	crawler.ScanCity(func(e *colly.HTMLElement) {
		if testData[e.Attr("data-type")] != e.Text {
			t.Error("Scan city ID Error")
		}
	})
}

func TestScanCity(t *testing.T) {
	const endPoint = "city"

	testData := map[string]string{
		"city10017": "基隆市",
		"city65000": "新北市",
	}

	type args struct {
		api string
		f   func(e *colly.HTMLElement)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "scan city id ",
			args: args{
				api: endPoint,
				f: func(e *colly.HTMLElement) {
					if testData[e.Attr("data-type")] != e.Text {
						t.Error("Scan city ID Error")
					}
				},
			},
		},
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "Get" {
			t.Error("Request method should be Get")
		}

		if r.RequestURI != endPoint {
			t.Error("Request URI should be " + endPoint)
		}
		_ = r.Write(bytes.NewBufferString(mockTestCityID))

		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanCity(tt.args.api, tt.args.f)
		})
	}
}
