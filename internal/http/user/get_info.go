package user

import (
	"github.com/gofiber/fiber/v2"
	"test/internal/entity"
)

type GetInfoResp struct {
	Id        entity.UserID `json:"id"`
	Login     string        `json:"login"`
	Nickname  string        `json:"nickname"`
	Firstname string        `json:"first_name"`
	Lastname  string        `json:"last_name"`
}

func (h *Handler) GetInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*entity.Claims)
		if !ok {
			return fiber.ErrUnauthorized
		}
		user, err := h.userService.GetUser(c.Context(), claims.UserID)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		return c.JSON(fiber.Map{
			"data": mapUserToGetInfoResp(user),
		})

	}
}

func mapUserToGetInfoResp(user entity.User) GetInfoResp {
	return GetInfoResp{
		Id:        user.ID,
		Login:     user.Login,
		Nickname:  user.Nickname,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}
}
