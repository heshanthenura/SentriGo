<!-- # SentriGo -->

<img src="./banner.png">

## Lightweight Intrusion Detection System (IDS) for Windows

🚀 **SentriGo** is a lightweight, real-time packet sniffing tool and Intrusion Detection System (IDS) built using Go and the `gopacket` library.

## 📝 Description

SentriGo listens to a selected network interface and analyzes packets in real-time. It's designed for educational and home-lab usage to understand basic network intrusion detection techniques.

Currently, it includes:

- SYN flood detection
- ICMP flood detection
- NMAP scan detection

## 🔍 How to Add a New Detection

To create a new detection module, follow these steps:

### 1. Create a New Detector File

Create a new file inside the `internal/detectors` directory. Name the file after your detection logic, e.g., `portscan.go`.
Add the following boilerplate code:

```go
package detectors

import "github.com/google/gopacket"

type <DetectionName> struct{}

func (d *<DetectionName>) Detect(packet gopacket.Packet) {
	// Implement your detection logic here
}
```

Replace `<DetectionName>` with the actual name of your detector (e.g., `PortScanDetector`).

Inside the `Detect` function, you can access various packet layers and implement your custom intrusion detection logic using the `gopacket` library.

### 2. Register Your Detector

Open the file `internal/common/sniffer.go`, and locate the `detectorList` initialization.

Add your new detector to the list:

```go
detectorList := []detectors.Detector{
	&detectors.SynFloodDetector{},
	&detectors.ICMPFloodDetector{},
	&detectors.NMAPScanDetector{},
	&detectors.<DetectionName>{}, // Add your new detector here
}
```

Again, replace `<DetectionName>` with the name of your struct.

✅ Your new detection module is now active and will be called for every captured packet.

---

## ⚙️ Prerequisites

### 🪟 OS Support:

✅ **Tested on Windows**

✅ **Linux**

### 📥 Requirements:

1. [Go Programming Language](https://golang.org/dl/) (v1.18+)
2. [Npcap](https://npcap.com/#download) (must be installed)

> 📌 **Npcap** is a packet capture driver required to access network interfaces on Windows.
