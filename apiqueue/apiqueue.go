package apiqueue

import (
  "log"
  "time"

  . "github.com/moryg/eve_analyst/config"
)

var queue chan IRequest
var rps   chan bool

type IRequest interface {
  Execute() string
}

func Start() {
  if (Config.EveAPI.RPS < 1) {
    log.Fatal("Missing EvE API Requests per second limit in config.json")
  }
  if (Config.EveAPI.Parallel < 1) {
    log.Fatal("Missing EvE API parralel requests limit in config.json")
  }

  queue = make(chan IRequest, 10000)
  rps = make(chan bool, Config.EveAPI.RPS)

  for ii := 0; ii < Config.EveAPI.Parallel; ii++ {
    go executor()
  }
}

func decrementRPS() {
  time.Sleep(time.Second)
  <- rps
}

func Enqueue (r IRequest) {
  queue <- r
}

func executor() {
  var r IRequest
  for {
    r = <- queue
    rps <- false
    defer decrementRPS()

    time.Sleep(time.Millisecond * 500)
    log.Printf("Request completed %v", r.Execute())
  }
}