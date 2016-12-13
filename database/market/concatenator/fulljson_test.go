package concatenator_test

import (
	"encoding/json"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"io/ioutil"
	"testing"
)

type deconJson struct {
	Items     concatenator.Region `json:"items"`
	ItemCount int                 `json:"totalCount"`
	PageCount int                 `json:"pageCount"`
}

func TestFullJson(test *testing.T) {
	// TODO - implement proper test
	test.SkipNow()
	raw, err := ioutil.ReadFile("./full_sample.json")
	if err != nil {
		test.Fatal(err)
	}
	raw2, err := ioutil.ReadFile("./full_sample_pg2.json")
	if err != nil {
		test.Fatal(err)
	}

	var (
		res, res2 deconJson
	)

	err = json.Unmarshal(raw, &res)
	if err != nil {
		test.Fatal(err)
	}

	err = json.Unmarshal(raw2, &res2)
	if err != nil {
		test.Fatal(err)
	}

	if res.PageCount != 2 {
		test.Fatal("Page count does not match")
	}

	region := res.Items

	region.Merge(&res2.Items)

	for stationID, station := range region.Prices {
		for itemID, price := range station {
			if price.Min < 0.01 {
				test.Fatalf("Invalid minimum for item %d in station %d", itemID, stationID)
			}
		}
	}
}
