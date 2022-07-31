package controller

import (
	"github.com/gofiber/fiber/v2"
)

func DefaultController(context *fiber.Ctx) error {

	context.Status(204)

	return context.Send(nil)
}
