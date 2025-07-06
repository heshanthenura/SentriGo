package detectors

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/heshanthenura/SentriGo/internal/types"
)

type NMAPScanDetector struct{}

var nmapMap = make(map[string]*types.NMAPTracker)

const (
	NMAPPortThreshold       = 5
	NMAPScanDetectionWindow = 3 * time.Second
	NMAPScanAlertCooldown   = 10 * time.Second
)

func (s *NMAPScanDetector) Detect(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)

	if ipLayer == nil || tcpLayer == nil {
		return
	}

	ip := ipLayer.(*layers.IPv4)
	tcp := tcpLayer.(*layers.TCP)

	srcIP := ip.SrcIP.String()
	dstPort := uint16(tcp.DstPort)
	now := time.Now()

	tracker, exists := nmapMap[srcIP]
	if !exists {
		nmapMap[srcIP] = &types.NMAPTracker{
			IP:        srcIP,
			StartTime: now,
			Count:     1,
			LastAlert: time.Time{},
			Ports:     map[uint16]bool{dstPort: true},
		}
		return
	}

	if now.Sub(tracker.StartTime) > NMAPScanDetectionWindow {
		tracker.StartTime = now
		tracker.Ports = map[uint16]bool{dstPort: true}
		tracker.Count = 1
		return
	}

	tracker.Ports[dstPort] = true
	tracker.Count++

	if len(tracker.Ports) > NMAPPortThreshold {
		if now.Sub(tracker.LastAlert) > NMAPScanAlertCooldown {
			fmt.Printf("\033[35m[Nmap Scan]\033[0m Detected from %s: scanned %d unique ports in %v\n",
				srcIP, len(tracker.Ports), NMAPScanDetectionWindow)
			tracker.LastAlert = now
		}
	}
}
