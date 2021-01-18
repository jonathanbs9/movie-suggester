package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jonathanbs9/movie-suggester/internal/logs"
)

func jwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func signToken(tokenkey, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	// Seteamos claims =>
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = id

	// Generamos un token encodeado y lo enviamos como response
	t, err := token.SignedString([]byte(tokenkey))

	if err != nil {
		return ""
	}

	return t
}

/* Creamos una funcion para extraer el ID de usuario desde
 * un JWT.
 */
func extractUserIDFromJWT(bearer, tokenKey string) string {
	// Le extraigo el "bearer " al token.
	token := bearer[7:]
	logs.Info(token)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	// Si no devuelve el token, error..
	if err != nil {
		return ""
	}

	// Si hay token ...
	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		//Obtengo el id
		return claims["sub"].(string)
	}

	return ""
}
