package apiqueue

import (
	. "github.com/moryg/eve_analyst/apiqueue/requests/usertoken"
	"sync"
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
