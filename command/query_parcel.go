package command

// Package is called aw
import (
	"alfred-utils-go/common"
	"fmt"
	aw "github.com/deanishe/awgo"
	mapset "github.com/deckarep/golang-set"
	"log"
	"time"
)

type ParcelResult struct {
	Status   int    `json:"status"`
	Info     string `json:"info"`
	Action   string `json:"action"`
	ShowTime int    `json:"show_time"`
	Data     struct {
		Shipment struct {
			ShipmentID                 string `json:"shipment_id"`
			ClientReference            string `json:"client_reference"`
			ExtNumber                  string `json:"ext_number"`
			OuterCarrierCode           string `json:"outer_carrier_code"`
			OuterCarrierTrackingNumber string `json:"outer_carrier_tracking_number"`
			Status                     string `json:"status"`
			Country                    string `json:"country"`
			Postcode                   string `json:"postcode"`
			Traces                     []struct {
				Time string `json:"time"`
				Info string `json:"info"`
			} `json:"traces"`
		} `json:"shipment"`
	} `json:"data"`
	GridViewReload int    `json:"grid_view_reload"`
	RequestTime    string `json:"request_time"`
}

type QueryParcelCommand struct{}

func (c *QueryParcelCommand) Name() string {
	return "parcel"
}

var (
	orderCacheName   = "order-info-cache.json" // Filename of cached repo list
	orderMaxCacheAge = 7 * 24 * time.Hour      // How long to cache repo list for
)

func (c *QueryParcelCommand) Description() Description {
	return Description{
		Short: "Query the parcel info to Sweden.",
		Usage: []string{
			"parcel order-number",
		},
	}
}

func (c *QueryParcelCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	var orderNumber string
	if len(args) >= 1 { // has order number arg
		orderNumber = args[0]
	}

	log.Printf("[main] Query parcel of order=[%s], current cache dir:[%s]  [%s]", orderNumber, wf.Cache.Dir)

	// check the input ip is valid
	if orderNumber != "" && len(orderNumber) != 18 {
		log.Printf("Entering order number ....")
		wf.NewItem(fmt.Sprintf("Entering order number %s ...", orderNumber)).Valid(true).Subtitle(orderNumber)
		return wf
	}

	var jsonResult *ParcelResult

	orderNumberList := []string{}
	if orderNumber == "" {
		if wf.Cache.Expired(orderCacheName, orderMaxCacheAge) {
			wf.ClearCache()
		}

		log.Print("List order numbers in cache")

		if wf.Cache.Exists(orderCacheName) {
			if err := wf.Cache.LoadJSON(orderCacheName, &orderNumberList); err != nil {
				wf.FatalError(err)
			} else {
				for _, orderNumber := range orderNumberList {
					title := fmt.Sprintf("%s", orderNumber)
					wf.NewItem(title).Valid(false).Arg(title).Subtitle(title).Autocomplete(title)
				}
			}
		}
		return wf
	} else {
		if orderNumber != "" { // search by order number
			if jsonResult == nil {
				// search online
				getParcelInfo(orderNumber, &jsonResult)
				log.Print("Search online with result : ", jsonResult)
			}
		}

		if jsonResult == nil || jsonResult.Status == 0 {
			title := fmt.Sprintf(jsonResult.Info)
			wf.NewItem(title).Valid(false).Arg(title).Subtitle(orderNumber)
		} else {
			// append traces
			traces := jsonResult.Data.Shipment.Traces
			for _, trace := range traces {
				title := fmt.Sprintf("%s", trace.Info)
				wf.NewItem(title).Valid(false).Arg(title).Subtitle(fmt.Sprintf("%s", trace.Time))
			}

			// Fetch success
			if wf.Cache.Exists(orderCacheName) {
				wf.Cache.LoadJSON(orderCacheName, &orderNumberList)
			}

			existOrderNumberList := mapset.NewSet()
			if len(orderNumberList) > 0 {
				for _, orderNumberTemp := range orderNumberList {
					existOrderNumberList.Add(orderNumberTemp)
				}
			}
			existOrderNumberList.Add(orderNumber)

			wf.Cache.StoreJSON(orderCacheName, existOrderNumberList)
		}
	}

	// Send results to Alfred
	return wf
}

func init() {
	RegisterCommand(&QueryParcelCommand{})
}

func getParcelInfo(orderNumber string, jsonResult interface{}) {
	url := fmt.Sprintf("http://xingtx.sah.nextsls.com/tracking/app?inajax=1&tracking_number=%s", orderNumber)
	log.Printf(orderNumber)
	common.HttpGetJSON(url, &jsonResult)
}
