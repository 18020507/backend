package handler

import (
	"encoding/json"
	"fmt"
	"hieu/pkgs/database"
	"net/http"
)

type MovieSelect struct {
	Name string `json:"id"`
}

func MovieSelected(w http.ResponseWriter, req *http.Request) {
	movieSelect := req.URL.Query()

	key, val := movieSelect["data"]
	if !val || len(key) == 0 {
		fmt.Println("key not present")
	}
	var movieName = key[0]

	movie_room_selected := database.GetMovieRoom(movieName)

	bb, err := json.Marshal(movie_room_selected)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, string(bb))

	// w.WriteHeader(200)
	// w.Write([]byte(strings.Join(key, ",")))
}
