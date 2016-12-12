package concatenator_test

import (
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"testing"
)

func Test_RegionMerge(test *testing.T) {
	regA := concatenator.NewRegion(1)
	regA.Add(100, 100, 1, 1)
	regA.Add(100, 100, 2, 1)
	regB := concatenator.NewRegion(1)
	regB.Add(300, 100, 1, 1)

	regA.Merge(regB)

	pr, ok := regA.Prices[1][1]
	if !ok {
		test.Log("No item 1-1 after merge")
		test.FailNow()
	}

	if pr.Mean != 200 {
		test.Log("Incorrect mean on 1-1 after merge")
		test.FailNow()
	}
}

func Test_AsyncMerge(test *testing.T) {

}
