package server

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	. "github.com/moryg/eve_analyst/server/types"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	Port int
}

type empt struct {
}

func (e empt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\nPath: %s\nMethod: %s\n", r.Host, r.URL.RequestURI(), r.Method)
	bs, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Body: %s\n", string(bs))
}

func (s *Server) Start(routeLoaders []RouteLoader) error {
	if s.Port == 0 {
		return errors.New("Server port not set")
	}

	router := httprouter.New()
	for _, loader := range routeLoaders {
		loader(router)
	}

	// DEV HANDLE!!
	router.NotFound = empt{}

	log.Printf("Listening to localhost:%d", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), router))

	return nil
}
