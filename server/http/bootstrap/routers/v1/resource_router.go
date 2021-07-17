package v1

import (
	"github.com/faizalnurrozi/phincon-browse/server/http/handlers"
	v1 "github.com/faizalnurrozi/phincon-browse/server/http/handlers/v1"
	"github.com/faizalnurrozi/phincon-browse/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ResourceRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ResourceRoute) RegisterRoute() {

	// Initiate handler
	handler := v1.ResourceHandler{Handler: route.Handler}
	basicAuthMiddleware := middlewares.BasicAuth{Contract: route.Handler.UcContract}

	// List of route
	pokemonGroupRouters := route.RouteGroup.Group("/pokemon")
	pokemonGroupRouters.Use(basicAuthMiddleware.BasicAuthNew())
	pokemonGroupRouters.Get("", handler.Browse)
	pokemonGroupRouters.Get("/:id", handler.ReadBy)
}
