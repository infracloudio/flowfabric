package k8s

import (
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods returns all the pods present in a namespace
func GetPods(ns string) (p *corev1.PodList, err error) {

	// List pods
	p, err = ClientSet.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to list pods in namespace '%s'. Error: '%s'", ns, err.Error())
		log.Println(errMsg)
		return p, fmt.Errorf(errMsg)
	}

	// Successful
	return p, nil
}
