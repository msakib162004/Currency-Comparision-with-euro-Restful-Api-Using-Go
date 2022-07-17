package controllers

import (
	"Currency-comaprasion/db"
	"Currency-comaprasion/structures"
	"encoding/json"
	"log"
	"net/http"
)

func Analyzer(w http.ResponseWriter, r *http.Request) {
	var (
		currencyName     string
		avg, min, max    float64
		RatesAnalyze     = make(map[string]map[string]float64)
		AnalyzerResponse structures.AnalyzerResponse
	)
	db := db.Connection()
	sqlStatement := `SELECT currencyname, AVG(currencyrate),MIN(currencyrate),MAX(currencyrate) FROM currencyrate GROUP BY currencyname`
	result, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		result.Scan(&currencyName, &avg, &min, &max)
		RatesAnalyze[currencyName] = map[string]float64{"min": min, "max": max, "avg": avg}
	}
	c := RatesAnalyze
	AnalyzerResponse = structures.AnalyzerResponse{
		Base:          "EUR",
		Rates_analyze: c,
	}
	q, _ := json.Marshal(AnalyzerResponse)
	w.Write(q)
}
