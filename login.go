package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	url := authRootUrl
	if url == "" {
		url = authRootUrlDev + "/token"
	} else {
		url = authRootUrl + "/token"
	}

	var dat map[string]interface{}
	json.NewDecoder(r.Body).Decode(&dat)

	if dat["password"] == nil || dat["email"] == nil {
		invalid_request(w, nil, "missing email or password")
		return
	}

	password := dat["password"].(string)
	email := dat["email"].(string)

	s := fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password)
	jsonStr := []byte(s)
	b := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		invalid_request(w, err, "Unable to send request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		access_denied(w, err, "Unable to login")
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		s := buf.String()
		created_request(w, s)
	}

}
