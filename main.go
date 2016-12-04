package main

import (
	"github.com/moryg/eve_analyst/apiqueue"
	// "github.com/moryg/eve_analyst/apiqueue/control"
	. "github.com/moryg/eve_analyst/config"
	"github.com/moryg/eve_analyst/server"
	marketRoutes "github.com/moryg/eve_analyst/server/market/routes"
	"github.com/moryg/eve_analyst/server/types"
	"log"
)

func main() {
	// API Queue setup
	apiqueue.Start()

	// control.BootUp()

	// Server setup
	s := server.Server{Config.HttpPort}
	routes := []types.RouteLoader{
		marketRoutes.Load,
	}

	log.Fatal(s.Start(routes))
}
