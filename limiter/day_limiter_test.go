package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAllowCallInCurrentDay(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewFixedValueTimer(refTime)

	limiter := NewCounterLimiter(repo)

	dayLimiter := NewDayLimiter(timer, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := dayLimiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallInCurrentDay(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewFixedValueTimer(refTime)

	limiter := NewCounterLimiter(repo)

	dayLimiter := NewDayLimiter(timer, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow("client_1")
	allowed := dayLimiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}

func TestResetBucketByDay(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewDynamicTimer(refTime)

	limiter := NewCounterLimiter(repo)

	dayLimiter := NewDayLimiter(timer, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow("client_1")

	nextCallTime := refTime.Add(24 * time.Hour)
	timer.T = nextCallTime
	allowed := dayLimiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallAfterResetByDay(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewDynamicTimer(refTime)

	limiter := NewCounterLimiter(repo)

	dayLimiter := NewDayLimiter(timer, limiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = dayLimiter.Allow("client_1")

	nextCallTime := refTime.Add(time.Minute)
	timer.T = nextCallTime

	_ = dayLimiter.Allow("client_1")
	allowed := dayLimiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}

func TestInjectTimeBoxedDelimiter(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewFixedValueTimer(refTime)

	limiter := NewCounterLimiter(repo)

	timeboxLimiter := NewTimeBoxedLimiter(timer, limiter, time.Hour)

	dayLimiter := NewDayLimiter(timer, timeboxLimiter)
	dayLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := dayLimiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}
