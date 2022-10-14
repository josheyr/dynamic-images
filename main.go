package main

import (
	"github.com/gorilla/mux"
	"github.com/josheyr/dynamic-images/generate"
	"log"
	"net/http"
)

func main() {
	// gorilla mux

	router := mux.NewRouter().StrictSlash(true)

	// CORS allow all
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/single", generate.GenerateSingleImage).Methods("POST")
	router.HandleFunc("/dynamic", generate.GenerateDynamicImages).Methods("POST")
	router.HandleFunc("/emptypixel/{info}", generate.HandlePixel).Methods("GET")
	// serve index.html
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(":8080", router))
}
