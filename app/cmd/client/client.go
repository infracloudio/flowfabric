package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/infracloudio/flowfabric/app/pkg/config"
	pb "github.com/infracloudio/flowfabric/app/pkg/proto"
	"google.golang.org/grpc"
)

var (
	port    = ":" + config.SERVER_PORT
	address = "localhost"
)

func main() {

	// Parse command line options
	pod := flag.String("pod", "all", "Name of the pod to monitor network traffic")
	flag.Parse()

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
	s, err := c.Capture(ctx, &pb.CaptureRequest{Pod: *pod})
	if err != nil {
		log.Fatalf("Failed to create stream handler. Error: '%s'", err.Error())
	}

	// Receive stream
	for {
		ni, err := s.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive response. Error: '%s'", err.Error())
		}
		fmt.Printf("%-5s %s -> %s\n", "IPv4:", ni.SrcIp, ni.DstIp)
		fmt.Printf("%-5s %s -> %s\n", "Port:", ni.SrcPort, ni.DstPort)
		fmt.Println()
	}
}
