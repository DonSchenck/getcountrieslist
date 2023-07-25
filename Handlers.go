package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Country struct {
	countryID         int    `json:"countryid"`
	countryName       string `json:"countryname"`
	population        int    `json:"population"`
	landAreaKM        int    `json:"landareakm"`
	populationDensity int    `json:"populationdensity"`
}
type Countries []Country

func AllCountries(w http.ResponseWriter, r *http.Request) {

	var countries = Countries{}

	// Create database handle
	db, err := sql.Open("mysql", "countries:countries@tcp(countries)/countries")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM countries")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var country Country
		err := res.Scan(&country.countryID, &country.countryName, &country.population, &country.landAreaKM, &country.populationDensity)

		if err != nil {
			log.Fatal(err)
		}

		log.Print(country)

		countries = append(countries, country)
	}

	log.Print(countries)
	json.NewEncoder(w).Encode(countries)
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getcountrieslist")
}

func OneCountry(w http.ResponseWriter, r *http.Request) {
	var countries = Countries{}

	// Create database handle
	db, err := sql.Open("mysql", "countries:countries@tcp(countries)/countries")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM countries")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var country Country
		err := res.Scan(&country.countryID, &country.countryName, &country.population, &country.landAreaKM, &country.populationDensity)

		if err != nil {
			log.Fatal(err)
		}

		log.Print(country)

		countries = append(countries, country)
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["countryId"])
	if err == nil {
		json.NewEncoder(w).Encode(countries[id])
	}
}

func RandomCountry(w http.ResponseWriter, r *http.Request) {
	var countries = Countries{}

	// Create database handle
	db, err := sql.Open("mysql", "countries:countries@tcp(countries)/countries")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM countries")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var country Country
		err := res.Scan(&country.countryID, &country.countryName, &country.population, &country.landAreaKM, &country.populationDensity)

		if err != nil {
			log.Fatal(err)
		}

		log.Print(country)

		countries = append(countries, country)
	}

	json.NewEncoder(w).Encode(countries[rand.Intn(150)])
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1.0.0")
}

func WrittenIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go")
}
