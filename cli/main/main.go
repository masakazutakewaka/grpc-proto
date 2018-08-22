package main

import (
	"golang.org/x/net/context"
	"log"
	"os"

	"github.com/masakazutakewaka/grpc-proto/item"
)

func main() {
	itemURL := os.Getenv("ITEM_URL")

	client, err := item.NewClient(itemURL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.PostItem(context.Background(), "hat", 1555)
	if err != nil {
		log.Fatal(err)
	}

	item, err := client.GetItem(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(item)
}
