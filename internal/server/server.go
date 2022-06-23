package server

import (
	"net/http"
	"text/template"
)

//Server starts server on 127.0.0.1:8081/
func Server() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("./static/index.html"))
	tmplt.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("./static/about.html"))
	tmplt.Execute(w, nil)
}
