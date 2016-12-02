package usertoken

import (
  "time"
  "bytes"
  "encoding/json"
  "net/http"
  "sync"
  "log"
  "os"

  "github.com/moryg/eve_analyst/database"
  "github.com/moryg/eve_analyst/apiqueue/ratelimit"
  . "github.com/moryg/eve_analyst/config"
)

type UserToken struct {
  Id           int
  AccessToken  string
  ValidTo      time.Time
  InUse        *sync.Mutex
  rToken       string
}

// Refresh Request JSON Body
type reqBody struct{
  Grant string `json:"grant_type"`
  Token string `json:"refresh_token"`
}

type resBody struct{
  Atoken string `json:"access_token"`
  Type   string `json:"token_type"`
  Expiry int    `json:"expires_in"`
  Ref    string `json:"refresh_token"`
}

var (
  basicAuth, baseURL string
)

/**
 * Bootup function
 */
func init() {
  if (len(Config.EveAPI.BasicAuth) < 1) {
    log.Fatal("Missing EvE API Basic auth code in config.json")
  }

  baseURL = os.Getenv("API")
  basicAuth = Config.EveAPI.BasicAuth
}

/**
 * UserToken constructor
 */
func NewToken(userId int) *UserToken {
  u := new(UserToken)
  u.Id = userId
  u.rToken, _ = database.GetUserToken(userId)
  u.InUse = new(sync.Mutex)
  return u
}


func (t UserToken) IsValid() bool {
  return time.Now().Before(t.ValidTo)
}

func (t *UserToken) RefreshToken() {
  t.AccessToken = ""
  if t.rToken == "" {
    return
  }

  // Construct the request JSON
  b := reqBody{"refresh_token", t.rToken}
  jBody, err := json.Marshal(b)
  if err != nil {
    t.rToken = ""
    log.Println("usertoken.Refresh:" + err.Error())
    return;
  }

  // Build the request
  client := &http.Client{}
  req, err := http.NewRequest(
    "POST",
    (baseURL + "/oauth/token"),
    bytes.NewReader(jBody),
  )
  if err != nil {
    t.rToken = ""
    log.Println("usertoken.Refresh json encode:" + err.Error())
    return
  }
  req.Header.Set("Authorization", "Basic " + basicAuth)
  req.Header.Set("Content-Type", "application/json")

  // Execute the request
  ratelimit.Add()
  res, err := client.Do(req)
  ratelimit.Sub()

  // Handle response
  if err != nil {
    log.Println("usertoken.Refresh request exec:" + err.Error())
    t.rToken = ""
    return
  }
  defer res.Body.Close()

  // Parse json response
  var resJson resBody
  err = json.NewDecoder(res.Body).Decode(&resJson)

  // TODO - check for error msg in json
  if err != nil {
    log.Println("usertoken.Refresh rsp decode:" + err.Error())
    t.rToken = ""
    return
  }

  t.AccessToken = resJson.Atoken
  t.ValidTo = time.Now().Add(time.Duration(resJson.Expiry - 5) * time.Second)
}