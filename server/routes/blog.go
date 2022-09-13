package routes

import (
	"server/database"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetBlogs(c *fiber.Ctx) error {
	var blogs []models.Blog
	if err := findBlogs(&blogs); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(blogs)
}

func getIDParam(c *fiber.Ctx) (uint, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if id <= 0 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "ensure that id is an unsigned integer")
	}

	return uint(id), nil
}

func GetBlog(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	var blog models.Blog
	if err := findBlog(id, &blog); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(blog)
}

func findBlogs(blogs *[]models.Blog) error {
	if err := database.Instance.Preload("Author").Find(blogs).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	for i := range *blogs {
		(*blogs)[i].Author.Password = ""
	}

	return nil
}

func findBlog(id uint, blog *models.Blog) error {
	if err := database.Instance.Where(&models.Blog{ID: id}).Preload("Author").Take(&blog).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	blog.Author.Password = ""

	return nil
}

func CreateBlog(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := blog.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	email := getEmailViaLocals(c)

	var user models.User
	if err := findUser(email, &user); err != nil {
		return err
	}

	blog.AuthorID = user.ID
	blog.Author = user
	blog.Author.Password = ""

	if err := database.Instance.Create(&blog).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(blog)
}

func getEmailViaLocals(c *fiber.Ctx) string {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims["email"].(string)
}

func UpdateBlog(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := blog.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var found models.Blog
	if err := findBlog(id, &found); err != nil {
		return err
	}

	email := getEmailViaLocals(c)
	if err := authorizeEmail(email, found); err != nil {
		return err
	}

	found.Merge(blog)

	if err := database.Instance.Save(&found).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(found)
}

func DeleteBlog(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	var found models.Blog
	if err := findBlog(id, &found); err != nil {
		return err
	}

	email := getEmailViaLocals(c)
	if err := authorizeEmail(email, found); err != nil {
		return err
	}

	if err := database.Instance.Delete(&found).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func authorizeEmail(email string, blog models.Blog) error {
	if email != blog.Author.Email {
		return fiber.NewError(fiber.StatusUnauthorized, "user not authorized")
	}

	return nil
}
