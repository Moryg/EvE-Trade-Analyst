package database

import "log"

func GetUserTokens(userId int) {
  rows, err := db.Query("SELECT * FROM `oauth_token` WHERE `user_id` = ?", userId)
  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()
}
