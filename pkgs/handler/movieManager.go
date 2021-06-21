package handler

import (
	"encoding/json"
	"fmt"
	"hieu/pkgs/database"
	"net/http"
)

func MoviesManager(w http.ResponseWriter, req *http.Request) {

	movieManager := database.GetMovieManager()

	bb, err := json.Marshal(movieManager)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, string(bb))
}
