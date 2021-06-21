package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Movie_Managers struct {
	MovieName *string `json:"movieName"`
}

func GetMovieManager() []*Movie_Managers {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT movies.movie_name FROM movies")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var movieManger []*Movie_Managers

	for res.Next() {
		var movieMangers Movie_Managers
		err := res.Scan(&movieMangers.MovieName)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			panic(err)
		}

		movieManger = append(movieManger, &movieMangers)

	}
	return movieManger
}
