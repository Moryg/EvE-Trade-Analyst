package types

import (
	"github.com/julienschmidt/httprouter"
)

type RouteLoader func(*httprouter.Router) *httprouter.Router
