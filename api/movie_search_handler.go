package api

import "github.com/gofiber/fiber/v2"

/* Funcionalidad de fiber v2 que devuelva un error. Nos ahorra varias
 * cosas, ya que c/vez que tenemos un error direcamente lo devolvemos.
 * Ademas cuando tenemos un path que es correcto, el que hace el marshall
 *  de la estrucura json, lo devuelve en la misma linea. return JSON
 */
func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) error {
	res, err := w.search.Search(MovieFilter{
		Title:    c.Query("title"),
		Genre:    c.Query("genre"),
		Director: c.Query("director"),
	})

	if err != nil {
		return fiber.NewError(400, "No se pueden traer peliculas")
	}

	if len(res) == 0 {
		return c.JSON([]interface{}{})
	}

	return c.JSON(res)
}
