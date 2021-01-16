package api

import "strings"

func getMovieQueries() string {
	return "select id, title, cast, genre, release_date, director from movie"
}

func getMoviesQuery(filter MovieFilter) string {
	var (
		d, g, t string
		clause  = false
		q       = "select id, title, cast, genre, release_date, director from movie"
		b       = strings.Builder{}
	)

	b.WriteString(q)

	// Pregunto si la query tiene Director
	if filter.Director != "" {
		d = "director like '%" + filter.Director + "%'"
		clause = true
	}

	// Pregunto si la query tiene Género
	if filter.Genre != "" {
		g = "genre like '%" + filter.Genre + "%'"
		clause = true
	}

	// Pregunto si la query tiene Titulo
	if filter.Title != "" {
		g = "genre like '%" + filter.Title + "%'"
		clause = true
	}

	if clause {
		var i int
		b.WriteString("where")
		if d != "" {
			b.WriteString(d)
			i = 1
		}

		if g != "" {
			if i == 1 {
				b.WriteString(" or ")
			}
			b.WriteString(g)
			i = 2
		}

		if t != "" {
			if i == 1 || i == 2 {
				b.WriteString(" or ")
			}
			b.WriteString(t)
		}

		return b.String()
	}
	return b.String()
}

func CreateUserQuery() string {
	// Query para crear User
	return "insert into user (id, username, password) values (?, ?, ?)"
}

func GetLoginQuery() string {
	return "select id from user where usermame =  ? and password = ?"
}