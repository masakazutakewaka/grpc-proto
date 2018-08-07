package main

import (
	"log"

	"wear-proto/item"
)

func main() {
	dbURL = ""
	repo, err := item.NewPostgresRepository(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen to port 8080 ...")
	server := itemServer{}
	log.Fatal(item.ListenGRPC(repo, 8080))
}
