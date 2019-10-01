package k8s

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetAllNamespaces returns a slice of names of namespaces in the k8s cluster
func GetAllNamespaces() ([]string, error) {

	var namespaces []string

	// List namespaces
	ns, err := ClientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to list namespaces. Error: '%s'", err.Error())
		log.Println(errMsg)
		return namespaces, fmt.Errorf(errMsg)
	}

	// Add names for "Active" namespaces in a slice
	for _, namespace := range ns.Items {
		if namespace.Status.Phase == "Active" {
			namespaces = append(namespaces, namespace.ObjectMeta.Name)
		}
	}

	// Successfull
	return namespaces, nil
}
