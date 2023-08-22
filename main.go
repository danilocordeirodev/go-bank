package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	// server := NewAPIServer(":3000")
	// server.Run()
}
