package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/faizalnurrozi/phincon-browse/server/http/bootstrap/routers/v1"
	"github.com/faizalnurrozi/phincon-browse/server/http/handlers"
)

func (boot Bootstrap) RegisterRoute() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.Db,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	// Route for check health
	rootParentGroup := boot.App.Group("/product")
	rootParentGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("Work")
	})

	// Grouping v1 api
	v1Routers := v1.Routers{
		RouteGroup: rootParentGroup,
		Handler:    handlerType,
	}
	v1Routers.RegisterRoute()

	//// Occastion route
	//occassionRoute := routers.OccasionRoute{RouteGroup: apiV1, Handler: handlerType}
	//occassionRoute.RegisterRoute()
	//
	//// Seller Brand route
	//sellerBrandRoute := routers.SellerBrandRoute{RouteGroup: apiV1, Handler: handlerType}
	//sellerBrandRoute.RegisterRoute()
	//
	//// Gender route
	//genderRoute := routers.GenderRoute{RouteGroup: apiV1, Handler: handlerType}
	//genderRoute.RegisterRoute()
	//
	//// Seller Category route
	//categoryRoute := routers.CategoryRoute{RouteGroup: apiV1, Handler: handlerType}
	//categoryRoute.RegisterRoute()
	//
	//// Seller Group Color route
	//groupColorRoute := routers.ColorGroupRoute{RouteGroup: apiV1, Handler: handlerType}
	//groupColorRoute.RegisterRoute()
	//
	//// Item route
	//itemRoute := routers.ItemRoute{RouteGroup: apiV1, Handler: handlerType}
	//itemRoute.RegisterRoute()
}
