package main

import (
	"jurassic-park-api/api"
	"jurassic-park-api/data/postgres"
	"log"
)

func main() {
	store, err := postgres.NewStore()
	if err != nil {
		log.Printf("error creating new postgres store: %s\n", err)
		return
	}

	r := api.SetupRouter(store)
	err = r.Run(":8888")
	if err != nil {
		log.Printf("error starting server: %s\n", err)
		return
	}
}
