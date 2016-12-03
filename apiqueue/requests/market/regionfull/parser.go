package regionfull

import (
	"encoding/json"
	. "github.com/moryg/eve_analyst/apiqueue/requests"
	"net/http"
)

type rgJson struct {
	Items     []Order `json:"items"`
	ItemCount int     `json:"totalCount"`
	PageCount int     `json:"pageCount"`
}

func parseResBody(res *http.Response) (*rgJson, error) {
	var resJson rgJson
	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&resJson)
	if err != nil {
		return nil, err
	}

	return &resJson, nil
}
