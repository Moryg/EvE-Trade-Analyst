package apiqueue

import (
  "encoding/json"
  "bytes"
  "log"
  "net/http"
  "time"
)

type RefreshAccessToken struct {
  UserId int
}

func (r RefreshAccessToken) Execute() {
  time.Sleep(time.Second)
  usr := GetValidToken(r.UserId)
  if usr.RefreshToken == "" {
    return
  }

  rbody := struct{
    Grant string `json:"grant_type"`
    Token string `json:"refresh_token"`
  }{
    "refresh_token",
    usr.RefreshToken,
  }

  js, err := json.Marshal(rbody)
  if err != nil {
    usr.Invalidate()
    log.Printf("Refresh Token json encode error user: %d", usr.Id)
    return;
  }

  client := &http.Client{}
  req, _ := http.NewRequest("POST", "http://localhost:8888/oauth/token", bytes.NewReader(js))
  req.Header.Set("Authorization", basicAuth)
  req.Header.Set("Content-Type", "application/json")
  _, err = client.Do(req)

  if err != nil {
    log.Printf("Invalid refresh token user %d", usr.Id)
    usr.Invalidate()
    return
  }
}

func (r RefreshAccessToken) RequiresAuth() bool {
  return false
}