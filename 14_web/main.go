package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func thing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "<h1>Thing %s</h1>", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/thing/{id}", thing).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/", r)

	port := 3000

	fmt.Printf("Server starting at http://localhost:%d\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
