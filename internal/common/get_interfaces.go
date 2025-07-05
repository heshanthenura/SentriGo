package common

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func GetInterfaces() []pcap.Interface {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Error finding devices: %v", err)
	}

	for i, device := range devices {
		fmt.Printf("[%d] %s -- %s\n", i, device.Name, device.Description)
	}

	return devices
}
