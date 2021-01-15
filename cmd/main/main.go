package main

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/jonathanbs9/movie-suggester/api"
	"github.com/jonathanbs9/movie-suggester/internal"
)

func main() {
	app := fiber.New()
	internal.SetErrorHandler(app)
	api.SetupMoviesRoutes(app)
	api.SetupUsersRoutes(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretkey-changeme"),
	}))

	_ = app.Listen(3001)
}
