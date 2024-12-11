package user

import (
	"github.com/gofiber/fiber/v2"
	userServis "test/internal/service/user"
)

type RegisterReq struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (h *Handler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		if req.Login == "" || req.Password == "" || req.Nickname == "" {
			return fiber.ErrBadRequest
		}
		err := h.userService.Create(c.Context(), userServis.CreateUser{
			Login:     req.Login,
			Password:  req.Password,
			Nickname:  req.Nickname,
			Firstname: req.Firstname,
			Lastname:  req.Lastname},
		)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "user was registered success",
		})

	}
}
