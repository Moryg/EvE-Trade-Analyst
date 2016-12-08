package market

import (
	. "github.com/moryg/eve_analyst/database"
	"github.com/moryg/eve_analyst/database/market/concatenator"
	"log"
)

func CleanMarketRegion(regionID int, batchID string) {
	_, err := DB.Exec("DELETE FROM `orderSell` WHERE `regionId` = ? AND `batchId` <> ?", regionID, batchID)
	if err != nil {
		log.Println("savemarketdata.clean: " + err.Error())
		return
	}

	_, err = DB.Exec("CALL `calculateStationData`(?);", regionID)
	if err != nil {
		log.Println("savemarketdata.clean: " + err.Error())
		return
	}
}

func SaveMarketData(data *concatenator.Region, regionId int) error {
	sqlBase := "INSERT INTO `orderSell` (`stationId`, `itemId`, `regionId`, `min`, `mean`, `max`, `upFlag`) VALUES "
	sqlEnd := " ON DUPLICATE KEY UPDATE `min` = VALUES(`min`), `max` = VALUES(`max`), `mean` = VALUES(`mean`), `upFlag` = 1;"

	sql, err := data.ConstructSQL()
	if err != nil {
		return err
	}

	sql = sqlBase + sql + sqlEnd
	_, err = DB.Exec(sql)

	if err != nil {
		return err
	}

	cleanUp1 := "DELETE FROM `orderSell` WHERE `regionId` = ? AND `upFlag` = 0;"
	cleanUp2 := "UPDATE `orderSell` SET `upFlag` = 0 WHERE `regionId` = ?;"

	_, _ = DB.Exec(cleanUp1, regionId)
	_, err = DB.Exec(cleanUp2, regionId)
	if err != nil {
		return err
	}

	return nil
}
