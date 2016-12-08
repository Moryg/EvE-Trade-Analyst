package regionfull

import (
	"fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	"github.com/moryg/eve_analyst/apiqueue/ratelimit"
	db "github.com/moryg/eve_analyst/database/market"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"log"
	"net/http"
	"os"
)

func (r *Request) execute() {
	// Prepare and execute the request
	r.url = fmt.Sprintf(baseUrl, os.Getenv("API"), r.regionID, r.page)
	// r.url = "http://localhost:8888/market/1/orders/all/"
	ratelimit.Add()
	res, err := http.Get(r.url)
	ratelimit.Sub()

	// Request error
	if err != nil {
		log.Println("regionfull.execute: " + err.Error())
		r.requestComplete(concatenator.NewRegion())
		return
	}

	// Parse rsp json to structure
	rsp, err := parseResBody(res)
	if err != nil {
		log.Println("regionfull.execute parse:" + err.Error())
		r.requestComplete(concatenator.NewRegion())
		return
	}

	// If this is the first page, start a channel
	if r.requestBatch == nil {
		r.requestBatch = make(chan *concatenator.Region, rsp.PageCount)
	}

	// Enqueue all other pages if there are any
	if r.page == 1 && rsp.PageCount > 1 {
		for ii := 2; ii <= rsp.PageCount; ii++ {
			apiqueue.Enqueue(r.newPage(ii))
		}
	}

	if r.page != 1 {
		r.requestComplete(&rsp.Items)
		return
	}

	mainData := rsp.Items
	mainData.RegionID = r.regionID

	items := rsp.PageCount - 1
	var subData *concatenator.Region
	for items > 0 {
		subData = <-r.requestBatch
		_, _ = subData.Prices[1][1]
		items--
		log.Printf("Region %d - %d pages remaining", r.regionID, items)
	}

	err = db.SaveMarketData(&mainData, r.regionID)
	if err != nil {
		log.Println(err)
	}
	return
}
