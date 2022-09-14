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

func generateHashFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	return string(hash), nil
}

func compareHashWithPassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return nil
}

func SignUp(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	password := user.Password
	hash, err := generateHashFromPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := database.Instance.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	user.Password = password

	return c.Status(fiber.StatusCreated).JSON(user)
}

func SignIn(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	var found models.User
	findUser(user.Email, &found)

	if err := compareHashWithPassword(found.Password, user.Password); err != nil {
		return err
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
