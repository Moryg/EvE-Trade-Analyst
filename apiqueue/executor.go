package apiqueue

import (
  "time"
)


func executor() {
  var r IRequest
  for {
    r = <- queue
    rps <- false

    defer decrementRPS()
    r.Execute()
  }
}

func decrementRPS() {
  time.Sleep(time.Second)
  <- rps
}