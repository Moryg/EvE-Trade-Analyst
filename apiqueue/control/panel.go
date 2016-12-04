package control

import (
	"github.com/moryg/eve_analyst/apiqueue/requests/market/regionfull"
)

func BootUp() {
	MarketFull()
}

func MarketFull() {
	// TMP - do not update other regions for now
	// regionfull.Update(10000002) // The Forge
	// regionfull.Update(10000047) // Providence
	regionfull.Update(10000043)
}
