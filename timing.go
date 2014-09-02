package logger

import "time"

var defaultTiming timingI = timing{}

type timingI interface {
	getCurrentTime() time.Time
}

type timing struct {
}

func (t timing) getCurrentTime() time.Time {
	return time.Now()
}
