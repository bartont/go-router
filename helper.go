package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func serveError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}

func redirectHandler(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, http.StatusMovedPermanently)
	}
	// usage: http.HandleFunc("/1", redirectHandler("/one"))
}

func parseForm(r *http.Request) error {
	if r.Form != nil {
		return errors.New("Request body was parsed already.")
	}
	tmp := r.URL.RawQuery
	r.URL.RawQuery = ""
	if err := r.ParseForm(); err != nil {
		return err
	}
	r.URL.RawQuery = tmp
	return nil
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Page not found!")
}

func notlsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in notlsHandler")
	fullUrl := "https://localhost" + r.RequestURI
	http.Redirect(w, r, fullUrl, http.StatusMovedPermanently)
}
