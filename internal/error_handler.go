package internal

import (
	"github.com/gofiber/fiber"
)

/* Vamos a tener una funcion que
 *
 */

func SetErrorHandler(app *fiber.App) {

	/* Setemos una funcion que va manejar los errores
	*
	 */
	app.Settings.ErrorHandler = func(ctx *fiber.Ctx, err error) {
		// Status Code Default to 500
		code := fiber.StatusInternalServerError
		msg := "No se puede procesar la llamada"

		// Retrieve the custom statusCode if it's a fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			msg = e.Message
		}

		// Send Custom error page
		err = ctx.Status(code).JSON(internalError{
			Message: msg,
		})

		if err != nil {
			ctx.Status(500).SendString("Internal Server Error! ")
		}
	}
}

type internalError struct {
	Message string `json: "message"`
}
