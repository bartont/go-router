package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := authRootUrl
	if url == "" {
		url = authRootUrlDev + "/token"
	} else {
		url = authRootUrl + "/token"
	}

	var dat map[string]interface{}
	json.NewDecoder(r.Body).Decode(&dat)

	if dat["password"] == nil || dat["email"] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "missing email or password")
		return
	}

	password := dat["password"].(string)
	email := dat["email"].(string)

	s := fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password)
	jsonStr := []byte(s)
	b := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to send request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to login", err)

	} else {
		w.WriteHeader(http.StatusCreated)

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		s := buf.String()

		fmt.Fprintf(w, s)
	}

}
