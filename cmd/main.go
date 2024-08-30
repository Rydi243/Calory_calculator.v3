package main

import (
	"Calorie_calculator/internal/app"
	"Calorie_calculator/internal/config"
	"fmt"
	"log"
	"net/http"

	db2 "Calorie_calculator/internal/pkg/db"
)

func main() {
	db, err := db2.NewDB(config.PgDSN)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := app.NewService(db)

	fmt.Printf("Starting server at port 8080\n")
	if err = http.ListenAndServe(service.Port, service); err != nil {
		log.Fatal(err)
	}

}
