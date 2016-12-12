package control

import (
	"github.com/moryg/eve_analyst/apiqueue/requests/market/regionfull"
)

func BootUp() {
	MarketFull()
}

func MarketFull() {
	// return
	// TMP - do not update other regions for now
	regionfull.Update(10000002) // The Forge
	// regionfull.Update(10000047) // Providence
	// regionfull.Update(10000043) // Domain
	// regionfull.Update(10000001) // Derelik
	// regionfull.Update(10000011)
	// regionfull.Update(10000012)
	// regionfull.Update(10000015)
	// regionfull.Update(10000020)
	// regionfull.Update(10000030)
	// regionfull.Update(10000032)
	// regionfull.Update(10000033)
	// regionfull.Update(10000044)
}
