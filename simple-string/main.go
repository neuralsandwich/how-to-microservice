package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrEmptyString = errors.New("empty string")

type stringRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	C int `json:"c"`
}

type uppercaseResponse struct {
	S string `json:"s"`
}

func count(s string) int {
	return len(s)
}

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	var request stringRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err)
	}
	if request.S == "" {
		json.NewEncoder(w).Encode(ErrEmptyString)
	}

	response := &countResponse{
		C: count(request.S),
	}
	json.NewEncoder(w).Encode(response)
}

func uppercaseHandler(w http.ResponseWriter, r *http.Request) {
	var request stringRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err)
	}

	if request.S == "" {
		json.NewEncoder(w).Encode(ErrEmptyString)
	}

	response := &uppercaseResponse{
		S: uppercase(request.S),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/uppercase", uppercaseHandler)
	http.ListenAndServe(":8080", nil)
}
