package routes

import (
	"os"
	"server/database"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func parseAndValidateUser(c *fiber.Ctx, user *models.User, exceptions ...string) error {
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.Validate(exceptions...); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func SignUp(c *fiber.Ctx) error {
	var user models.User
	if err := parseAndValidateUser(c, &user); err != nil {
		return err
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
	if err := parseAndValidateUser(c, &user, "Name"); err != nil {
		return err
	}

	var found models.User
	findUser(user.Email, &found)

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour).Unix(),
		"email": found.Email,
	})

	signed, err := token.SignedString([]byte(os.Getenv("SERVER_ACCESS_TOKEN_SECRET")))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": signed,
	})
}

func findUser(email string, user *models.User) error {
	if err := database.Instance.Where(&models.User{Email: email}).Take(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return nil
}

func WhoAmI(c *fiber.Ctx) error {
	email := getEmailViaLocals(c)

	var user models.User
	if err := findUser(email, &user); err != nil {
		return err
	}

	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(user)
}
