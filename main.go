package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

	var countries = []Country{}

	for res.Next() {
		var country Country
		err := res.Scan(&country.countryID, &country.countryName, &country.population, &country.landAreaKM, &country.populationDensity)

		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}
}
