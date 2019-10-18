package network

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/infracloudio/flowfabric/app/pkg/k8s"
)

// Capture captures network traffic and adds pod info into the network
// traffic output
func Capture(iface string) {

	var (
		snapshotlen int32 = 1024
		promiscuous       = false
		err         error
		timeout     time.Duration = -1 * time.Second
		handle      *pcap.Handle

		// Will reuse these for each packet
		ethLayer layers.Ethernet
		ipLayer  layers.IPv4
		tcpLayer layers.TCP
	)

	// Open device
	handle, err = pcap.OpenLive(iface, snapshotlen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&ipLayer,
			&tcpLayer,
		)
		foundLayerTypes := []gopacket.LayerType{}

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			fmt.Println("Trouble decoding layers: ", err)
		}

		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeIPv4 {

				srcVal := ipLayer.SrcIP.String()
				dstVal := ipLayer.DstIP.String()

				// Add pod info
				if _, ok := k8s.IPPodMap[srcVal]; ok {
					srcVal = k8s.IPPodMap[srcVal]["Name"]
				}

				if _, ok := k8s.IPPodMap[dstVal]; ok {
					dstVal = k8s.IPPodMap[dstVal]["Name"]
				}

				fmt.Println("IPv4: ", srcVal, "->", dstVal)
			}
			if layerType == layers.LayerTypeTCP {
				fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
				fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
			}
		}
	}
}

func Info(pod, iface string, info chan NetworkInfo, stop chan bool) {

	log.Printf("Monitoring network traffic for '%s' pod(s) on interface '%s'", pod, iface)

	var (
		snapshotlen int32         = 1024
		promiscuous bool          = false
		timeout     time.Duration = -1 * time.Second

		err    error
		handle *pcap.Handle
	)

	// Open device
	handle, err = pcap.OpenLive(iface, snapshotlen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Read packets from packet source
	for packet := range packetSource.Packets() {
		select {
		case <-stop:
			log.Printf("Stopping network capture for '%s' pods(s) on interface '%s'", pod, iface)
			return
		default:

			var ni NetworkInfo

			// Check if the packet has IPv4 layer
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {

				ip, _ := ipLayer.(*layers.IPv4)

				srcVal := ip.SrcIP.String()
				dstVal := ip.DstIP.String()

				// Add pod info
				if _, ok := k8s.IPPodMap[srcVal]; ok {
					srcVal = k8s.IPPodMap[srcVal]["Name"]
				}

				if _, ok := k8s.IPPodMap[dstVal]; ok {
					dstVal = k8s.IPPodMap[dstVal]["Name"]
				}

				ni.SrcIP = srcVal
				ni.DstIP = dstVal

			} else {
				// IP layer not present
				continue
			}

			// Check if the packet has TCP layer
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)

				ni.SrcPort = tcp.SrcPort.String()
				ni.DstPort = tcp.DstPort.String()
			}
			info <- ni
		}
	}
}
