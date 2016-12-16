/**
 * Get a list of stations that are missing from the DB (market orders have player stations, which are not listed in the static export)
 */
package station

import (
	. "github.com/moryg/eve_analyst/database"
	"log"
	// "strconv"
)

func GetMissingStations() []uint64 {
	// TODO - OPTIMIZE Feels like it performs slowly for what it does
	var (
		id  uint64
		ids []uint64
	)
	sql := "SELECT `t1`.`id` FROM (SELECT DISTINCT `stationId` `id` FROM `orderSell` WHERE `stationId` < 1000000000000) `t1` LEFT JOIN `station` ON `station`.`id` = `t1`.`id` WHERE `station`.`id` IS NULL LIMIT 100;"

	rows, err := DB.Query(sql)
	if err != nil {
		log.Println("station.GetMissingStations: " + err.Error())
		return ids
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println("station.GetMissingStations: " + err.Error())
			continue
		}

		ids = append(ids, id)
		// iId, err := strconv.Atoi(id)
		// if err == nil {
		// 	ids = append(ids, iId)
		// }
	}

	return ids
}
