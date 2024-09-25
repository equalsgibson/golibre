package main

import (
	"context"
	"log"
	"os"

	"github.com/equalsgibson/golibre/golibre"
)

func main() {
	// Set up a new golibre service
	ctx := context.Background()
	service := golibre.NewService(
		"api.libreview.io",
		golibre.Authentication{
			Email:    os.Getenv("EMAIL"),    // Your email address
			Password: os.Getenv("PASSWORD"), // Your password
		},
	)

	connections, err := service.Connection().GetAllConnectionData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print a count of all the patients that you are connected to, with a list of patient IDs
	log.Printf("You have %d patients that are sharing their data with you.\n\n", len(connections))

	for i, connection := range connections {
		log.Printf("\tPatient %d: %s\n", i, connection.PatientID)
	}
}
