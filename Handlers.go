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
	countryID         int
	countryName       string
	population        int
	landAreaKM        int
	populationDensity int
}

func AllCountries(w http.ResponseWriter, r *http.Request) {

	var countries []*Country

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
		country := new(Country)
		err := res.Scan(&country.countryID, &country.countryName, &country.population, &country.landAreaKM, &country.populationDensity)

		if err != nil {
			log.Fatal(err)
		}

		log.Print(country)

		countries = append(countries, country)
	}

	if err := json.NewEncoder(w).Encode(countries); err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(countries)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getcountrieslist")
}

func OneCountry(w http.ResponseWriter, r *http.Request) {
	var countries []*Country

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
		country := new(Country)
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
	var countries []*Country

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
		country := new(Country)
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
