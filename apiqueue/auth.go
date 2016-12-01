package apiqueue

import (
  "sync"
  . "github.com/moryg/eve_analyst/apiqueue/usertoken"
)

var (
  tokens     map[int]*UserToken
  masterLock *sync.Mutex
)

func init() {
  tokens = make(map[int]*UserToken)
  masterLock = new(sync.Mutex)
}

func GetValidToken(userId int) string {
  var usr *UserToken

  masterLock.Lock()
  usr, ok := tokens[userId]

  if !ok {
    usr = NewToken(userId)
    tokens[userId] = usr
  }

  // Lock the token itself
  usr.InUse.Lock()
  // After we are done, unlock the token (AFTER the potential refresh!)
  defer usr.InUse.Unlock()

  // We can unlock the master at this point
  masterLock.Unlock()

  if !usr.IsValid() {
    // Access token is invalid, refresh it
    usr.RefreshToken()
  }

  return usr.AccessToken
}



// func (u *UserToken) Invalidate() {
//   u.RefreshToken = ""
//   // usr := GetValidToken(u.Id)
//   // usr.RefreshToken = ""
//   // usr.AccessToken = ""
//   // usr.ValidTo = time.Now().Add(time.Second * 5)
//   // tokens[u.Id] = usr
// }

// func (u UserToken) NewAccessToken(aToken string, duration int) {
//   u.AccessToken = aToken
//   if duration > 5 {
//     duration = duration - 5
//   }
//   u.ValidTo = time.Now().Add(time.Second * time.Duration(duration))
//   tokens[u.Id] = u
// }