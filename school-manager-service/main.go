package main

import (
	"log"

	"github.com/Ara-Chobanyan/school-manager/school-manager-service/api"
	"github.com/Ara-Chobanyan/school-manager/school-manager-service/app"
	"github.com/Ara-Chobanyan/school-manager/school-manager-service/repository/postgres"
)

func main() {
	repo, err := postgres.NewPostgresRepository()

	if err != nil {
		log.Fatalf("Something went wrong: %v", repo)
	}

	service := app.NewManagerService(repo)

	rpc := api.NewAdapter(service)

	err = rpc.Run()

	if err != nil {
		log.Fatalf("Could not run: %v", err)
	}

}
