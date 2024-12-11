package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"test/internal/entity"
	"time"
)

func (mw *MwManager) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()["Authorization"][0]
		if header == "" {
			return fiber.ErrUnauthorized
		}
		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return fiber.ErrUnauthorized
		}
		token := parts[1]

		claims := &entity.Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return mw.cfg.PublicKey, nil
		})
		if err != nil {
			return err
		}

		now := time.Now()

		if claims.ExpiresAt.Equal(now) || claims.ExpiresAt.Before(now) {
			return fiber.ErrUnauthorized
		}

		c.Locals("claims", claims)
		return c.Next()

	}
}
