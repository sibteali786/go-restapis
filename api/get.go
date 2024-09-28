package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/femm/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0]) // because id is a slice of string ([]string) and not a single string because its protocol like /api/id=32,21
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
