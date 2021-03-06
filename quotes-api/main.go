package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/gopher-quotes/handlers"
	"github.com/s1ntaxe770r/gopher-quotes/pkg/db"
	"github.com/sirupsen/logrus"
)

func main() {
	db := db.NewDB()
	h := handlers.NewQoutesHandler(db)
	r := mux.NewRouter()
	r.HandleFunc("/create", h.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", h.GetQuotes).Methods("GET")
	r.HandleFunc("/stats", h.GetStats).Methods("GET")
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	log.Println("Starting quotes api on 8080")
	logrus.Info(http.ListenAndServe(":8080", r))

}
