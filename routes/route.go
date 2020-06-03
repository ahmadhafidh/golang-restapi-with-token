package routes

import (
	"fmt"
	"log"
	"net/http"
	"spc-api-v2/services"

	"github.com/gorilla/mux"
)

func RouteStart() {
	router := mux.NewRouter()
	router.HandleFunc("/token/get", services.GetToken).Methods("GET")
	router.HandleFunc("/token/bni", services.GetTokenFromBNI).Methods("GET")
	fmt.Println("Connected to port 8003")
	log.Print(http.ListenAndServe(":8003", router))
}
