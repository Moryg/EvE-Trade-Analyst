package regionfull

import (
	"encoding/json"
	// "github.com/moryg/eve_analyst/database/market/concatenator"
	. "github.com/moryg/eve_analyst/shared"
	"net/http"
)

type item struct {
}

type rgJson struct {
	Items     []Order `json:"items"`
	ItemCount int     `json:"totalCount"`
	PageCount int     `json:"pageCount"`
}

func parseResBody(res *http.Response) ([]Order, error) {
	items := []Order{}
	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
