package network

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// UpdateIfaceCIDRMap adds all the interfaces and respective CIDRs to the IfaceCIDRMap
func UpdateIfaceCIDRMap() error {

	// Get all the network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get network Interfaces. Error: '%s'\n", err.Error())
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	// Add all the network interfaces and CIDR to the IfaceCIDRMap
	for _, i := range ifaces {

		// If network interface is not UP, continue ahead!
		if !strings.Contains(i.Flags.String(), "up") {
			// log.Printf("Interface '%s' is not UP", i.Name)
			continue
		}

		// Get all the interface addresses
		addrs, err := i.Addrs()
		if err != nil {
			errMsg := fmt.Sprintf("Failed to fetch interface addresses. Error: '%s'", err.Error())
			log.Println(errMsg)
			return fmt.Errorf(errMsg)
		}

		// Add valid CIDRs to the IfaceCIDRMap
		for _, a := range addrs {
			if CIDRPattern.MatchString(a.String()) {
				IfaceCIDRMap[i.Name] = a.String()
			}
		}
	}

	// Successfull
	return nil
}
