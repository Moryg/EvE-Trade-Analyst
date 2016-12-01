package apiqueue

import (
  "log"
  . "github.com/moryg/eve_analyst/config"
)

var (
  queue     chan IRequest
)

type IRequest interface {
  Execute()
  RequiresAuth() bool
}

func Enqueue (r IRequest) {
  queue <- r
}

func Start() {
  if (Config.EveAPI.Parallel < 1) {
    log.Fatal("Missing EvE API parralel requests limit in config.json")
  }

  queue = make(chan IRequest, 10000)

  for ii := 0; ii < Config.EveAPI.Parallel; ii++ {
    go executor()
  }
}