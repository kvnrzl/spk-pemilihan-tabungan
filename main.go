package main

import (
	"log"
)

func main() {
	r := InitServer()
	log.Fatal(r.Run(":8080"))
}
