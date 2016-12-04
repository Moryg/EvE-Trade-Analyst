package routes

import (
	"github.com/julienschmidt/httprouter"
)

func Load(r *httprouter.Router) *httprouter.Router {
	r.GET("/api/market/overpriced", getOverpriced)
	return r
}
