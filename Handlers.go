package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllCountries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(countries)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getcountrieslist")
}

func OneCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["countryId"])
	if err == nil {
		json.NewEncoder(w).Encode(countries[id])
	}
}

func RandomCountry(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(countries[rand.Intn(150)])
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1.0.0")
}

func WrittenIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go")
}
