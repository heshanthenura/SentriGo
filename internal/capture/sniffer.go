package capture

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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
	for packet := range packetSource.Packets() {
		// if netLayer := packet.NetworkLayer(); netLayer != nil {
		// 	src, dst := netLayer.NetworkFlow().Endpoints()
		// 	fmt.Printf("📦 Packet from %v ➜ %v\n", src, dst)
		// }
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		icmpLayer := packet.Layer(layers.LayerTypeICMPv4)

		if ipLayer != nil && icmpLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			icmp, _ := icmpLayer.(*layers.ICMPv4)
			if icmp.TypeCode.Type() == 8 {
				fmt.Printf("ICMP Packet ➜ Type: %d, Code: %d, From: %s ➜ To: %s\n",
					icmp.TypeCode.Type(), icmp.TypeCode.Code(),
					ip.SrcIP, ip.DstIP)
			}

		}
	}
}
