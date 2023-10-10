package utils

import "time"

func Float32Ptr(a float32) *float32 {
	return &a
}

func Float64Ptr(a float64) *float64 {
	return &a
}

func IntPtr(a int) *int {
	return &a
}

func Int32Ptr(a int32) *int32 {
	return &a
}

func Int64Ptr(a int64) *int64 {
	return &a
}

func StringPtr(a string) *string {
	return &a
}

func TimePtr(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &t
}

func DateUtc(year int, month time.Month, day int) time.Time {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return t
}

func DateUtcPtr(year int, month time.Month, day int) *time.Time {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return &t
}
