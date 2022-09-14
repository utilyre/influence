package middlewares

import "github.com/gofiber/fiber/v2"

func NewParamID() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if id <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "ensure that id is an unsigned integer")
		}

		c.Locals("id", uint(id))
		return c.Next()
	}
}
