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

func TestRailwayCrawle_ScanCityWithFunc(t *testing.T) {
	type fields struct {
		baseURL string
	}
	type args struct {
		scanner RailwayCrawler
		f       func(e *colly.HTMLElement)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &RailwayCrawle{
				baseURL: tt.fields.baseURL,
			}
		})
	}
}

func TestRailwayCrawleScanCity(t *testing.T) {
	const endPoint = "city"

	testData := make(map[string]string){
		"city10017": "基隆市",
		"city65000": "新北市",
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "Get" {
			t.Error("Request method should be Get")
		}

		if r.RequestURI != endPoint {
			t.Error("Request URI should be " + endPoint)
		}

		_ = r.Write(bytes.NewBufferString(testCityID))

		w.WriteHeader(http.StatusInternalServerError)
	}))

	defer testServer.Close()
	var crawler RailwayCrawler
	crawler = &RailwayCrawle{
		testServer.URL + endPoint,
	}
	crawler.ScanCity(func(e *colly.HTMLElement) {
		e.Attr("data-type"))
		fmt.Println(e.Text)

	})
}

const (
	testCityID = "<div class=\"line-inner\">\n\t                <div class=\"line-head\">\n\t                 \t\t\t            <a href=\"#mainline\" aria-controls=\"mainline\" role=\"tab\" data-toggle=\"tab\" title=\"縣市\">縣市</a>\n\t\t\t        \t\t\t        \t\t\t        \t\t\t        \t\t\t        \t\t\t        </div>\n\t                <ul>\t  \n\t                    \t\t\t\t\t       <li><button type=\"button\" class=\"btn tipCity active\" title=\"常用\" data-type=\"cityHot\">常用</button></li>\n\t\t\t\t\t    \t\t                \n\t\t                \t\t                \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"基隆市\" data-type=\"city10017\">基隆市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"新北市\" data-type=\"city65000\">新北市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"臺北市\" data-type=\"city63000\">臺北市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"桃園市\" data-type=\"city68000\">桃園市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"新竹縣\" data-type=\"city10004\">新竹縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"新竹市\" data-type=\"city10018\">新竹市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"苗栗縣\" data-type=\"city10005\">苗栗縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"臺中市\" data-type=\"city66000\">臺中市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"彰化縣\" data-type=\"city10007\">彰化縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"南投縣\" data-type=\"city10008\">南投縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"雲林縣\" data-type=\"city10009\">雲林縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"嘉義縣\" data-type=\"city10010\">嘉義縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"嘉義市\" data-type=\"city10020\">嘉義市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"臺南市\" data-type=\"city67000\">臺南市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"高雄市\" data-type=\"city64000\">高雄市</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"屏東縣\" data-type=\"city10013\">屏東縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"臺東縣\" data-type=\"city10014\">臺東縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"花蓮縣\" data-type=\"city10015\">花蓮縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t\t                    \t\t\t\t\t            <li><button type=\"button\" class=\"btn tipCity\" title=\"宜蘭縣\" data-type=\"city10002\">宜蘭縣</button></li>\n\t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t\t\t\t\t        \t                    \t                    \t                </ul>\n\t            </div>"
)
