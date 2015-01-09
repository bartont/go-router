package main

import (
	//jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "Hello home")
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", Middleware(r))
	http.ListenAndServe(":3000", nil)
}
