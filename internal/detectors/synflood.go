package detectors

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/heshanthenura/SentriGo/internal/types"
)

type SynFloodDetector struct{}

var synMap = make(map[string]*types.Tracker)

const (
	SYNFloodThreshold       = 30
	SYNFloodDetectionWindow = 1 * time.Second
	SYNFloodAlertCooldown   = 5 * time.Second
)

func (s *SynFloodDetector) Detect(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)

	if ipLayer == nil || tcpLayer == nil {
		return
	}

	ip, _ := ipLayer.(*layers.IPv4)
	tcp, _ := tcpLayer.(*layers.TCP)

	if tcp.SYN && !tcp.ACK {
		srcIP := ip.SrcIP.String()
		tracker, exists := synMap[srcIP]
		now := time.Now()
		if !exists {
			synMap[srcIP] = &types.Tracker{
				IP:        srcIP,
				StartTime: now,
				Count:     1,
				LastAlert: time.Time{},
			}
			return
		}
		if now.Sub(tracker.StartTime) <= SYNFloodDetectionWindow {
			tracker.Count++
			if tracker.Count > SYNFloodThreshold {
				if now.Sub(tracker.LastAlert) > SYNFloodAlertCooldown {
					fmt.Printf("\033[36m[SYN Flood]\033[0m Detected from %s: %d SYNs in %v\n", srcIP, tracker.Count, SYNFloodDetectionWindow)
					tracker.LastAlert = now
				}
			}
		} else {
			tracker.Count = 1
			tracker.StartTime = now
		}
	}
}
