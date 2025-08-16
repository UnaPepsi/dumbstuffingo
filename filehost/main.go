package main

import (
	"filehost/internal/api"
	"filehost/internal/db"
	"log"
)

func main() {
	err := db.Initialize()
	if err != nil {
		log.Fatalf("An error ocurred while trying to connect to postgres: %v", err.Error())
	}
	defer db.Close()
	if err := api.Listen(); err != nil {
		log.Fatalf("An error ocurred while listening to API: %v", err.Error())
	}
}
