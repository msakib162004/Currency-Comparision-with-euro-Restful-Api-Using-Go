package routers

import (
	"Currency-comaprasion/controllers"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/rates/latest", controllers.Latest)
	mux.HandleFunc("/rates/analyze", controllers.Analyzer)
	mux.HandleFunc("/rates/{date}", controllers.DateRate)

	return mux
}
