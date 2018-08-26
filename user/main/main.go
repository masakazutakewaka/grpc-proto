package main

import (
	"log"
	"os"

	"github.com/masakazutakewaka/grpc-proto/user"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")

	repo, err := user.NewPostgresRepository(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen to port 8080 ...")
	log.Fatal(user.ListenGRPC(repo, 8080))
}
