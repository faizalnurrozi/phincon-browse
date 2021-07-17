package v1

import (
	"github.com/faizalnurrozi/phincon-browse/server/http/handlers"
	v1 "github.com/faizalnurrozi/phincon-browse/server/http/handlers/v1"
	"github.com/gofiber/fiber/v2"
)

type ResourceRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ResourceRoute) RegisterRoute() {

	// Initiate handler
	handler := v1.ResourceHandler{Handler: route.Handler}

	// List of route
	colorGroupRouters := route.RouteGroup.Group("/pokemon")
	colorGroupRouters.Get("", handler.Browse)
	colorGroupRouters.Get("/:id", handler.ReadBy)
}
