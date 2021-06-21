package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MovieDetail struct {
	Id               *int    `json:"id"`
	MovieName        *string `json:"movieName"`
	MovieTime        *string `json:"movieTime"`
	DayRelease       *string `json:"dayRelease"`
	MovieLanguage    *string `json:"movieLanguage"`
	Producer         *string `json:"producer"`
	ShortDescription *string `json:"shortDescription"`
}

func GetMovieDetail(id string) []*MovieDetail {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT movies.id, movies.movie_name, movies.movie_time, movies.day_release, movies.movie_language, movies.producer, movies.short_description FROM movies WHERE movies.id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var movieDetail []*MovieDetail

	for res.Next() {
		var moviesDetail MovieDetail
		err := res.Scan(&moviesDetail.Id, &moviesDetail.MovieName, &moviesDetail.MovieTime, &moviesDetail.DayRelease, &moviesDetail.MovieLanguage, &moviesDetail.Producer, &moviesDetail.ShortDescription)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			panic(err)
		}

		movieDetail = append(movieDetail, &moviesDetail)
	}
	return movieDetail
}
