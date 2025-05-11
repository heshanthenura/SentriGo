package capture

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/heshanthenura/SentriGo/internal/detection"
)

var (
	started  bool
	stopChan chan bool
)

func GetAvailableDevices() []pcap.Interface {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal("Failed to list devices:", err)
	}

	if len(devices) == 0 {
		log.Fatal("No network interfaces found.")
	}

	return devices
}

func SniffInterface(interfaceName string) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", interfaceName, err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("📡 Sniffing started on", interfaceName)

	for {
		select {
		case <-stopChan:
			fmt.Println("🛑 Sniffing stopped")
			return
		case packet := <-packetSource.Packets():
			if packet == nil {
				continue
			}
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			icmpLayer := packet.Layer(layers.LayerTypeICMPv4)

			if ipLayer != nil && icmpLayer != nil {
				icmp, _ := icmpLayer.(*layers.ICMPv4)
				if icmp.TypeCode.Type() == 8 {
					detection.ICMPFloodDetection(ipLayer.(*layers.IPv4), icmp)
				}
			}
		}
	}
}

func ToggleSniffing(interfaceName string) {
	if started {
		stopChan <- true
		started = false
	} else {
		stopChan = make(chan bool)
		started = true
		go SniffInterface(interfaceName)
	}
}
