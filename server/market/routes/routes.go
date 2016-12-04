package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/moryg/eve_analyst/server/market/routes/overpriced"
)

func Load(r *httprouter.Router) *httprouter.Router {
	r.GET("/api/market/overpriced", overpriced.Get)
	return r
}
