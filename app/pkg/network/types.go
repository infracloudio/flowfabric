package network

import (
	"regexp"
)

var (

	// CIDRPattern - pattern to validate CIDRs
	CIDRPattern = regexp.MustCompile("((\\d){1,3}\\.){3}(\\d){1,3}\\/(\\d){1,3}")

	// IfaceCIDRMap - map of network interface names and associated CIDR
	IfaceCIDRMap = make(map[string]string)
)
