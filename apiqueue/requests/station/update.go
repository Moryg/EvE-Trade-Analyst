package station

import (
	"fmt"
	"github.com/moryg/eve_analyst/apiqueue"
	// "github.com/moryg/eve_analyst/apiqueue/updatecache"
	// "time"
)

const (
	url string = "https://esi.tech.ccp.is/latest/universe/stations/%d/"
)

// var cache = updatecache.NewInt(time.Hour * 24)

type request struct {
	stationId uint64
	url       string
}

func Update(id uint64) {
	// if !cache.Add(id) {
	// 	// This station has been recently updated!
	// 	return
	// }
	u := request{}
	u.stationId = id
	u.url = fmt.Sprintf(url, id)
	apiqueue.Enqueue(u)
}

func (r request) Execute() {
	(&r).execute()
}

func (r request) RequiresAuth() bool {
	return false
}
