package database

import (
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	. "github.com/moryg/eve_analyst/config"
	"log"
	"sync"
)

var (
	DB   *sqlx.DB
	lock *sync.Mutex
)

func init() {
	dbi, err := sqlx.Connect("mysql", Config.MySQL.String())
	if err != nil {
		log.Fatal(err)
	}

	DB = dbi
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(5)
	lock = new(sync.Mutex)
}

func Ping() {
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
