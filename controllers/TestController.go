package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request interface{}
	var response interface{}
	errors := decoder.Decode(&request)
	if errors != nil {
		log.Print("no input")
		response = "no input"
	} else {
		response = request
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
