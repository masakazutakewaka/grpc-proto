package main

import (
	"golang.org/x/net/context"
	//"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	coordinatePb "github.com/masakazutakewaka/grpc-proto/coordinate/pb"
	itemPb "github.com/masakazutakewaka/grpc-proto/item/pb"
	userPb "github.com/masakazutakewaka/grpc-proto/user/pb"
)

func run(itemURL string, userURL string, coordinateURL string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := itemPb.RegisterItemServiceHandlerFromEndpoint(ctx, mux, itemURL, opts)
	if err != nil {
		return err
	}
	err = userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userURL, opts)
	if err != nil {
		return err
	}
	err = coordinatePb.RegisterCoordinateServiceHandlerFromEndpoint(ctx, mux, coordinateURL, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	itemURL := os.Getenv("ITEM_URL")
	userURL := os.Getenv("USER_URL")
	coordinateURL := os.Getenv("COORDINATE_URL")

	defer glog.Flush()

	if err := run(itemURL, userURL, coordinateURL); err != nil {
		glog.Fatal(err)
	}
}
