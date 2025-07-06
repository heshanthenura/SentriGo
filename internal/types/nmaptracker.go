package types

import "time"

type NMAPTracker struct {
	IP        string
	StartTime time.Time
	Count     int
	LastAlert time.Time
	Ports     map[uint16]bool
}
