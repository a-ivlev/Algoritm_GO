package main

import (
	"shop/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	s := &server{
		rep: repository.NewMapDB(),
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", s.listItemHandler).Methods("GET")
	router.HandleFunc("/item", s.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", s.updateItemHandler).Methods("PUT")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
