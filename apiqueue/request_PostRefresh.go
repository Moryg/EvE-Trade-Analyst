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
    log.Printf("(PostRefresh) Json encode error usr: %d", usr.Id)
    return;
  }

  client := &http.Client{}
  req, _ := http.NewRequest("POST", "http://localhost:8888/oauth/token", bytes.NewReader(js))
  req.Header.Set("Authorization", basicAuth)
  req.Header.Set("Content-Type", "application/json")
  rsp, err := client.Do(req)

  if err != nil {
    log.Printf("(PostRefresh) Invalid Rtoken, usr %d", usr.Id)
    usr.Invalidate()
    return
  }

  defer rsp.Body.Close()

  var rspBody struct{
    Atoken string `json:"access_token"`
    Type   string `json:"token_type"`
    Expiry int    `json:"expires_in"`
    Ref    string `json:"refresh_token"`
  }

  // js = json.Unmarshal(rsp.Body, rspBody)
  err = json.NewDecoder(rsp.Body).Decode(&rspBody)
  if err != nil {
    log.Println("(PostRefresh) JSON Decode error" + err.Error())
    return
  }

  usr.NewAccessToken(rspBody.Atoken, rspBody.Expiry)
}

func (r RefreshAccessToken) RequiresAuth() bool {
  return false
}