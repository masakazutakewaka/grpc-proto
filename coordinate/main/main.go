package main

import (
	"log"
	"os"

	"github.com/masakazutakewaka/grpc-proto/coordinate"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	itemURL := os.Getenv("ITEM_URL")
	userURL := os.Getenv("USER_URL")

	repo, err := coordinate.NewPostgresRepository(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen to port 8080 ...")
	log.Fatal(coordinate.ListenGRPC(repo, itemURL, userURL, 8080))
}
