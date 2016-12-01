package apiqueue

import (
  "time"
)


func executor() {
  var r IRequest
  for {
    r = <- queue

    if r.RequiresAuth() {
      // get valid token
    }

    rps <- false
    defer decrementRPS()

    r.Execute()
  }
}

func decrementRPS() {
  time.Sleep(time.Second)
  <- rps
}