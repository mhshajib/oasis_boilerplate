package utils

import "time"

// Convert UTC to Dhaka Timezone
func ConvertToDhakaTime(utcTime time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Dhaka")
	return utcTime.In(loc)
}

// Convert Dhaka Timezone to UTC
func ConvertToUTCTime(dhakaTime time.Time) time.Time {
	return dhakaTime.UTC()
}

func DateToHumanReadableFormat(time time.Time) string {
	return time.Format("2006-01-02 03:04:05 PM")
}

func IsExpired(expirationTime time.Time) bool {
	// Get the current time
	currentTime := time.Now()
	// Check if the current time is after the expiration time
	return currentTime.After(expirationTime)
}
