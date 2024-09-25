package main

import (
	"context"
	"fmt"
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
	fmt.Printf("You have %d patients that are sharing their data with you.\n", len(connections))

	for i, connection := range connections {
		fmt.Printf("\t-> Patient %d: ID: %s\n", i+1, connection.PatientID)
	}
}
