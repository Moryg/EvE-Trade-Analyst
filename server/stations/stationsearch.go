package stations

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/moryg/eve_analyst/database/station"
	"github.com/moryg/eve_analyst/shared"
	// "log"
	"net/http"
)

type searchRsp struct {
	Stations []shared.Station `json:"stations"`
}

func search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	query := params.ByName("query")
	if len(query) < 3 {
		shared.SendError(w, "Query too short", http.StatusNotAcceptable)
		return
	}

	rsp := new(searchRsp)
	rsp.Stations = station.FindByName(query)

	json.NewEncoder(w).Encode(rsp)
}
