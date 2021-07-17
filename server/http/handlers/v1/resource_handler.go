package v1

import (
	handlers2 "github.com/faizalnurrozi/phincon-browse/domain/handlers"
	"github.com/faizalnurrozi/phincon-browse/server/http/handlers"
	v1 "github.com/faizalnurrozi/phincon-browse/usecase/v1"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ResourceHandler struct {
	handlers.Handler
}

func NewOccasionHandler(handler handlers.Handler) handlers2.IResourceHandler {
	return &ResourceHandler{Handler: handler}
}

// function handler for browse all data
func (handler ResourceHandler) Browse(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := v1.NewResourceUseCase(handler.UcContract)
	res, pagination, err := uc.Browse(page, limit)

	return handler.SendResponse(ctx, handlers.ResponseWithMeta, res, pagination, err, http.StatusUnprocessableEntity)
}
