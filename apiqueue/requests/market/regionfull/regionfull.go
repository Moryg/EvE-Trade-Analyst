package regionfull

import (
	"fmt"
	"time"
)

const (
	baseUrl string = "%s/market/%d/orders/all/?page=%d"
)

type Request struct {
	regionID     int
	page         int
	url          string
	uid          string
	requestBatch chan bool
}

func New(regionID int) Request {
	r := Request{}
	r.regionID = regionID
	r.page = 1
	r.uid = fmt.Sprintf("rugfull_%d_%d", regionID, time.Now().Unix())
	return r
}

func (src *Request) newPage(page int) Request {
	r := New(src.regionID)
	r.page = page
	r.uid = src.uid
	r.requestBatch = src.requestBatch
	return r
}

func (r Request) Execute() {
	(&r).execute()
}

func (r *Request) requestComplete() {
	if r.requestBatch != nil {
		r.requestBatch <- false
	}
}

func (r Request) RequiresAuth() bool {
	return false
}
