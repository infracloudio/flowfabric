package server

import (
	"log"
	"net"

	"github.com/infracloudio/flowfabric/app/pkg/config"
	"github.com/infracloudio/flowfabric/app/pkg/network"
	pb "github.com/infracloudio/flowfabric/app/pkg/proto"
	"github.com/infracloudio/flowfabric/app/pkg/runtime"
	"google.golang.org/grpc"
)

var (
	port = ":" + config.SERVER_PORT
)

// server struct is used to implement ff.NetworkCaptureServer
type server struct{}

// Capture implements ff.ff.NetworkCaptureServer.Capture
func (s *server) Capture(req *pb.CaptureRequest, stream pb.NetworkCapture_CaptureServer) error {

	// sync channels
	var (
		networkInfo = make(chan network.NetworkInfo)
		stop        = make(chan bool)
	)

	// Initiate network capture
	go network.Info(req.Pod, runtime.Iface, networkInfo, stop)

	// Send back stream response
	for {
		ni := <-networkInfo

		// create response
		r := &pb.CaptureResponse{
			SrcIp:   ni.SrcIP,
			DstIp:   ni.DstIP,
			SrcPort: ni.SrcPort,
			DstPort: ni.DstPort,
		}
		err := stream.Send(r)
		if err != nil {
			log.Printf("Failed to send stream to client. Error: '%s'", err.Error())
			stop <- true
			return err
		}
	}
}

// StartServer starts the gRPC server
func StartServer() {

	log.Printf("Starting server at port '%s'", port)

	// create a tcp server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen at port '%s'. Error: '%s'", port, err.Error())
	}
	log.Printf("Listening at port '%s'", port)

	// create a grpc server
	s := grpc.NewServer()

	// register grpc server
	pb.RegisterNetworkCaptureServer(s, &server{})

	// serve requests
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve request. Error: '%s'", err.Error())
	}
}
