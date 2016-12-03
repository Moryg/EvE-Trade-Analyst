package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/moryg/eve_analyst/config"
	"log"
	"sync"
)

func DevQuery() {

}

var (
	DB      *sql.DB
	batches map[string]*upBatch
	lock    *sync.Mutex
)

func init() {
	dbi, err := sql.Open("mysql", Config.MySQL.String())
	if err != nil {
		log.Fatal(err)
	}

	DB = dbi
	batches = make(map[string]*upBatch)
	lock = new(sync.Mutex)
}

func Ping() {
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type upBatch struct {
	finished int
	total    int
	lock     *sync.Mutex
}

func newBatch(total int) *upBatch {
	b := new(upBatch)
	b.finished = 0
	b.total = total
	b.lock = new(sync.Mutex)
	return b
}

func (u *upBatch) Lock() {
	u.lock.Lock()
}

func (u *upBatch) Unlock() {
	u.lock.Unlock()
}

func (u *upBatch) Increment() bool {
	u.finished++
	return u.finished == u.total
}
