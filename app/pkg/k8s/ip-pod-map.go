package k8s

import (
	"log"
	"time"

	"github.com/infracloudio/flowfabric/app/pkg/config"
)

// UpdateIPPodMap adds all the pods in all the namespaces into the IPPodMap
func UpdateIPPodMap() error {

	// Get all namespaces
	namespaces, err := GetAllNamespaces()
	if err != nil {
		return err
	}

	// Add all pods in all namespaces to the IPPodMap
	for _, ns := range namespaces {

		// Get all pods in namespaces
		p, err := GetPods(ns)
		if err != nil {
			return err
		}

		// Populate the IPPodMap
		for _, pod := range p.Items {

			var (
				podName   = pod.Name
				podIP     = pod.Status.PodIP
				podLabels = pod.ObjectMeta.Labels
			)

			// Add pod name to the map of pod labels
			podLabels["Name"] = podName
			IPPodMap[podIP] = podLabels
		}
	}

	// Sucessfull
	return nil
}

// PeriodicUpdateIPPodMap periodically updates the IPPodMap
func PeriodicUpdateIPPodMap() {

	// func to update IPPodMap
	update := func() {
		if err := UpdateIPPodMap(); err != nil {
			log.Printf("Failed to update IPPodMap. Error: '%s'", err.Error())
			return
		}
	}

	// Update for the first time on startup
	update()

	// Initialize ticker
	ticker := time.NewTicker(time.Duration(config.POD_UPDATE_FREQ) * time.Second)
	defer ticker.Stop()

	// Periodically update IPPodMap
	for {
		select {
		case <-ticker.C:
			log.Println("Periodic update of IPPodMap...")
			update()
		}
	}
}
