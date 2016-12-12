package concatenator_test

import (
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"strings"
	"testing"
)

func TestRegionQuery(test *testing.T) {
	region := concatenator.NewRegion(1)

	result, err := region.ConstructSQL()
	if err == nil {
		test.Error("Expected error but none received")
	}

	region.RegionID = 3

	result, err = region.ConstructSQL()
	if err == nil {
		test.Error("Expected error but none received")
	}

	region.Add(100, 100, 1, 2)
	region.Add(60, 100, 1, 2)
	region.Add(200, 100, 1, 2)

	// (`stationId`, `itemId`, `regionId`, `min`, `mean`, `max`, `upFlag`)
	expected := "(1, 2, 3, 60.00, 120.00, 200.00, 1)"
	result, err = region.ConstructSQL()
	if err != nil {
		test.Error(err)
	}
	if result != expected {
		test.Errorf("Value '%s' does not match expected '%s'", result, expected)
	}

	region.Add(1, 1, 1, 5)
	result, err = region.ConstructSQL()
	if err != nil {
		test.Error(err)
	}
	if !strings.Contains(result, "(1, 2, 3, 60.00, 120.00, 200.00, 1)") ||
		!strings.Contains(result, "(1, 5, 3, 1.00, 1.00, 1.00, 1)") ||
		strings.Count(result, ",") != 13 {
		test.Errorf("Incorrect value: '%s'", result)
	}

	region.Add(50.656, 50, 8, 2)
	result, err = region.ConstructSQL()
	if err != nil {
		test.Error(err)
	}
	if !strings.Contains(result, "(1, 2, 3, 60.00, 120.00, 200.00, 1)") ||
		!strings.Contains(result, "(1, 5, 3, 1.00, 1.00, 1.00, 1)") ||
		!strings.Contains(result, "(8, 2, 3, 50.66, 50.66, 50.66, 1)") ||
		strings.Count(result, ",") != 20 {
		test.Errorf("Incorrect value: '%s'", result)
	}
}
