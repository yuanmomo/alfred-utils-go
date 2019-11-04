package command

// Package is called aw
import (
	"alfred-utils-go/common"
	"fmt"
	aw "github.com/deanishe/awgo"
	"log"
	"time"
)

type SearchIPLocationCommand struct{}

type IpApiResult struct {
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

var (
	cacheName   = "ip-cache.json" // Filename of cached repo list
	maxCacheAge = 24 * time.Hour  // How long to cache repo list for
)

func (c *SearchIPLocationCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	var ip string
	if len(args) >= 1 { // has ip arg
		ip = args[0]
	}

	log.Printf("[main] search ip=[%s], current cache dir:[%s]", ip, wf.Cache.Dir)

	// check the input ip is valid
	if ip != "" && ! common.IsIpV4(ip) {
		log.Printf("Entering ip ....")
		wf.NewItem(fmt.Sprintf("Entering ip %s ...", ip)).Valid(true).Subtitle(ip)
		return wf
	}

	var jsonResult *IpApiResult

	if ip == "" { // search own ip location
		searchOnline(ip, &jsonResult);
		log.Print("Search own location online with result : ", jsonResult)
	} else {
		// If the cache has expired, set Rerun (which tells Alfred to re-run the
		// workflow), and start the background update process if it isn't already
		// running.
		if wf.Cache.Expired(cacheName, maxCacheAge) {
			wf.ClearCache();
		}

		// Try to load Results
		results := []*IpApiResult{}
		if wf.Cache.Exists(cacheName) {
			if err := wf.Cache.LoadJSON(cacheName, &results); err != nil {
				wf.FatalError(err)
			}
		}

		for _, res := range results {
			if res == nil || res.Query != ip {
				continue
			}
			log.Print("Find in cache : ", jsonResult)
			jsonResult = res;
		}

		if jsonResult == nil { // cannot find in local cache
			// search online
			searchOnline(ip, &jsonResult);
			log.Print("No cache, search online with result : ", jsonResult)
			results = append(results, jsonResult);
			wf.Cache.StoreJSON(cacheName, results);
		}
	}

	if jsonResult == nil {
		title := fmt.Sprintf("Search error")
		wf.NewItem(title).Valid(true).Arg(title).Subtitle(ip)
	} else {
		title := fmt.Sprintf("%s, %s, %s", jsonResult.Country, jsonResult.RegionName, jsonResult.City)
		wf.NewItem(title).Valid(true).Arg(title).Subtitle(fmt.Sprintf("%v, %v",jsonResult.Isp,jsonResult.Org))
	}

	// Send results to Alfred
	return wf
}

func searchOnline(ip string, jsonResult interface{}) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	log.Printf(ip)
	common.HttpGetJSON(url, &jsonResult)
}

func init() {
	RegisterCommand(&SearchIPLocationCommand{})
}
