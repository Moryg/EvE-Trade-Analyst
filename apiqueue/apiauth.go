package apiqueue

import (
  "time"
  "github.com/moryg/eve_analyst/database"
)

type UserToken struct {
  Id           int
  AccessToken  string
  RefreshToken string
  ValidTo      time.Time
}

var tokens map[int]UserToken

func init() {
  tokens = make(map[int]UserToken)
}

func GetValidToken(userId int) UserToken {
  usr, ok := tokens[userId]
  if ok {
    return usr
  }

  rToken, err := database.GetUserToken(userId)
  tokens[userId] = UserToken{userId, "", rToken, time.Now()}
  if err != nil {
    return tokens[userId]
  }

  return tokens[userId]
}