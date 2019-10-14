package k8s

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/infracloudio/flowfabric/app/pkg/config"
)

// CreateClientSet creates kubernetes client. If the application is running
// inside a k8s cluster it does not need kubeconfig. But, if deployed outside of
// k8s cluster then "KUBECONFIG" env var needs to be set OR kubeconfig file
// needs to be present at default kubeconfig path i.e "$HOME/.kube/config"
func CreateClientSet() (*kubernetes.Clientset, error) {

	// Check for incluster config
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {

		log.Println("Application NOT running inside a k8s cluster. Figuring out kubeconfig of the intended cluster...")

		// Check KUBECONFIG variable
		kubeconfigPath := config.KUBECONFIG
		if kubeconfigPath == "" {

			log.Println("'KUBECONFIG' env variable not set. Expecting kubeconfig at default path i.e $HOME/.kube/config")

			// Default KUBECONFIG path
			kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
		}

		// Build k8s config from kubeconfig file path
		kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			log.Printf("Failed to build config from flags. Error: '%s'", err.Error())
			return nil, err
		}

		// Create clientset
		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			log.Printf("Failed to create client set. Error: '%s'", err.Error())
			return nil, err
		}

		// Successfull
		return clientset, nil

	}

	log.Println("Application running inside a k8s cluster")

	// Create client set
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Printf("Failed to create client set. Error: '%s'", err.Error())
		return nil, err
	}

	// Successfull
	return clientset, nil
}
