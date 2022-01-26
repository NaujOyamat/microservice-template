package main

import (
	"log"

	"Github.com/NaujOyamat/microservice-template/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
