/**
 * Update or insert a station
 */
package station

import (
	"fmt"
	. "github.com/moryg/eve_analyst/database"
	"log"
	"strconv"
	"strings"
)

func Update(id, systemID int, name string) {
	sql := "insert into `station` (`id`, `systemID`, `regionID`, `name`)" +
		"SELECT ? AS `id`, ? as `systemId`, regionID, ? as `name`, ? AS `sortName`" +
		" FROM `system` WHERE `system`.`id` = ?" +
		" ON DUPLICATE KEY UPDATE" +
		" `systemID` = VALUES(`systemID`)," +
		" `regionID` = VALUES(`regionID`)," +
		" `name` = VALUES(`name`)" +
		" `sortName` = VALUES(`sortName`);"

	_, err := DB.Exec(sql, id, systemID, name, systemID, getSortName(name))
	if err != nil {
		log.Println("station.update: " + err.Error())
	}
}

func getSortName(name string) string {
	parts := strings.Split(name, " ")
	for ii, part := range parts {
		num, err := strconv.ParseUint(part, 10, 64)
		if err != nil {
			continue
		}
		parts[ii] = fmt.Sprintf("%03d", num)
	}

	return strings.Join(parts, " ")
}
