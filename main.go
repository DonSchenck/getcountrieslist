package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Country struct {
	countryID         int
	countryName       string
	population        int
	landAreaKM        int
	populationDensity int
}

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

		fmt.Printf("%v\n", country)
	}

}
