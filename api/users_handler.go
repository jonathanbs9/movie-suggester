package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {

	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "No se puede crear un usuario")
	}

	// En caso que vaya todo bien, creamos un token
	res.JWT = signToken(w.tokenKey, res.ID)
	return c.JSON(res)
}

/*token := jwt.New(jwt.SigningMethodHS256)

	// Seteamos claims =>
	claims := token.Claims.(jwt.MapClaims)
	claims["test-name"] = "Jonathan"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generamos un token encodeado y lo enviamos como response
	t, err := token.SignedString([]byte("mysecret-changeme"))

	res.JWT = t

	// En caso que no haya error creamos un json y  nos va a devolver
	_ = c.JSON(res)
}*/

/*func (w *WebServices) WishListHandler(c *fiber.Ctx) error {
	var cmd WishMovieCMD
	_ = c.BodyParser(&cmd)
	bearer := c.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.AddWishMovie(userID, cmd.MovieID, cmd.Comment)

	if err != nil {
		return fiber.NewError(400, "No se puede agregar a la lista de deseos")
	}
	// caso exitoso, devolvemos mensaje que se agregó la pelicula
	return c.JSON(struct {
		Wishlist string `json:"result"`
	}{
		Wishlist: "Película agregada a la WishList",
	})
}*/

/* Crea un espacio de video con html
*  Endpoint :3001/users/video
*  Guardamos un file en el root del proyecto.
 */
func (w *WebServices) ServeVideo(c *fiber.Ctx) error {
	c.Set("Content-Type", "video/mp4")
	err := c.SendFile("ruta-del-archivo.mp4", false)

	// Manejamos el error, en caso que haya alguno
	if err != nil {
		return fiber.NewError(400, "No se puede mostrar el video.")

	}
	return nil
}

/* Login Handler
 *
 */
func (w *WebServices) LoginHandler(c *fiber.Ctx) error {
	var cmd LoginCMD
	err := c.BodyParser(&cmd)

	if err != nil {
		return fiber.NewError(400, "No se puede parsear los parametros")
	}

	id := w.users.Login(cmd)

	// Si no existe el User
	if id == "" {
		return fiber.NewError(400, "Usuario no ha sido encontrado")
	}

	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: signToken(w.tokenKey, id),
	})
}

type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WishMovieCMD struct {
	MovieID string `json:"movie_id"`
	Comment string `json:"comment"`
}
