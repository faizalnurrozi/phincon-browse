package handlers

import "github.com/gofiber/fiber/v2"

type IResourceHandler interface {
	Browse(ctx *fiber.Ctx) (err error)

	ReadBy(ctx *fiber.Ctx) (err error)
}
