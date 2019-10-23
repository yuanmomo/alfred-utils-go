package command

// Package is called aw
import (
	"alfred-utils-go/common"
	"fmt"
	aw "github.com/deanishe/awgo"
)

type SearchIPLocationCommand struct{}

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


func (c *SearchIPLocationCommand) Name() string {
	return "lookip"
}

func (c *SearchIPLocationCommand) Description() Description {
	return Description{
		Short: "Search IP location",
		Usage: []string{"lip 8.8.8.8"},
	}
}

func (c *SearchIPLocationCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	var ip string
	//fmt.Println(os.Args)
	if len(args) >= 1 { // has ip arg
		ip = args[0]
	}
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)

	//fmt.Println(ip)
	var jsonResult *IpResult
	common.HttpGetJSON(url, &jsonResult)

	//fmt.Println(jsonResult)

	title := fmt.Sprintf("%s, %s, %s", jsonResult.Country, jsonResult.RegionName, jsonResult.City)
	wf.NewItem(title).Valid(true).Arg(title).Subtitle(jsonResult.Query)
	// Send results to Alfred
	return  wf
}

func init() {
	RegisterCommand(&SearchIPLocationCommand{})
}


