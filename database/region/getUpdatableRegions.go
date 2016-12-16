package region

import (
	"github.com/jmoiron/sqlx"
	. "github.com/moryg/eve_analyst/database"
	"log"
)

var getUpdatable *sqlx.Stmt

func init() {
	var err error

	// Increase number of regions later
	getUpdatable, err = DB.Preparex("SELECT `id` FROM `region` WHERE (`lastUpdate` IS NULL OR `lastUpdate` < DATE_SUB(NOW(), INTERVAL 30 minute)) ORDER BY `lastUpdate` ASC LIMIT 1;")

	if err != nil {
		log.Fatal(err)
	}
}

func GetUpdatableRegions() []uint64 {
	var id uint64
	results := []uint64{}

	rows, err := getUpdatable.Query()
	if err != nil {
		log.Println(err)
		return results
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			continue
		}
		results = append(results, id)
	}

	return results
}
