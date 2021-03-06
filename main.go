package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pari-27/tinyURLProject/service"
	"github.com/pari-27/tinyURLProject/store"
	"log"
	"net/http"
)

func main() {
	db, err := store.Init()
	if err != nil {
		fmt.Println("Failed to setup database connection")
	}
	deps := service.Init(db)
	router := mux.NewRouter()
	router.HandleFunc("/url", service.CreateTinyURL(deps)).Methods(http.MethodPost)
	router.HandleFunc("/url", service.GetTinyURL(deps)).Methods(http.MethodGet)
	router.HandleFunc("/url", service.DeleteTinyURL(deps)).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))

}
