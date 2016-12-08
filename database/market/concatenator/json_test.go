package concatenator_test

import (
	"bytes"
	"encoding/json"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"testing"
)

type rgJson struct {
	Items *concatenator.Region `json:"items"`
}

func TestParseJson(test *testing.T) {
	jsonString := "{\"items\": [{\"buy\": false,\"issued\": \"2016-09-06T11:15:05\",\"price\": 400.0,\"volume\": 200,\"duration\": 365,\"id\": 911203054,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 1,\"type\": 1},{\"buy\": false,\"issued\": \"2016-09-28T11:04:37\",\"price\": 500.0,\"volume\": 300,\"duration\": 365,\"id\": 911203055,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 1,\"type\": 1},{\"buy\": false,\"issued\": \"2016-10-06T11:06:32\",\"price\": 5000.0,\"volume\": 300,\"duration\": 365,\"id\": 911203056,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 2,\"type\": 1},{\"buy\": false,\"issued\": \"2016-11-15T11:59:13\",\"price\": 7000.0,\"volume\": 300,\"duration\": 365,\"id\": 911203057,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 2,\"type\": 1},{\"buy\": false,\"issued\": \"2016-10-25T11:08:38\",\"price\": 50000.0,\"volume\": 432,\"duration\": 365,\"id\": 911203058,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 2,\"type\": 5},{\"buy\": false,\"issued\": \"2016-10-18T11:23:34\",\"price\": 12341.0,\"volume\": 1,\"duration\": 365,\"id\": 911203059,\"minVolume\": 1,\"volumeEntered\": 1,\"range\": \"region\",\"stationID\": 8,\"type\": 1},{\"buy\": true,\"issued\": \"2016-12-02T11:21:22\",\"price\": 10000,\"volume\": 1000,\"duration\": 365,\"id\": 911203095,\"minVolume\": 1,\"volumeEntered\": 13606,\"range\": \"station\",\"stationID\": 1,\"type\": 1},{\"buy\": true,\"issued\": \"2016-12-02T18:04:32\",\"price\": 18000,\"volume\": 200,\"duration\": 365,\"id\": 911203096,\"minVolume\": 1,\"volumeEntered\": 13606,\"range\": \"station\",\"stationID\": 2,\"type\": 1}],\"totalCount\": 273534,\"next\": {\"href\": \"https://crest-tq.eveonline.com/market/10000002/orders/all/?page=2\"},\"pageCount\": 10}"

	r := new(rgJson)
	err := json.NewDecoder(bytes.NewBufferString(jsonString)).Decode(r)
	if err != nil {
		test.Log(err)
		test.FailNow()
	}

	chunk, ok := r.Items.Prices[1][1]
	if !ok {
		test.Log("Missing 1-1 value")
		test.FailNow()
	} else if chunk.Mean != float64(460) {
		test.Log("Incorrect value 1-1")
		test.FailNow()
	}

	chunk, ok = r.Items.Prices[2][1]
	if !ok {
		test.Log("Missing 2-1 value")
		test.FailNow()
	} else if chunk.Mean != float64(6000) {
		test.Log("Incorrect value 2-1")
		test.FailNow()
	}

	chunk, ok = r.Items.Prices[66][23]
	if ok {
		test.Log("Unexpected value")
		test.FailNow()
	}
}
