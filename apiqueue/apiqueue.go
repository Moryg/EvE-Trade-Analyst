package apiqueue

import (
  "log"

  . "github.com/moryg/eve_analyst/config"
)

var (
  queue     chan IRequest
  rps       chan bool
  basicAuth string
)

type IRequest interface {
  Execute()
  RequiresAuth() bool
}

func Enqueue (r IRequest) {
  queue <- r
}

func Start() {
  if (Config.EveAPI.RPS < 1) {
    log.Fatal("Missing EvE API Requests per second limit in config.json")
  }
  if (Config.EveAPI.Parallel < 1) {
    log.Fatal("Missing EvE API parralel requests limit in config.json")
  }
  if (len(Config.EveAPI.BasicAuth) < 1) {
    log.Fatal("Missing EvE API Basic auth code in config.json")
  }

  queue = make(chan IRequest, 10000)
  rps = make(chan bool, Config.EveAPI.RPS)
  basicAuth = "Basic " + Config.EveAPI.BasicAuth

  for ii := 0; ii < Config.EveAPI.Parallel; ii++ {
    go executor()
  }
}