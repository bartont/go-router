package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	port = ":8001"
)

var (
	publicKey []byte
)

func init() {
	pbk, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/demo.rsa.pub")
	if err != nil {
		log.Fatal("Unable to read public key", err)
	} else {
		publicKey = pbk
	}
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "Hello home")
}

func validate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/*		token, err := jwt.ParseFromRequest(r, func(t *jwt.Token) (interface{}, error) {
					return publicKey, nil
				})
				fmt.Println(token)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized) // Default is unauthorized
					log.Println(err)
					fmt.Fprintf(w, err.Error())
				} else if token.Valid {
					h.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized) // Default is unauthorized
					log.Println(err)
					fmt.Fprintf(w, err.Error())
				}*/
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", validate(r))

	log.Println("Listening on port " + port + ". Go to http://localhost:8000/")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
