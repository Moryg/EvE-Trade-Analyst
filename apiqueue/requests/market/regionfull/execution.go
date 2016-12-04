package regionfull

import (
	"fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	"github.com/moryg/eve_analyst/apiqueue/ratelimit"
	"github.com/moryg/eve_analyst/database/market"
	"log"
	"math"
	"net/http"
	"os"
)

func (r *Request) execute() {
	// Prepare and execute the request
	r.url = fmt.Sprintf(baseUrl, os.Getenv("API"), r.regionID, r.page)
	ratelimit.Add()
	res, err := http.Get(r.url)
	ratelimit.Sub()

	// Request error
	if err != nil {
		log.Println("regionfull.execute: " + err.Error())
		r.requestComplete()
		return
	}

	// Parse rsp json to structure
	rsp, err := parseResBody(res)
	if err != nil {
		log.Println("regionfull.execute parse:" + err.Error())
		r.requestComplete()
		return
	}

	// If this is the first page, start a channel
	if r.requestBatch == nil {
		r.requestBatch = make(chan bool, rsp.PageCount)
	}

	// Enqueue all other pages if there are any
	if r.page == 1 && rsp.PageCount > 1 {
		for ii := 2; ii <= rsp.PageCount; ii++ {
			apiqueue.Enqueue(r.newPage(ii))
		}
	}

	// Chunk and save the data
	chunk := 2000
	items := len(rsp.Items)
	splits := int(math.Ceil(float64(items) / float64(chunk)))
	innerChannel := make(chan bool, splits)
	for start, end := 0, chunk; start < items-1; start, end = start+chunk, end+chunk {
		if end > items {
			end = items
		}

		go market.SaveMarketData(r.uid, r.regionID, innerChannel, rsp.Items[start:end])
	}

	// Wait for all chunks  of this request to finish
	for splits > 0 {
		<-innerChannel
		splits--
	}

	// Request data handling done if not first page
	r.requestComplete()
	if r.page != 1 {
		return
	}

	// Wait for all pages to finish
	items = rsp.PageCount
	for items > 0 {
		<-r.requestBatch
		items--
		log.Printf("Region %d - %d pages remaining", r.regionID, items)
	}

	// Clean up duty
	market.CleanMarketRegion(r.regionID, r.uid)
}
