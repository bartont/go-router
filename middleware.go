package main

import (
	"fmt"
	"log"
	"net/http"
)

func validate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()
		if path == "/registration" || path == "/login" {
			h.ServeHTTP(w, r)
		} else {
			url := authRootUrl
			if url == "" {
				url = authRootUrlDev + "/validate"
			} else {
				url = authRootUrl + "/validate"
			}

			auth := r.Header.Get("Authorization")

			req, err := http.NewRequest("PUT", url, nil)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println(err)
				fmt.Fprintf(w, err.Error())
			}

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
		}
	})
}
