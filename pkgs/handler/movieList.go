package handler

import (
	"encoding/json"
	"fmt"
	"hieu/pkgs/database"
	"net/http"
)

func MoviesList(w http.ResponseWriter, req *http.Request) {

	movie := database.GetMovie()

	bb, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, string(bb))
}
