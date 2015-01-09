package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	validateUrl = "http://localhost:8000/validate"
)

func validate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		req, _ := http.NewRequest("PUT", validateUrl, nil)
		req.Header.Set("Authorization", auth)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			fmt.Fprintf(w, err.Error())
		} else if resp.StatusCode == 200 {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			fmt.Fprintf(w, "unauthorized")
		}
	})
}
