package main

import (
	"golang.org/x/net/context"
	"log"
	"os"

	"github.com/masakazutakewaka/grpc-proto/coordinate"
	"github.com/masakazutakewaka/grpc-proto/item"
	"github.com/masakazutakewaka/grpc-proto/user"
)

func main() {
	itemURL := os.Getenv("ITEM_URL")
	userURL := os.Getenv("USER_URL")
	coordinateURL := os.Getenv("COORDINATE_URL")

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

	_, err = itemClient.PostItem(context.Background(), "shoes", 2255)
	if err != nil {
		log.Fatal(err)
	}

	items, err := itemClient.GetItems(context.Background(), []int32{1, 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(items)

	coordinateClient, err := coordinate.NewClient(coordinateURL)
	if err != nil {
		log.Fatal(err)
	}
	defer coordinateClient.Close()

	err = coordinateClient.PostCoordinate(context.Background(), 1, []int32{1, 2})
	if err != nil {
		log.Fatal(err)
	}

	coordinate, err := coordinateClient.GetCoordinatesByUser(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(coordinate)
}
