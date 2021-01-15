package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) {

	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		err = fiber.NewError(400, "No se puede crear un usuario")
		c.Next()
		return
	}
	// En caso que vaya todo bien, creamos un toke
	token := jwt.New(jwt.SigningMethodHS256)

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
}

func (w *WebServices) WishListHandler(c *fiber.Ctx) {

	_ = c.JSON(struct {
		Wishlist string `json:"wish_list"`
	}{
		Wishlist: "Algunas peliculas acá",
	})
}
