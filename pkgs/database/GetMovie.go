package database

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Movies struct {
	Id          *int    `json:"id"`
	MovieName   *string `json:"movieName"`
	MovieTime   *string `json:"movieTime"`
	DayRelease  *string `json:"dayRelease"`
	MovieTypeId *int    `json:"movieTypeId"`
	IdMovieType *int    `json:"idMovieType"`
	MovieType   *string `json:"movieType"`
}

func (movies *Movies) String() string {
	c, _ := json.Marshal(movies)
	return string(c)
}

func GetMovie() []*Movies {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT movies.id, movies.movie_name, movies.movie_time, movies.day_release, movie_types.id, movies.movie_type_id, movie_types.movie_type_name FROM movies INNER JOIN movie_types ON movies.movie_type_id = movie_types.id")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var movie []*Movies

	for res.Next() {
		var movies Movies
		err := res.Scan(&movies.Id, &movies.MovieName, &movies.MovieTime, &movies.DayRelease, &movies.MovieTypeId, &movies.IdMovieType, &movies.MovieType)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			panic(err)
		}

		movie = append(movie, &movies)

	}
	return movie
}
