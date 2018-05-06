package lib

import "time"

func GetTimeFromTimezone(timezone string, localTime *time.Time) time.Time {
	location, _ := time.LoadLocation(timezone)
	return localTime.In(location)
}
