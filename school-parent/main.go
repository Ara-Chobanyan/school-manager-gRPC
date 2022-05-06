package main

import (
	"log"

	"github.com/Ara-Chobanyan/school-manager/school-parent/api"
	"github.com/Ara-Chobanyan/school-manager/school-parent/app"
	"github.com/Ara-Chobanyan/school-manager/school-parent/repository/postgres"
)

func main() {
	// initialized the repo logic
	repo, err := postgres.NewPostgresRepository()

	if err != nil {
		log.Fatalf("Something went wrong NewPostgresRepository: %v", err)
	}

	// Inject our repo inside of the services interface
	service := app.NewParentService(repo)

	// Open up our gRPC interface and inject our services inside of it
	rpc := api.NewAdapter(service)

	// Start the rpc server
	err = rpc.Run()
	// check for any errors
	if err != nil {
		log.Fatalf("Could not run: %v", err)
	}

}
