package common

import "time"

type ICMPFloodRecord struct {
	Source string
	Count  int
	Start  time.Time
	End    time.Time
	Warned bool
}
