package detectors

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/heshanthenura/SentriGo/internal/types"
)

type ICMPFloodDetector struct{}

var icmpMap = make(map[string]*types.Tracker)

const (
	ICMPFloodThreshold       = 100
	ICMPFloodDetectionWindow = 5 * time.Second
	ICMPFloodAlertCooldown   = 10 * time.Second
)

func (s *ICMPFloodDetector) Detect(packet gopacket.Packet) {
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	icmp6Layer := packet.Layer(layers.LayerTypeICMPv6)

	if icmpLayer == nil && icmp6Layer == nil {
		return
	}

	var srcIP string
	if ip4Layer := packet.Layer(layers.LayerTypeIPv4); ip4Layer != nil {
		ip4 := ip4Layer.(*layers.IPv4)
		srcIP = ip4.SrcIP.String()
	} else if ip6Layer := packet.Layer(layers.LayerTypeIPv6); ip6Layer != nil {
		ip6 := ip6Layer.(*layers.IPv6)
		srcIP = ip6.SrcIP.String()
	} else {
		return
	}

	now := time.Now()

	tracker, exists := icmpMap[srcIP]
	if !exists {
		icmpMap[srcIP] = &types.Tracker{
			IP:        srcIP,
			StartTime: now,
			Count:     1,
			LastAlert: time.Time{},
		}
		return
	}

	if now.Sub(tracker.StartTime) <= ICMPFloodDetectionWindow {
		tracker.Count++
		if tracker.Count > ICMPFloodThreshold {
			if now.Sub(tracker.LastAlert) > ICMPFloodAlertCooldown {
				fmt.Printf("\033[31m[ICMP Flood]\033[0m Detected from %s: %d ICMP packets in %v\n", srcIP, tracker.Count, ICMPFloodDetectionWindow)
				tracker.LastAlert = now
			}
		}
	} else {
		tracker.Count = 1
		tracker.StartTime = now
	}
}
