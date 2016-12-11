package station

import (
	// "database/sql"
	"github.com/jmoiron/sqlx"
	. "github.com/moryg/eve_analyst/database"
	"github.com/moryg/eve_analyst/shared"
	"log"
)

var stmtSearch *sqlx.Stmt

func init() {
	var err error
	stmtSearch, err = DB.Preparex("SELECT `station`.`id`, `station`.`name`, `region`.`name` `region`, `system`.`name` `system` FROM `station` JOIN `region` ON `region`.`id` = `station`.`regionId` JOIN `system` ON `system`.`id` = `station`.`systemId` WHERE `station`.`name` LIKE ? ORDER BY `station`.`sortName` ASC LIMIT 5;")
	if err != nil {
		log.Fatal(err)
	}
}

func FindByName(name string) []shared.Station {
	res := []shared.Station{}
	err := stmtSearch.Select(&res, "%"+name+"%")
	if err != nil {
		log.Println(err)
		return []shared.Station{}
	}

	return res
}
