package main

import (
	"golang.org/x/net/context"
	"log"
	"os"

	"github.com/masakazutakewaka/grpc-proto/item"
	"github.com/masakazutakewaka/grpc-proto/user"
)

func main() {
	itemURL := os.Getenv("ITEM_URL")
	userURL := os.Getenv("USER_URL")

	userClient, err := user.NewClient(userURL)
	if err != nil {
		log.Fatal(err)
	}
	defer userClient.Close()

	err = userClient.PostUser(context.Background(), "takewaka")
	if err != nil {
		log.Fatal(err)
	}

	user, err := userClient.GetUser(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	itemClient, err := item.NewClient(itemURL)
	if err != nil {
		log.Fatal(err)
	}
	defer itemClient.Close()

	_, err = itemClient.PostItem(context.Background(), "hat", 1555)
	if err != nil {
		log.Fatal(err)
	}

	item, err := itemClient.GetItem(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(item)
}
