package main

// Package is called aw
import (
	"alfred-utils-go/common"
	"fmt"
	"github.com/deanishe/awgo"
	"os"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

type IpResult struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

// Your workflow starts here
func run() {

	var ip string
	//fmt.Println(os.Args)
	if len(os.Args) > 1 { // has ip arg
		ip = os.Args[1]
	}
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)

	//fmt.Println(ip)
	var jsonResult *IpResult
	http_util.HttpGetJSON(url, &jsonResult)

	//fmt.Println(jsonResult)

	title := fmt.Sprintf("%s, %s, %s", jsonResult.Country, jsonResult.RegionName, jsonResult.City)

	wf.NewItem(title).Valid(true).Arg(title).Subtitle(jsonResult.Query)

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
