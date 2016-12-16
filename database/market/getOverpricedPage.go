package market

import (
	"github.com/jmoiron/sqlx"
	. "github.com/moryg/eve_analyst/database"
	. "github.com/moryg/eve_analyst/shared"
	"log"
)

var getOverPriceStmt *sqlx.Stmt

func init() {
	var err error
	getOverPriceStmt, err = DB.Preparex("SELECT `buy`.`itemId`, `item`.`name`, `buy`.`min` `bMin`, `buy`.`mean` `bMean`, `sell`.`min` `sMin`, `sell`.`mean` `sMean`, CAST(`sell`.`min` / `buy`.`min` AS DECIMAL(20,2)) `rMin`, CAST(`sell`.`mean` / `buy`.`mean` AS DECIMAL(20,2)) `rMean` FROM `orderSell` `buy` JOIN `orderSell` `sell` ON `sell`.`stationId` = ? AND `sell`.`itemId` = `buy`.`itemId` JOIN `item` ON `item`.`id` = `buy`.`itemId` WHERE `buy`.`stationID` = ? ORDER BY (IF(? = 'mean', `rMean`, `rMin`)) DESC LIMIT ?, ?;")

	if err != nil {
		log.Fatal("Failed preparing getOverPriced query statement")
	}
}

func GetOverPricedPage(buyId, sellId int64, page int, sortCol string) []CompItem {
	var items []CompItem

	perPage := 100
	page -= 1 // 0 Based paging in query
	if page < 0 {
		page = 0
	}
	page = page * perPage
	err := getOverPriceStmt.Select(&items, sellId, buyId, sortCol, page, perPage)
	if err != nil {
		log.Printf("getOverpriced query: " + err.Error())
	}
	return items
}
