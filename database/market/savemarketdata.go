package market

import (
	. "github.com/moryg/eve_analyst/database"
	. "github.com/moryg/eve_analyst/shared"
	"log"
	"strings"
)

func CleanMarketRegion(regionID int, batchID string) {
	_, err := DB.Exec("UPDATE `orderSell` SET `deletedAt` = NOW() WHERE `regionId` = ? AND `batchId` <> ?", regionID, batchID)
	if err != nil {
		log.Println("savemarketdata.clean: " + err.Error())
	}
}

func SaveMarketData(batchID string, regionID int, c chan bool, orders []Order) {
	var err error
	query := "INSERT INTO `orderSell` (`id`, `itemId`, `stationId`, `regionID`, `price`, `volume`, `expiry`, `batchId`, `deletedAt`) VALUES "
	inserts := []interface{}{}
	row := "(?,?,?,?,?,?,?,?,NULL),"

	for _, order := range orders {
		if order.Buy {
			// ignore buy orders for now
			continue
		}
		query += row
		inserts = append(inserts, order.Id, order.ItemID, order.StationID, regionID, order.Price, order.Volume, order.ExpiryDate(), batchID)
	}

	query = strings.TrimSuffix(query, ",")
	query += "ON DUPLICATE KEY UPDATE `itemId`=VALUES(`itemId`),`stationId`=VALUES(`stationId`),`price`=VALUES(`price`),`volume`=VALUES(`volume`),`expiry`=VALUES(`expiry`),`batchId`=VALUES(`batchId`),`regionID`=VALUES(`regionID`);"
	stmt, err := DB.Prepare(query)

	if err == nil {
		_, err = stmt.Exec(inserts...)
	}

	if err != nil {
		log.Println(err)
	}

	c <- true
}
