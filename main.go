package main

import (
	"Currency-comaprasion/db"
	"Currency-comaprasion/models"
	"Currency-comaprasion/routers"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	models.InsertCurrencyModels(db.Connection())
	mux := routers.Routers()
	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}
