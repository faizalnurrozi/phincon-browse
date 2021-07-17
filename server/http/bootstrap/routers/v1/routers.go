package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/faizalnurrozi/phincon-browse/server/http/handlers"
)

type Routers struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (routers Routers) RegisterRoute() {
	apiV1 := routers.RouteGroup.Group("/v1")

	brandRoutes := BrandRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	brandRoutes.RegisterRoute()

	genderRoutes := GenderRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	genderRoutes.RegisterRoute()

	categoryGroupRoutes := CategoryGroupRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	categoryGroupRoutes.RegisterRoute()

	categoryRoutes := CategoryRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	categoryRoutes.RegisterRoute()

	colorGroupRoutes := ColorGroupRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	colorGroupRoutes.RegisterRoute()

	colorRoutes := ColorRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	colorRoutes.RegisterRoute()

	occasionRoutes := OccasionRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	occasionRoutes.RegisterRoute()

	productRoutes := ProductRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	productRoutes.RegisterRoute()

	itemRoutes := ItemV1Route{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	itemRoutes.RegisterRoute()

	publicRoutes := PublicRoute{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	publicRoutes.RegisterRoute()
}
