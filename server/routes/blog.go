package routes

import (
	"fmt"
	"server/database"
	"server/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetBlogs(c *fiber.Ctx) error {
	strID := c.Params("id")
	if strID == "" {
		var blogs []models.Blog
		if err := findBlogs(&blogs); err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(blogs)
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("'%s' is not an integer", strID))
	}

	var blog models.Blog
	if err := findBlog(uint(id), &blog); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(blog)
}

func findBlogs(blogs *[]models.Blog) *fiber.Error {
	if err := database.Instance.Preload("Author").Find(blogs).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	for i := range *blogs {
		(*blogs)[i].Author.Password = ""
	}

	return nil
}

func findBlog(id uint, blog *models.Blog) *fiber.Error {
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

	validate := validator.New()
	if err := validate.Struct(&blog); err != nil {
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(&blog); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var found models.Blog
	if err := findBlog(uint(id), &found); err != nil {
		return err
	}

	email := getEmailViaLocals(c)
	if found.Author.Email != email {
		return fiber.NewError(fiber.StatusUnauthorized, "malformed token")
	}

	found.Apply(blog)

	if err := database.Instance.Save(&found).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(found)
}

func DeleteBlog(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var found models.Blog
	if err := findBlog(uint(id), &found); err != nil {
		return err
	}

	email := getEmailViaLocals(c)
	if found.Author.Email != email {
		return fiber.NewError(fiber.StatusUnauthorized, "malformed token")
	}

	if err := database.Instance.Delete(&found).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
