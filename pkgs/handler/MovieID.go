package handler

import (
	"encoding/json"
	"fmt"
	"hieu/pkgs/database"
	"net/http"
)

type MovieId struct {
	Id int `json:"id"`
}

func MovieID(w http.ResponseWriter, req *http.Request) {
	// movieID := MovieId{}

	// err := json.NewDecoder(req.Body).Decode(&movieID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Fprintf(w, "error when parse body")
	// 	return
	// }
	// fmt.Println(movieID.Id)

	// fmt.Fprintf(w, "Movie with ID: %v", movieID.Id)
	movieID := req.URL.Query()

	key, val := movieID["id"]
	if !val || len(key) == 0 {
		fmt.Println("key not present")
	}

	// w.WriteHeader(200)
	// w.Write([]byte(strings.Join(key, ",")))

	movieDetail := database.GetMovieDetail(key[0])

	bb, err := json.Marshal(movieDetail)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, string(bb))
}
