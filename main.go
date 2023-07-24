package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var countries = []Country{}

func main() {
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

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
