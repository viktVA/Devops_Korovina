package user

import "github.com/gofiber/fiber/v2"

type AuthReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		if req.Login == "" || req.Password == "" {
			return fiber.ErrBadRequest
		}
		token, err := h.userService.Auth(c.Context(), req.Login, req.Password)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token})
	}

}
