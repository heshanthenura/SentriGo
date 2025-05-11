package detection

import (
	"log"
	"time"

	"github.com/google/gopacket/layers"
	"github.com/heshanthenura/SentriGo/internal/common"
)

var ICMPFloodRecords = make(map[string]*common.ICMPFloodRecord)

func ICMPFloodDetection(ip *layers.IPv4, icmp *layers.ICMPv4) {
	AddICMPFloodRecord(ip.SrcIP.String())
}

func AddICMPFloodRecord(ip string) {
	if _, exists := ICMPFloodRecords[ip]; !exists {
		ICMPFloodRecords[ip] = &common.ICMPFloodRecord{
			Source: ip,
			Count:  1,
			Start:  time.Now(),
			Warned: false,
		}
	} else {
		ICMPFloodRecords[ip].Count++
		ICMPFloodRecords[ip].End = time.Now()
	}
	packet, _ := FindICMPFloodRecord(ip)
	duration := GetDurationSeconds(packet.Start, packet.End)
	log.Printf("[ ICMP Flood Detection ] Source: %s, Count: %d, Start: %v, End: %v, Duration: %.2f seconds\n",
		packet.Source, packet.Count, packet.Start, packet.End, duration)
	DetectFlood(packet)
}

func FindICMPFloodRecord(ip string) (*common.ICMPFloodRecord, bool) {
	record, exists := ICMPFloodRecords[ip]
	return record, exists
}

func GetDurationSeconds(start, end time.Time) float64 {
	var dif = end.Sub(start).Seconds()
	if dif < 0 {
		return 0
	}
	return dif
}

func DetectFlood(record *common.ICMPFloodRecord) {
	duration := GetDurationSeconds(record.Start, record.End)

	if record.Count >= 10 && duration <= 30 {
		if !record.Warned {
			log.Println("Flood Detected")
			record.Warned = true
		}
		record.Count = 0
		record.Start = time.Now()
		record.End = time.Now()
		record.Warned = false
	} else if duration > 30 && record.Count < 10 {
		log.Printf("Removing record for %s: inactive or normal activity.\n", record.Source)
		delete(ICMPFloodRecords, record.Source)
	}
}
