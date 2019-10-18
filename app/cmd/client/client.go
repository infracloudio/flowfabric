package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/infracloudio/flowfabric/app/pkg/config"
	pb "github.com/infracloudio/flowfabric/app/pkg/proto"
	"google.golang.org/grpc"
)

var (
	port    = ":" + config.SERVER_PORT
	address = "localhost"
)

// setupServerConn sets up a server connection to the gRPC server
func setupServerConn(url string) *grpc.ClientConn {

	// setup a connection to the server
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the gRPC server. Error: '%s'", err.Error())
	}
	return conn
}

// requestNetworkInfo requests server for network information
func requestNetworkInfo(pod string, dedup bool, conn *grpc.ClientConn) {

	var (
		ctx          = context.Background()
		c            = pb.NewNetworkCaptureClient(conn)
		networkCache = make(map[string]bool)
	)

	// request stream
	s, err := c.Capture(ctx, &pb.CaptureRequest{Pod: pod})
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

		// Deduplicate network data
		if dedup {

			// Check cache for deduplication
			possibleVal1 := fmt.Sprintf("%s-%s", ni.SrcIp, ni.DstIp)
			possibleVal2 := fmt.Sprintf("%s-%s", ni.DstIp, ni.SrcIp)

			// If value already present in network cache go ahead, if not add
			// it to network cache
			if networkCache[possibleVal1] || networkCache[possibleVal2] {
				continue
			} else {
				networkCache[possibleVal1] = true
			}
		}

		// Filter pod info
		if pod == "all" || pod == ni.SrcIp || pod == ni.DstIp {
			// Foramtted output of network info
			fmt.Printf("%-5s %s -> %s\n", "IPv4:", ni.SrcIp, ni.DstIp)
			fmt.Printf("%-5s %s -> %s\n", "Port:", ni.SrcPort, ni.DstPort)
			fmt.Println()
		} else {
			// Filter out undesired values
			continue
		}
	}
}

func main() {

	// Parse command line options
	pod := flag.String("pod", "all", "Name of the pod to monitor network traffic")
	d := flag.String("dedup", "true", "(true|false) Deduplicate network information")
	flag.Parse()

	// Convert dedup string to bool, this conversion is necessary due to a
	// small bug in go flags where all bool flags need to be set with "=".
	// Hence user convenience, taking dedup as string.
	dedup, err := strconv.ParseBool(*d)
	if err != nil {
		log.Fatalf("Incorrect value for cli option '-dedup'. Error: '%s'", err.Error())
	}

	// Setup server connection
	url := address + port
	conn := setupServerConn(url)
	defer conn.Close()

	// Request network information
	requestNetworkInfo(*pod, dedup, conn)
}
