package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/infracloudio/flowfabric/app/pkg/config"
	pb "github.com/infracloudio/flowfabric/app/pkg/ff"
	"google.golang.org/grpc"
)

var (
	port    = ":" + config.SERVER_PORT
	address = "localhost"
)

func main() {

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
	s, err := c.Capture(ctx)
	if err != nil {
		log.Fatalf("Failed to create stream handler. Error: '%s'", err.Error())
	}

	waitc := make(chan struct{})

	// Receive stream
	go func() {
		for {
			r, err := s.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive response. Error: '%s'", err.Error())
			}
			fmt.Println("IPv4: ", r.Src, "->", r.Dst)
		}
	}()

	// Send stream till an interrupt is received
	signalChan := make(chan os.Signal, 1)
	for {
		select {
		case <-signalChan:
			log.Printf("Closing client...")
			if err := s.Send(&pb.CaptureInfo{Stop: true}); err != nil {
				log.Fatalf("Failed to send STOP stream message to server. Error: '%s'", err.Error())
			}
		default:
			if err := s.Send(&pb.CaptureInfo{Stop: false}); err != nil {
				log.Fatalf("Failed to send stream message to server. Error: '%s'", err.Error())
			}
		}
	}

	s.CloseSend()
	<-waitc
}
