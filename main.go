package main

import (
	_ "embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/josheyr/dynamic-images/generate"
	"log"
	"net/http"
	"os/exec"
)

//go:embed index.html
var indexHtml string

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
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(indexHtml))
		if err != nil {
			log.Println(err)
		}
	})

	go func() {
		cmd := exec.Command("cmd", "/C", "start", "http://localhost:8080")
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
