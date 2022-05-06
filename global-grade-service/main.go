package main

import (
	"log"

	"github.com/Ara-Chobanyan/school-manager/global-grade-service/api"
	"github.com/Ara-Chobanyan/school-manager/global-grade-service/app"
	"github.com/Ara-Chobanyan/school-manager/global-grade-service/repository/postgres"
)

func main() {
	repo, err := postgres.NewPostgresRepository()

	if err != nil {
		log.Fatalf("something went wrong with repository: %v", err)
	}

	service := app.NewGlobalService(repo)

	rpc := api.NewAdapter(service)

	err = rpc.Run()

	if err != nil {
		log.Fatalf("Could not run: %v", err)
	}

}
