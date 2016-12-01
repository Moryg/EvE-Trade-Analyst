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

  _ = s
  _ = routes

  log.Printf("Token 1: \"%s\"", apiqueue.GetValidToken(1))
  log.Printf("Token 2: \"%s\"", apiqueue.GetValidToken(2))
  log.Printf("Token 1: \"%s\"", apiqueue.GetValidToken(1))
  // log.Fatal(s.Start(routes))
}