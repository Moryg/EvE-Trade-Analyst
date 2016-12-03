package station

import (
	"github.com/moryg/eve_analyst/apiqueue/ratelimit"
	"github.com/moryg/eve_analyst/database/station"
	"log"
	"net/http"
)

func (r *request) execute() {
	ratelimit.Add()
	res, err := http.Get(r.url)
	ratelimit.Sub()

	if err != nil {
		log.Println("station.execute: " + err.Error())
		return
	}

	rsp, err := parseResBody(res)
	if err != nil {
		log.Println("station.execute parse:" + err.Error())
		return
	}

	station.Update(r.stationId, rsp.SystemID, rsp.Name)
	log.Printf("Updated station %d\n", r.stationId)
}
