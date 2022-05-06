package main

import (
	"log"

	"github.com/Ara-Chobanyan/school-manager/school-data/api"
	"github.com/Ara-Chobanyan/school-manager/school-data/app"
	"github.com/Ara-Chobanyan/school-manager/school-data/repository/postgres"
)

func main() {
	repo, err := postgres.NewPostgresRepository()

	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	service := app.NewDataService(repo)

	rpc := api.NewAdapter(service)

	err = rpc.Run()

	if err != nil {
		log.Fatalf("Could not run: %v", err)
	}

}
