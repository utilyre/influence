package routes

import (
	"os"
	"server/database"
	"server/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	password := user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user.Password = string(hash)

	if err := database.Instance.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	user.Password = password

	return c.Status(fiber.StatusCreated).JSON(user)
}

func SignIn(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.StructExcept(&user, "Name"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var found models.User
	if err := database.Instance.Where(&models.User{Email: user.Email}).Take(&found).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour),
		"email": found.Email,
	})

	signed, err := token.SignedString([]byte(os.Getenv("SERVER_ACCESS_TOKEN_SECRET")))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": signed,
	})
}
