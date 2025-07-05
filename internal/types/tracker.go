package types

import "time"

type Tracker struct {
	IP        string
	StartTime time.Time
	Count     int
	LastAlert time.Time
}
