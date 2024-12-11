package err_handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {
	var (
		e          *fiber.Error
		statusCode int
	)
	if errors.As(err, &e) {
		statusCode = e.Code
	}
	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}
	return ctx.Status(statusCode).JSON(fiber.Map{
		"err": err.Error(),
	})
}
