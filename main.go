package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))

}
