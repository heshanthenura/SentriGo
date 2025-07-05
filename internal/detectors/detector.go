package detectors

import "github.com/google/gopacket"

type Detector interface {
	Detect(packet gopacket.Packet)
}
