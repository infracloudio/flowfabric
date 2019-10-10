package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/infracloudio/flowfabric/app/pkg/config"
	pb "github.com/infracloudio/flowfabric/app/pkg/proto"
	"google.golang.org/grpc"
)

var (
	port    = ":" + config.SERVER_PORT
	address = "localhost"
)

func main() {

	// pod name
	pod := "default"
	if len(os.Args) > 1 {
		pod = os.Args[1]
	}

	// setup a connection to the server
	url := address + port
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the gRPC server. Error: '%s'", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()
	c := pb.NewNetworkCaptureClient(conn)

	// request stream
	s, err := c.Capture(ctx, &pb.CaptureRequest{Pod: pod})
	if err != nil {
		log.Fatalf("Failed to create stream handler. Error: '%s'", err.Error())
	}

	// Receive stream
	for {
		r, err := s.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive response. Error: '%s'", err.Error())
		}
		fmt.Println("IPv4: ", r.Src, "->", r.Dst)
	}
}
