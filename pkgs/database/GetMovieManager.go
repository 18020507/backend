package database

import (
	"hieu/pkgs/connectDB"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Movie_Managers struct {
	MovieName *string `json:"movieName"`
}

func GetMovieManager() []*Movie_Managers {
	db := connectDB.Connect()

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
