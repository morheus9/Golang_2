package server

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {}

//Server starts server on 127.0.0.1:8080/
func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "44e") })
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "About Page") })

	log.Fatal(http.ListenAndServe(":8080", nil))

}
