package regionfull

import (
	// "fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	// "time"
	db "github.com/moryg/eve_analyst/database/market"
	"log"
	"time"
)

const (
	// baseUrl string = "https://crest-tq.eveonline.com/market/%d/orders/all/?page=%d" // CREST
	baseUrl string = "https://esi.tech.ccp.is/latest/markets/%d/orders/?page=%d" // ESI
)

type Request struct {
	regionID   uint64
	page       int
	url        string
	uid        string
	Statistics *concatenator.Region
	RawOrders  []string
}

func Update(id uint64) {
	r := create(id)
	apiqueue.Enqueue(r)
}

func create(regionID uint64) Request {
	r := Request{}
	r.regionID = regionID
	r.page = 1
	return r
}

func (src *Request) newPage() Request {
	r := create(src.regionID)
	r.page = src.page + 1
	r.Statistics = src.Statistics
	r.RawOrders = src.RawOrders
	return r
}

func (r Request) Execute() {
	(&r).execute()
}

func (this *Request) requestComplete(err error) {
	log.Printf("Done with region %d", this.regionID)
	if err != nil {
		log.Printf("regionfull.complete: region %d, error:%s\n", this.regionID, err.Error())
	}

	t0 := time.Now()
	err = db.SaveRegionStatistics(this.Statistics, this.regionID)
	log.Printf("Saving region %d to DB took %.3fs", this.regionID, time.Now().Sub(t0).Seconds())
	if err != nil {
		log.Printf("regionfull.complete: region %d, error:%s\n", this.regionID, err.Error())
	}
}

func (r Request) RequiresAuth() bool {
	return false
}
