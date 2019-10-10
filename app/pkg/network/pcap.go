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

var (
	snapshotlen int32 = 1024
	promiscuous       = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle

	// Will reuse these for each packet
	ethLayer layers.Ethernet
	ipLayer  layers.IPv4
	tcpLayer layers.TCP
)

// Capture captures network traffic and adds pod info into the network
// traffic output
func Capture(iface string) {

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

func Info(iface string, info chan NetworkInfo, stop chan bool) {

	log.Printf("Serving Info from interface '%s'", iface)

	// Open device
	handle, err = pcap.OpenLive(iface, snapshotlen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		select {
		case <-stop:
			log.Printf("Stop network capture...")
			return
		default:
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

					networkInfo := NetworkInfo{Src: srcVal, Dst: dstVal}
					info <- networkInfo

				}
				if layerType == layers.LayerTypeTCP {
					fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
					fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
				}
			}
		}
	}
}
