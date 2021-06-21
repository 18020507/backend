package handler

import (
	"encoding/json"
	"hieu/pkgs/database"
	"net/http"
)

type Show struct {
	Start          string `json:"start"`
	End            string `json:"end"`
	Day_release    string `json:"day_release"`
	Movie_Selected string `json:"movie_selected"`
	Movie_room     int    `json:"movie_room"`
	Movie_format   string `json:"movie_format"`
}

func CreateShow(w http.ResponseWriter, req *http.Request) {
	show := Show{}
	err := json.NewDecoder(req.Body).Decode(&show)
	if err != nil {
		panic(err)
	}
	// fmt.Println(show)
	// fmt.Println(show.Start)
	// fmt.Println(show.End)
	// fmt.Println(show.Day_release)
	// fmt.Println(show.Movie_Selected)
	// fmt.Println(show.Movie_room)
	// fmt.Println(show.Movie_format)

	database.CreateShow(show.Start, show.End, show.Day_release, show.Movie_Selected, show.Movie_room, show.Movie_format)
}
