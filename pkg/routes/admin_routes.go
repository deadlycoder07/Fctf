package routes

import (
	"github.com/deadlycoder07/Fctf/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(a *fiber.App) {
	//group routes
	route := a.Group("/api/v1")

	//Routes for settings

	// GET Routes
	route.Get("/settings", middleware.JWTProtected(), nil)

	// POST Routes

}
