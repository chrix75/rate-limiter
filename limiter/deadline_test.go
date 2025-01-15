package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAllowCallInCurrentDay(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	dayLimiter := NewDayLimiter(refTime, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := dayLimiter.Allow(refTime, "client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallInCurrentDay(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	dayLimiter := NewDayLimiter(refTime, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow(refTime, "client_1")
	allowed := dayLimiter.Allow(refTime, "client_1")

	// then
	assert.False(t, allowed)
}

func TestResetBucketByDay(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	dayLimiter := NewDayLimiter(refTime, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow(refTime, "client_1")

	nextCallTime := refTime.Add(24 * time.Hour)
	allowed := dayLimiter.Allow(nextCallTime, "client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallAfterResetByDay(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	dayLimiter := NewDayLimiter(refTime, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow(refTime, "client_1")

	nextCallTime := refTime.Add(time.Minute)
	_ = dayLimiter.Allow(nextCallTime, "client_1")
	allowed := dayLimiter.Allow(nextCallTime, "client_1")

	// then
	assert.False(t, allowed)
}

func TestInjectTimeBoxedDelimiter(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	timeboxLimiter := NewTimeBoxedLimiter(refTime, limiter, time.Hour)

	dayLimiter := NewDayLimiterWithTimeBox(timeboxLimiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := dayLimiter.Allow(refTime, "client_1")

	// then
	assert.True(t, allowed)
}
