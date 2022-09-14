package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/middlewares"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
	setupRoutes(app)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("SERVER_PORT"))
	app.Listen(addr)
}

func setupRoutes(app *fiber.App) {
	auth := jwtware.New(jwtware.Config{
		SigningMethod: jwtware.HS256,
		SigningKey:    []byte(os.Getenv("SERVER_ACCESS_TOKEN_SECRET")),
	})

	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("/signup", routes.SignUp)
	users.Post("/signin", routes.SignIn)
	users.Get("/whoami", auth, routes.WhoAmI)

	blogs := api.Group("/blogs")
	blogs.Get("/", routes.GetBlogs)
	blogs.Get("/:id", middlewares.ParamID, routes.GetBlog)
	blogs.Post("/", auth, middlewares.BodyBlog, routes.CreateBlog)
	blogs.Put("/:id", auth, middlewares.ParamID, middlewares.BodyBlog, routes.UpdateBlog)
	blogs.Delete("/:id", auth, middlewares.ParamID, routes.DeleteBlog)
}
