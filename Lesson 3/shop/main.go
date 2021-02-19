package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shop/notification"
	"shop/repository"
	"shop/service"
	"strings"
)

func main() {
	var tokenStr string
	flag.StringVar(&tokenStr, "t", "", "token for telegram api")

	var loginEmail string
	flag.StringVar(&loginEmail, "email", "", "Login mail.ru from which letters will be sent")

	var passwordEmail string
	flag.StringVar(&passwordEmail, "passwd", "", "Password for Email")

	flag.Parse()

	host := "smtp." + strings.Split(loginEmail, "@")[1]

	notif, err := notification.NewTelegramBot(tokenStr, 425969694)
	if err != nil {
		log.Fatal(err)
	}

	notifEmail, err := notification.NewEmail(loginEmail, passwordEmail, host)
	if err != nil {
		log.Fatal(err)
	}

	rep := repository.NewMapDB()
	service := service.NewService(rep, notif, notifEmail)
	//service := service.NewService(rep, notif)
	s := &server{
		service: service,
		rep:     rep,
	}

	router := mux.NewRouter()

	router.HandleFunc("/items", s.listItemHandler).Methods("GET")
	router.HandleFunc("/items", s.createItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", s.updateItemHandler).Methods("PUT")

	router.HandleFunc("/orders", s.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", s.createOrderHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
