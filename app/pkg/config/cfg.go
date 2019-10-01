package config

import (
	"os"
)

var (

	// KUBECONFIG env variable to be set if application running outside of k8s cluster (optional)
	KUBECONFIG = os.Getenv("KUBECONFIG")

	// POD_UPDATE_FREQ frequency (secs) at which the pod labels and pod IPs
	// would be updated
	POD_UPDATE_FREQ = 300
)
