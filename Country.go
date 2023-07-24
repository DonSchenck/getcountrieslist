package main

type Country struct {
	countryID         int    `json:"countryid"`
	countryName       string `json:"countryname"`
	population        int    `json:"population"`
	landAreaKM        int    `json:"landareakm"`
	populationDensity int    `json:"populationdensity"`
}
type Countries []Country
