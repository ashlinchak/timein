package lib

import (
	"time"

	"4d63.com/tz"
)

func GetTimeFromTimezone(timezone string, localTime *time.Time) time.Time {
	// https://github.com/golang/go/issues/21881
	location, _ := tz.LoadLocation(timezone)

	return localTime.In(location)
}
