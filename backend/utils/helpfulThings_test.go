package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetAge(t *testing.T) {
	got, err := GetAge(DateUtc(2000, 1, 2), DateUtc(2001, 1, 1))
	require.NoError(t, err)
	require.Equal(t, 0, got)
	got, err = GetAge(DateUtc(2000, 1, 2), DateUtc(2001, 1, 2))
	require.NoError(t, err)
	require.Equal(t, 1, got)

	got, err = GetAge(DateUtc(2000, 1, 2), DateUtc(2000, 1, 1))
	require.Error(t, err)
	require.Empty(t, got)
}

func TestToPtr(t *testing.T) {
	int32Ptr := Int32Ptr(1)
	require.NotNil(t, int32Ptr)
	require.Equal(t, int32(1), *int32Ptr)

	int64Ptr := Int64Ptr(1)
	require.NotNil(t, int64Ptr)
	require.Equal(t, int64(1), *int64Ptr)

	strPtr := StringPtr("abc")
	require.NotNil(t, strPtr)
	require.Equal(t, "abc", *strPtr)

	intPtr := ToPtr(1)
	require.NotNil(t, intPtr)
	require.Equal(t, 1, *intPtr)

	intVal := ToValOrZero(intPtr)
	require.Equal(t, 1, intVal)
	intPtr = nil
	intVal = ToValOrZero(intPtr)
	require.Empty(t, intVal)

	intVal = ToIntOrZero(Float32Ptr(10.1))
	require.Equal(t, 10, intVal)
	var f32Ptr *float32
	intVal = ToIntOrZero(f32Ptr)
	require.Empty(t, intVal)

	f32Val := ToFloat32OrZero(Float64Ptr(10.1))
	require.Equal(t, float32(10.1), f32Val)
	f32Val = ToFloat32OrZero(f32Ptr)
	require.Empty(t, f32Val)

	f64Val := ToFloat64OrZero(IntPtr(10))
	require.Equal(t, 10.0, f64Val)
	f64Val = ToFloat64OrZero(f32Ptr)
	require.Empty(t, f64Val)
}

func TestTimePtr(t *testing.T) {
	t1 := TimePtr(2000, 12, 28, 23, 59, 59, 9000, time.UTC)
	require.NotEmpty(t, t1)
	require.Equal(t, 2000, t1.Year())
	require.Equal(t, 23, t1.Hour())
	require.Equal(t, 9000, t1.Nanosecond())

	t1 = DateUtcPtr(2001, 1, 2)
	require.NotEmpty(t, t1)
	require.Equal(t, 2001, t1.Year())
	require.Equal(t, time.Month(1), t1.Month())
	require.Equal(t, 2, t1.Day())
}
