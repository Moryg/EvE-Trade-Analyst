package market

import (
	. "github.com/moryg/eve_analyst/database"
	. "github.com/moryg/eve_analyst/shared"
	"log"
)

var sql_GetOverPricedPage string

func init() {
	sql_GetOverPricedPage = "SELECT `item`.`id`, `item`.`groupID`, `item`.`volume`, `item`.`name`," +
		"`sell`.`min` `sellMin`, `sell`.`mean` `sellMean`," +
		"`buy`.`min` `buyMin`, `buy`.`mean` `buyMean`," +
		"CAST((`sell`.`min` / `buy`.`min`)AS DECIMAL(20,2)) `minRatio`," +
		"CAST((`sell`.`mean` / `buy`.`mean`)AS DECIMAL(20,2)) `meanRatio` " +
		"FROM `item` " +
		// Join BUY subselect
		"JOIN (SELECT `itemId`," +
		"MIN(`price`) `min`," +
		"CAST(SUM(`price` * `volume`)/SUM(`volume`) AS DECIMAL(20,2)) `mean` " +
		"FROM `orderSell` " +
		"WHERE `stationID` = ? " + // Buy station ID
		"GROUP BY `itemId`" +
		") `buy` ON `buy`.`itemId` = `item`.`id` " +
		// Join SELL subselect
		"JOIN (SELECT `itemId`," +
		"MIN(`price`) `min`," +
		"CAST(SUM(`price` * `volume`)/SUM(`volume`) AS DECIMAL(20,2)) `mean` " +
		"FROM `orderSell` " +
		"WHERE `stationID` = ? " + // Sell station ID
		"GROUP BY `itemId`" +
		") `sell` ON `sell`.`itemId` = `item`.`id` " +
		"ORDER BY `meanRatio` DESC " +
		"LIMIT ?, ?;" // Page values
}

func GetOverPricedPage(buyId, sellId int64, page int) []CompItem {
	var items []CompItem

	perPage := 100
	page -= 1 // 0 Based paging in query
	if page < 0 {
		page = 0
	}
	page = page * perPage
	err := DB.Select(&items, sql_GetOverPricedPage, buyId, sellId, page, perPage)
	if err != nil {
		log.Printf("getOverpriced query: " + err.Error())
	}
	return items
}
