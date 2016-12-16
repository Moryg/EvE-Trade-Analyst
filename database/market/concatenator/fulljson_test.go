package concatenator_test

import (
	"encoding/json"
	"fmt"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	// "github.com/moryg/eve_analyst/shared"
	"io/ioutil"
	// "os"
	"strconv"
	"testing"
)

type deconJson struct {
	Items concatenator.Region `json:"items"`
	// Items     []shared.Order `json:"items"`
	// PageCount int `json:"pageCount"`
}

func TestJsonDecompile(test *testing.T) {
	jsonFiles := []string{"./full_sample.json", "./full_sample_pg2.json"}
	masterRegion := concatenator.NewRegion()
	src, err := ioutil.ReadFile("./full_stats.json")
	if err != nil {
		test.Fatal(err)
	}

	var (
		sta, ite string
		holder   deconJson
		static   map[string]map[string][2]float64
	)

	err = json.Unmarshal(src, &static)
	if err != nil {
		test.Fatal(err)
	}

	for _, file := range jsonFiles {
		raw, err := ioutil.ReadFile(file)
		if err != nil {
			test.Fatal(err)
		}

		err = json.Unmarshal(raw, &holder)
		if err != nil {
			test.Fatal(err)
		}

		masterRegion.Merge(&holder.Items)
	}

	if len(masterRegion.Prices) != len(static) {
		test.Log("Number of stations does not match")
		test.FailNow()
	}

	for stationId, items := range masterRegion.Prices {
		sta = strconv.FormatUint(stationId, 10)

		if len(items) != len(static[sta]) {
			test.Log("Number of stations does not match")
			test.FailNow()
		}
		for itemId, price := range items {
			ite = strconv.FormatUint(itemId, 10)
			val, ok := static[sta][ite]
			if !ok {
				test.Logf("Station %d item %d not found", stationId, itemId)
				test.Fail()
			}

			if fmt.Sprintf("%.2f", price.Min) != fmt.Sprintf("%.2f", val[0]) {
				test.Logf("Station %d item %d minimum mismatch", stationId, itemId)
				test.Fail()
			}

			if fmt.Sprintf("%.2f", price.Mean) != fmt.Sprintf("%.2f", val[1]) {
				test.Logf("Station %d item %d mean mismatch", stationId, itemId)
				test.Fail()
			}
		}
	}
}
