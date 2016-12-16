package control

import (
	"github.com/moryg/eve_analyst/apiqueue/requests/market/regionfull"
	stationApi "github.com/moryg/eve_analyst/apiqueue/requests/station"
	"github.com/moryg/eve_analyst/database/region"
	"github.com/moryg/eve_analyst/database/station"
	"log"
	"time"
)

var (
	regionUpdater *time.Ticker
)

func BootUp() {
	if regionUpdater != nil {
		log.Fatal("Ticker is not nil")
	}

	regionUpdater = time.NewTicker(time.Second * 30)

	go func() {
		for {
			<-regionUpdater.C
			go MarketBatch()
			go StationBatch()
		}
	}()

	go MarketBatch()
	go StationBatch()
}

func StationBatch() {
	missing := station.GetMissingStations()
	for _, id := range missing {
		stationApi.Update(id)
	}
}

func MarketBatch() {

	res := region.GetUpdatableRegions()
	for _, id := range res {
		log.Printf("Starting update of region %d", id)
		regionfull.Update(id)
	}
	// TMP - do not update other regions for now
	return
	regionfull.Update(10000047) // Providence
	// regionfull.Update(10000002) // The Forge
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
