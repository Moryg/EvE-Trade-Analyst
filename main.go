package main

import (
  "log"

  "github.com/moryg/eve_analyst/server"
)

func main() {
  s := server.Server{8888}
  routes := []server.RouteLoader {}

  log.Fatal(s.Start(routes))

  // httpServer.Start(8001, routes)
}