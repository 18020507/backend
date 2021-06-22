package handler

import (
	"encoding/json"
	"fmt"
	"hieu/pkgs/database"
	"net/http"
)

func Cinema(w http.ResponseWriter, req *http.Request) {

	cinema := database.GetCinema()

	bb, err := json.Marshal(cinema)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, string(bb))
}
