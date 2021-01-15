package api

import (
	"github.com/gofiber/utils"
	"github.com/jonathanbs9/movie-suggester/internal/database"
	"github.com/jonathanbs9/movie-suggester/internal/logs"
)

type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserSummary struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login()
}

type UserService struct {
	// Creamos una referencia a la base de datos
	*database.MySqlClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {
	id := utils.UUID()
	// Ejecutamos la query para guardar un user
	_, err := us.Exec(CreateUserQuery(), id, cmd.Username, cmd.Password)

	if err != nil {
		logs.Error("No se puede insertar ... =>" + err.Error())
		return nil, err
	}

	return &UserSummary{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

func (us *UserService) Login() {

}
