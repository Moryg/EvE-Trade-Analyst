package concatenator_test

import (
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"testing"
)

func TestGetAvg(test *testing.T) {
	p := concatenator.NewPrice(3.5, 20)

	if p.Mean != 3.5 {
		test.Log("Initial average failed")
		test.FailNow()
	}

	p.Add(4.5, 20)

	if p.Mean != 4 {
		test.Log("First addition failed")
		test.FailNow()
	}

	p.Add(3, 40)

	if p.Mean != 3.5 {
		test.Log("Second addition failed")
		test.FailNow()
	}
}
