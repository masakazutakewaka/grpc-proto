package item

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/masakazutakewaka/grpc-proto/item/pb"
)

type Client struct {
	conn *grpc.Connection
}

func NewClient(url string) (*Client, error) {
}

func GetItem() {
}

func GetItems() {
}

func PostItem() {
}
