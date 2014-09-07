package logger

import "time"

// Default Timing, can be mocked in tests
var defaultTiming timingI = timing{}

// Interface for Timing
type timingI interface {
	getCurrentTime() time.Time
}

// Timing is getting current time
type timing struct {
}

// Getting current time
func (t timing) getCurrentTime() time.Time {
	return time.Now()
}
