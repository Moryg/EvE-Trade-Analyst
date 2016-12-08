package concatenator_test

import (
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"testing"
)

func TestPriceQuery(test *testing.T) {
	price := concatenator.NewPrice(100, 100)
	price.Add(60, 100)
	price.Add(200, 100)

	q := price.String()
	expected := "60.00, 120.00, 200.00"
	if q != expected {
		test.Logf("Value '%s' does not match expected '%s'", q, expected)
		test.Fail()
	}

	price = concatenator.NewPrice(10, 1)
	price.Add(5, 2)
	expected = "5.00, 6.67, 10.00"
	q = price.String()
	if q != expected {
		test.Logf("Value '%s' does not match expected '%s'", q, expected)
		test.Fail()
	}

	price = concatenator.NewPrice(5, 1)
	price.Add(2.5, 2)
	expected = "2.50, 3.33, 5.00"
	q = price.String()
	if q != expected {
		test.Logf("Value '%s' does not match expected '%s'", q, expected)
		test.Fail()
	}
}
