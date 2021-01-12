package main

import (
	"github.com/gofiber/fiber"
	"github.com/jonathanbs9/movie-suggester/api"
)

func main() {
	app := fiber.New()
	api.SetupMoviesRoutes(app)
	_ = app.Listen(3001)
}
