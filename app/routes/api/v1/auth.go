package v1

import (
	controllerV1 "goshaka/app/controllers/v1"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", validator.LoginValidator, controllerV1.Login)
}