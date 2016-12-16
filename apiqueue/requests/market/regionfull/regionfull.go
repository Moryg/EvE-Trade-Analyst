package regionfull

import (
	"fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"time"
)

const (
	baseUrl string = "%s/market/%d/orders/all/?page=%d"
)

type Request struct {
	regionID     uint64
	page         int
	url          string
	uid          string
	requestBatch chan *concatenator.Region
}

func Update(id uint64) {
	r := create(id)
	apiqueue.Enqueue(r)
}

func create(regionID uint64) Request {
	r := Request{}
	r.regionID = regionID
	r.page = 1
	r.uid = fmt.Sprintf("rugfull_%d_%d", regionID, time.Now().Unix())
	return r
}

func (src *Request) newPage(page int) Request {
	r := create(src.regionID)
	r.page = page
	r.uid = src.uid
	r.requestBatch = src.requestBatch
	return r
}

func (r Request) Execute() {
	(&r).execute()
}

func (r *Request) requestComplete(data *concatenator.Region) {
	if r.requestBatch != nil {
		r.requestBatch <- data
	}
}

func (r Request) RequiresAuth() bool {
	return false
}
