package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jonathanbs9/movie-suggester/api"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code por defecto 500
			code := fiber.StatusInternalServerError
			var msg string

			// Recupera el codido de estado personalizado si es un fiber error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "No se puede procesar la llamada HTTP"
			}

			// Envía custom error page
			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})
			return nil
		},
	})

	key := "tokenKey"
	app.Use(recover.New())
	api.SetupMoviesRoutes(app, key)
	api.SetupUsersRoutes(app, key)
	_ = app.Listen(":3001")
}

type internalError struct {
	Message string `json:"message"`
}
