package database

import (
	"hieu/pkgs/connectDB"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Cinemas struct {
	Id         *int    `json:"id"`
	CinemaName *string `json:"cinemaName"`
}

func GetCinema() []*Cinemas {
	db := connectDB.Connect()

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
