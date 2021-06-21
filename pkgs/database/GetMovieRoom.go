package database

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Movie_Rooms struct {
	ID            *int    `json:"id"`
	MovieRoomName *string `json:"movieName"`
}

func GetMovieRoom(cinema string) []*Movie_Rooms {
	db, err := sql.Open("mysql", "root:anhvahieu2k@tcp(127.0.0.1:3306)/Cinema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	id, err := strconv.Atoi(cinema)
	if err != nil {
		panic(err)
	}

	res, err := db.Query("SELECT movie_rooms.id, movie_rooms.movie_room_name FROM movie_rooms WHERE movie_rooms.cinema_id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var movie_rooms []*Movie_Rooms

	for res.Next() {
		var movie_room Movie_Rooms
		err := res.Scan(&movie_room.ID, &movie_room.MovieRoomName)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			panic(err)
		}

		movie_rooms = append(movie_rooms, &movie_room)

	}
	return movie_rooms
}
