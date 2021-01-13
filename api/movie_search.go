package api

import (
	"github.com/jonathanbs9/movie-suggester/internal/database"
	"github.com/jonathanbs9/movie-suggester/internal/logs"
)

type MovieFilter struct {
	Title    string `json:"title,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Director string `json:"director,omitempty"`
}

type Movie struct {
	Id          string `json:id`
	Title       string `json:"title"`
	Cast        string `json:"cast"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
}

type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}

type MovieService struct {
	*database.MySqlClient
}

func (s *MovieService) Search(filter MovieFilter) ([]Movie, error) {

	tx, err := s.Begin()
	if err != nil {
		logs.Error("No se puede crear la transacción")
		return nil, err
	}

	rows, err := tx.Query(getMovieQueries())

	if err != nil {
		logs.Error("No se puede leer películas" + err.Error())
		// Hacemos un rollback de la transacción
		_ = tx.Rollback()
		return nil, err
	}

	var _movies []Movie

	// Leemos lo que tenemos en rows
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Cast, &movie.Genre, &movie.ReleaseDate, &movie.Director)
		if err != nil {
			logs.Error("No se pueden leer peliculas" + err.Error())
		}
		_movies = append(_movies, movie)
	}

	// Hacemos un commit de la transaccion. El ciclo de la transaccion queda listo
	_ = tx.Commit()

	/* Mock de movies
	m1 := Movie{
		Title:       "Blade Runner",
		Cast:        "Harrison Ford",
		ReleaseDate: time.Now(),
		Genre:       "Sciene Fiction",
		Director:    "James Camarón",
	}

	m2 := Movie{
		Title:       "Drive",
		Cast:        "Ryyan Gosling",
		ReleaseDate: time.Now(),
		Genre:       "Drama",
		Director:    "Some_director",
	}

	_movies = append(_movies, m1)
	_movies = append(_movies, m2)
	*/
	return _movies, nil
}
