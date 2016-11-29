package server

import (
  "log"
  "fmt"
  "net/http"
  "errors"

  "github.com/julienschmidt/httprouter"
)

type Server struct {
  Port  int
}

type RouteLoader func(*httprouter.Router) *httprouter.Router

func (s *Server) Start(routeLoaders []RouteLoader) error {
  if (s.Port == 0) {
    return errors.New("Server port not set")
  }

  router := httprouter.New()
  for _,loader := range routeLoaders {
    loader(router)
  }

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), router))

  return nil
}
