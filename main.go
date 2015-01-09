package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	port = ":8001"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
}

func TestoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testo page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/testo", TestoHandler)
	http.Handle("/", validate(r))

	log.Println("Listening on port " + port + ". Go to http://localhost:8000/")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
