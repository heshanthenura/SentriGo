package common

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"github.com/heshanthenura/SentriGo/internal/detectors"
)

func SniffInterface(device pcap.Interface) {
	fmt.Println("Sniffing on:", device.Description)

	handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)

	if err != nil {
		log.Fatalf("Error opening device: %v", err)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	detectorList := []detectors.Detector{
		&detectors.SynFloodDetector{},
		&detectors.ICMPFloodDetector{},
	}

	for packet := range packetSource.Packets() {
		for _, d := range detectorList {
			d.Detect(packet)
		}
	}
}
