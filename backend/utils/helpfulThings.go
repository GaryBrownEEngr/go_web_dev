package utils

import (
	"fmt"
	"time"
)

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

func ToPtr[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool | time.Time](in T) *T {
	return &in
}

// Return the underlying value or zero
func ToValOrZero[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool | time.Time](in *T) T {
	if in == nil {
		var ret T
		return ret
	}
	return *in
}

func ToIntOrZero[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](in *T) int {
	if in == nil {
		return 0
	}
	return int(*in)
}

func ToFloat32OrZero[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](in *T) float32 {
	if in == nil {
		return 0
	}
	return float32(*in)
}

func ToFloat64OrZero[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](in *T) float64 {
	if in == nil {
		return 0
	}
	return float64(*in)
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

func GetAge(dateOfBirth, compareDate time.Time) (int, error) {
	// Set both times to UTC
	dateOfBirth = dateOfBirth.UTC()
	compareDate = compareDate.UTC()

	// Only use year, month, day
	aYear, aMonth, aDay := dateOfBirth.Date()
	dateOfBirth = time.Date(aYear, aMonth, aDay, 0, 0, 0, 0, time.UTC)
	bYear, bMonth, bDay := compareDate.Date()
	compareDate = time.Date(bYear, bMonth, bDay, 0, 0, 0, 0, time.UTC)

	if compareDate.Before(dateOfBirth) {
		return 0, fmt.Errorf("Invalid, negative age")
	}

	age := bYear - aYear

	// Check if a full year hasn't passed yet for the last year
	anniversary := dateOfBirth.AddDate(age, 0, 0)
	if anniversary.After(compareDate) {
		age--
	}

	return age, nil
}
