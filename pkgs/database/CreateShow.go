package database

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MovieID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (movID *MovieID) String() string {
	c, _ := json.Marshal(movID)
	return string(c)
}

func CreateShow(start, end, day_release, movie_selected string, movie_room int, movie_format string) {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT movies.id, movies.movie_name FROM movies WHERE movies.movie_name = ?", movie_selected)

	if err != nil {
		log.Fatal(err)
	}

	var moviesID []*MovieID

	for res.Next() {
		var movieID MovieID
		err := res.Scan(&movieID.ID, &movieID.Name)

		if err != nil {
			panic(err)
		}
		moviesID = append(moviesID, &movieID)
	}

	addShow, err := db.Query("INSERT INTO movie_showtimes (time_begin, time_end, day_showtime, movie_id, movie_room_id, movie_format) VALUES (?, ?, ?, ?, ?, ?)", start, end, day_release, moviesID[0].ID, movie_room, movie_format)

	if err != nil {
		panic(err)
	}

	defer addShow.Close()
}
