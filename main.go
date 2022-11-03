package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	r := InitServer()
	log.Fatal(r.Run(":8080"))
}
