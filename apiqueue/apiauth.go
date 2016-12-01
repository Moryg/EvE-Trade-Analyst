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

func (u *UserToken) Invalidate() {
  u.RefreshToken = ""
  usr := GetValidToken(u.Id)
  usr.RefreshToken = ""
  usr.AccessToken = ""
  usr.ValidTo = time.Now().Add(time.Second * 5)
  tokens[u.Id] = usr
}

func (u UserToken) NewAccessToken(aToken string, duration int) {
  u.AccessToken = aToken
  if duration > 5 {
    duration = duration - 5
  }
  u.ValidTo = time.Now().Add(time.Second * time.Duration(duration))
  tokens[u.Id] = u
}