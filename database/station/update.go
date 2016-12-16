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

func Update(id uint64, systemID int, name string) {
	sql := "CALL saveStation(?, ?, ?, ?)"

	_, err := DB.Exec(sql, id, systemID, name, getSortName(name))
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
