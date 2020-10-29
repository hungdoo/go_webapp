package main

import (
	"go_webapp/api"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/prime", api.PrimeHandler).Methods("POST")

	// Add the static file directory
	staticFileDirectory := http.Dir("./statics/")
	staticFileHandler := http.StripPrefix("/statics/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/statics/").Handler(staticFileHandler).Methods("GET")

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve index.html
	w.Header().Set("Content-type", "text/html")

	p := path.Dir("./statics/index.html")
	http.ServeFile(w, r, p)
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
