package logger

import "time"

var defTimingOriginal = defaultTiming

type mockedTiming struct {
}

func (m mockedTiming) getCurrentTime() time.Time {
	location, _ := time.LoadLocation("UTC")
	return time.Date(1985, 1, 2, 3, 4, 5, 6, location)
}
