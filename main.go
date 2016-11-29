package main

import (
  "log"
  "github.com/moryg/eve_analyst/server"
  "github.com/moryg/eve_analyst/apiqueue"
  . "github.com/moryg/eve_analyst/config"
)

func main() {
  // API Queue setup
  apiqueue.Start()

  // Server setup
  s := server.Server{Config.HttpPort}
  routes := []server.RouteLoader {}

  log.Fatal(s.Start(routes))
}