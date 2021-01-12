package api

import "github.com/gofiber/fiber"

func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) {
	res, err := w.s.search.Search(MovieFilter{})

	if err != nil {
		err = fiber.NewError(400, "No se pueden traer peliculas")
		c.Next(err)
	}

	c.JSON(res)
}
