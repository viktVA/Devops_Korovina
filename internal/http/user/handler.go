package user

import (
	"github.com/gofiber/fiber/v2"
	"test/internal/middleware"
)

type Handler struct {
	userService UserService
}

func NewHandler(userService UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func MapUserRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/register", h.Register())
	g.Post("/auth", h.Auth())
	g.Get("/getInfo", mw.Auth(), h.GetInfo())
}
