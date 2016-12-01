package ratelimit

import (
  "time"
  "log"
  . "github.com/moryg/eve_analyst/config"
)

var rps chan bool

func init() {
  if (Config.EveAPI.RPS < 1) {
    log.Fatal("Missing EvE API Requests per second limit in config.json")
  }
  rps = make(chan bool, Config.EveAPI.RPS)
}

func Add() {
  rps <- false
}

func Sub() {
  go func () {
    time.Sleep(time.Second)
    <- rps
  }()
}