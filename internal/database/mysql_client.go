package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jonathanbs9/movie-suggester/internal/logs"
)

type MySqlClient struct {
	*sql.DB
}

func NewMySQLClient() *MySqlClient {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/movies")

	if err != nil {
		logs.Error("No se puede crear cliente Mysql")
		panic(err)
	}

	err = db.Ping()

	if err != nil {

	}

	return &MySqlClient{db}
}
