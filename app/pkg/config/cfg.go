package config

import (
	"os"
)

var (

	// SERVER_PORT is the port at which the gRPC server would listen
	SERVER_PORT = "50051"

	// KUBECONFIG env variable to be set if application running outside of k8s cluster (optional)
	KUBECONFIG = os.Getenv("KUBECONFIG")

	// POD_UPDATE_FREQ frequency (secs) at which the pod labels and pod IPs
	// would be updated
	POD_UPDATE_FREQ = 300
)
