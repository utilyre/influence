package middlewares

import (
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func BodyBlog(c *fiber.Ctx) error {
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
