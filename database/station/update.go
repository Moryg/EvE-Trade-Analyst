/**
 * Update or insert a station
 */
package station

import (
	. "github.com/moryg/eve_analyst/database"
	"log"
)

func Update(id, systemID int, name string) {
	sql := "insert into `station` (`id`, `systemID`, `regionID`, `name`)" +
		"SELECT ? AS `id`, ? as `systemId`, regionID, ? as `name`" +
		" FROM `system` WHERE `system`.`id` = ?" +
		" ON DUPLICATE KEY UPDATE" +
		" `systemID` = VALUES(`systemID`)," +
		" `regionID` = VALUES(`regionID`)," +
		" `name` = VALUES(`name`);"

	_, err := DB.Exec(sql, id, systemID, name, systemID)
	if err != nil {
		log.Println("station.update: " + err.Error())
	}
}
