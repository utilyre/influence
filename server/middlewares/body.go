package middlewares

import (
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func NewUserParser(exceptions ...string) func(*fiber.Ctx) error {
	var user models.User
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(&user); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := user.Validate(exceptions...); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		c.Locals("user", user)
		return c.Next()
	}
}

func NewBlogParser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var blog models.Blog
		if err := c.BodyParser(&blog); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := blog.Validate(); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		c.Locals("blog", blog)
		return c.Next()
	}
}
