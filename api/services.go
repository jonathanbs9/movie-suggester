package api

import (
	"github.com/jonathanbs9/movie-suggester/internal/database"
	embedded "github.com/tomiok/fuego-cache/clients/inmemory"
)

type Services struct {
	search MovieSearch
	users  UserGateway
}

func NewServices() Services {
	client := database.NewMySQLClient()
	return Services{
		search: &MovieService{client},
		users:  &UserService{client},
	}
}

type WebServices struct {
	Services
	tokenKey string
	cache    *embedded.FuegoInMemory
}

func start(tokenKey string) *WebServices {
	inMemoryDB := embedded.NewInMemory(false, "")
	return &WebServices{NewServices(), tokenKey, inMemoryDB}
}
