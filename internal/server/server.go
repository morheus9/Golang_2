package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/8host\n</h1>")
	})
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker!\n</h1>")
	})
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["name"]
		fmt.Fprintf(w, "<h1>Hello, %s!\n</h1>", title)
	})
	http.ListenAndServe(":80", r)
}

//Server starts server on 127.0.0.1:80/
