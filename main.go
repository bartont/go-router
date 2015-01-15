package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	port = ":8001"
)

var authRootUrl = os.Getenv("ROOT_URL_AUTH")
var authRootUrlDev = "http://localhost:8000"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
}

func testoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testo page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/registration", registrationHandler)
	r.HandleFunc("/testo", testoHandler)
	http.Handle("/", validate(r))

	log.Println("Listening on port " + port + ". Go to http://localhost:8000/")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
