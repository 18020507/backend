package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Cinemas struct {
	Id         *int    `json:"id"`
	CinemaName *string `json:"cinemaName"`
}

func GetCinema() []*Cinemas {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM cinemas")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var cinema []*Cinemas

	for res.Next() {
		var cinemas Cinemas
		err := res.Scan(&cinemas.Id, &cinemas.CinemaName)

		if err != nil {
			log.Fatal(err)
		}
		// bb, err := json.Marshal(&movies)

		if err != nil {
			panic(err)
		}

		cinema = append(cinema, &cinemas)
	}
	return cinema
}
