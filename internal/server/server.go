package server

import (
	"html/template"
	"log"
	"net/http"
)

func static(w http.ResponseWriter, _ *http.Request) {
	var tmplt = template.Must(template.ParseFiles("./static/index.html"))
	tmplt.Execute(w, nil)
}

func about(w http.ResponseWriter, _ *http.Request) {
	var tmplt = template.Must(template.ParseFiles("./static/about.html"))
	tmplt.Execute(w, nil)
}

//Server starts server on 127.0.0.1:8080/
func Server() {

	http.HandleFunc("/", static)
	http.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
