package database

import (
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  . "github.com/moryg/eve_analyst/config"
)

var db *sql.DB

func init() {
  dbi, err := sql.Open("mysql", Config.MySQL.String())
  if err != nil {
    log.Fatal(err)
  }

  db = dbi
}

func Ping() {
  err := db.Ping()
  if err != nil {
    log.Fatal(err)
  }
}