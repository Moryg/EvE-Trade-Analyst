package stations

import (
	"github.com/julienschmidt/httprouter"
)

func Load(r *httprouter.Router) *httprouter.Router {
	r.GET("/api/stations/search/:query", search)
	return r
}
