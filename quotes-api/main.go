package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/gopher-quotes/handlers"
	"github.com/s1ntaxe770r/gopher-quotes/pkg/db"
	"github.com/sirupsen/logrus"
)

func main() {
	host := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")

	db_uri := fmt.Sprintf("host=%s user=%s password=%s dbname=quotes port=5432 sslmode=disable", host, dbuser, dbpass)
	db := db.NewDB(db_uri)
	h := handlers.NewQoutesHandler(db)
	r := mux.NewRouter()
	r.HandleFunc("/create", h.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", h.GetQuotes).Methods("GET")
	r.HandleFunc("/stats", h.GetStats).Methods("GET")
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	log.Println("Starting quotes api on 8080")
	logrus.Info(http.ListenAndServe(":8080", r))

}
