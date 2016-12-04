package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	. "github.com/moryg/eve_analyst/server/types"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Port int
}

func (s *Server) Start(routeLoaders []RouteLoader) error {
	if s.Port == 0 {
		return errors.New("Server port not set")
	}

	router := httprouter.New()
	for _, loader := range routeLoaders {
		loader(router)
	}

	log.Printf("Listening to localhost:%d", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), router))

	return nil
}
