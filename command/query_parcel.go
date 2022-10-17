package command

// Package is called aw
import (
	"alfred-utils-go/common"
	"fmt"
	aw "github.com/deanishe/awgo"
	"log"
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

func (c *QueryParcelCommand) Description() Description {
	return Description{
		Short: "Query the parcel info to Sweden.",
		Usage: []string{
			"parcel order-number",
		},
	}
}

func (c *QueryParcelCommand) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) <= 0 {
		return wf
	}

	var orderNumber string
	if len(args) >= 1 { // has order number arg
		orderNumber = args[0]
	}

	log.Printf("Query parcel of order [%s]", orderNumber)

	var jsonResult *ParcelResult
	results := []*ParcelResult{}

	if orderNumber != "" { // search own ip location
		if jsonResult == nil { // cannot find in local cache
			// search online
			getParcelInfo(orderNumber, &jsonResult)
			log.Print("Search online with result : ", jsonResult)
			results = append(results, jsonResult)
		}
	}

	if jsonResult == nil || jsonResult.Status == 0 {
		title := fmt.Sprintf(jsonResult.Info)
		wf.NewItem(title).Valid(true).Arg(title).Subtitle(orderNumber)
	} else {
		// append shipment info

		// append traces
		traces := jsonResult.Data.Shipment.Traces
		for _, trace := range traces {
			title := fmt.Sprintf("%s", trace.Info)
			wf.NewItem(title).Valid(false).Arg(title).Subtitle(fmt.Sprintf("%s", trace.Time))
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
