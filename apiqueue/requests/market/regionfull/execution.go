package regionfull

import (
	"fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	"github.com/moryg/eve_analyst/apiqueue/ratelimit"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"log"
	"net/http"
)

func (r *Request) execute() {
	// Prepare and execute the request
	r.url = fmt.Sprintf(baseUrl, r.regionID, r.page)

	ratelimit.Add()
	res, err := http.Get(r.url)
	ratelimit.Sub()

	// Request error
	if err != nil {
		log.Println("regionfull.execute: " + err.Error())
		r.requestComplete(err)
		return
	}

	// Parse rsp json to structure
	orders, err := parseResBody(res)
	if err != nil {
		log.Println("regionfull.execute parse:" + err.Error())
		r.requestComplete(err)
		return
	}

	if r.Statistics == nil {
		r.Statistics = concatenator.NewRegion(r.regionID)
	}

	for _, order := range orders {
		if order.Buy {
			continue
		}
		r.RawOrders = append(
			r.RawOrders,
			fmt.Sprintf("(%d,%d,%d,%.2f,%d)", order.Id, order.StationId, order.ItemId, order.Price, order.Volume),
		)
		r.Statistics.Add(order.Price, order.Volume, order.StationId, order.ItemId)
	}

	log.Printf("Done with page %d of region %d", r.page, r.regionID)

	if len(orders) < 10000 {
		r.requestComplete(nil)
		return
	}

	apiqueue.Enqueue(r.newPage())
}
