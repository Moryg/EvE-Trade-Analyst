/**
 * Get a list of stations that are missing from the DB (market orders have player stations, which are not listed in the static export)
 */
package station

import (
	. "github.com/moryg/eve_analyst/database"
	"log"
	"strconv"
)

func GetMissingStations() []int {
	// TODO - OPTIMIZE Feels like it performs slowly for what it does
	var (
		id  string
		ids []int
	)
	sql := "SELECT DISTINCT `stationId` FROM `orderSell` `ord`" +
		" LEFT JOIN `station` `stat`" +
		" ON `stat`.`id` = `ord`.`stationId`" +
		" WHERE `stat`.`id` IS NULL" +
		" AND `stationId` < 1000000000000;" // Only stations, no citadels

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

		iId, err := strconv.Atoi(id)
		if err == nil {
			ids = append(ids, iId)
		}
	}

	return ids
}
