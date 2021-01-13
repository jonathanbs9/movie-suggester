package main

import (
	"github.com/gofiber/fiber"
	"github.com/jonathanbs9/movie-suggester/api"
	"github.com/jonathanbs9/movie-suggester/internal"
)

func main() {
	app := fiber.New()
	internal.SetErrorHandler(app)
	api.SetupMoviesRoutes(app)
	_ = app.Listen(3001)
}
