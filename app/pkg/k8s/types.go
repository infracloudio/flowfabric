package k8s

import (
	"k8s.io/client-go/kubernetes"
)

var (

	// ClientSet - kubernetes client
	ClientSet *kubernetes.Clientset

	// IPPodMap - Nested map of pod IPs and map of pod labels (which includes pod name)
	IPPodMap = make(map[string]map[string]string)
)
