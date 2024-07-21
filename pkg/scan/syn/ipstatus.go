package syn

import "time"

type IpStatus struct {
	ReceivedPort map[uint16]struct{}
	LastTime     time.Time
}
