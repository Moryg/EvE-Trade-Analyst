package overpriced

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/moryg/eve_analyst/database/market"
	"github.com/moryg/eve_analyst/shared"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	buyId, err := strconv.ParseInt(params.ByName("buyId"), 10, 64)
	if err != nil || buyId == 0 {
		shared.SendError(w, "Invalid buy station ID", 403)
		return
	}

	sellId, err := strconv.ParseInt(params.ByName("sellId"), 10, 64)
	if err != nil || sellId == 0 {
		shared.SendError(w, "Invalid selling station ID", 403)
		return
	}

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page == 1 {
		page = 1
	}

	sort := query.Get("sort")
	if sort == "" || sort != "min" && sort != "mean" {
		sort = "min"
	}

	data := shared.CompHolder{}
	data.Items = market.GetOverPricedPage(buyId, sellId, page, sort)
	buf, err := (&data).MarshalJSON()

	if err != nil {
		log.Printf("overpriced.GET json: %d %d %d - "+err.Error(), buyId, sellId, page)
		shared.SendError(w, "Failed encoding json", 501)
	}
	fmt.Fprintln(w, string(buf))
}
