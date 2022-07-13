package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tinyURLProject/service"
	"github.com/tinyURLProject/store"
	"log"
	"net/http"
)

func main() {
	db, err := store.Init()
	if err != nil {
		fmt.Println("Failed to setup database connection")
		return
	}
	deps := service.Init(db)
	router := mux.NewRouter()
	router.HandleFunc("/url", service.GetTinyURL(deps)).Methods(http.MethodPost)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())

}
