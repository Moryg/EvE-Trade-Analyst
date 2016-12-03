package station

import (
	"encoding/json"
	// . "github.com/moryg/eve_analyst/apiqueue/requests"
	"net/http"
)

type rspJSON struct {
	SystemID int    `json:"solar_system_id"`
	Name     string `json:"station_name"`
}

func parseResBody(res *http.Response) (*rspJSON, error) {
	var resJson rspJSON
	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&resJson)
	if err != nil {
		return nil, err
	}

	return &resJson, nil
}
