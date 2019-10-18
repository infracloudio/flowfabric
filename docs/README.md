# Flowfabric
Flowfabric helps user intelligently monitor and analyze network traffic in a kubernetes cluster. Flowfabric server needs to be installed on every node of the kubernetes cluster to monitor network traffic from all the pods in the kubernetes cluster.

#### Theory:
- Flowfabric pod is deployed on the host network by setting the `hostNetwork: true` in deployment configuration.
- As the flowfabric pod can access the host network, it identifies the network bridge on which all the pods are connected.
- It also creates a (golang) map of pod names and their respective IP addresses
- When the (gRPC) client sends a request to the (gRPC) server for capturing network traffic, the server captures the network traffic on the previously identified network interface and replaces IP addresses with pod names using the previously created (golang) map. This modified network information is sent back to the client as a gRPC server side stream.
- The golang `gopacket` package is used for monitoring network traffic.
- `libpcapdev` package is a prequisite that needs to be installed for `gopacket` library to work.

#### Build docker image:
1. Install docker for building a docker image
2. Clone this repository onto your local machine.
3. Create a docker image:
	```sh
	docker build -t flowfabric:latest .
	```
4. This step should create a docker image which can be used for deploying containers in a kubernetes setup.

#### Using the flowctl cli:
1. The `flowctl` cli is packaged along with the docker image created.
2. Exec into the flowfabric pod
	```sh
	# ls
	client	flowctl  server
	# ./flowctl -h
	Usage of ./flowctl:
	  -dedup string
	    	(true|false) Deduplicate network information (default "true")
	  -pod string
	    	Name of the pod to monitor network traffic (default "all")
	```
