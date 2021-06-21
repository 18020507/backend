package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

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
