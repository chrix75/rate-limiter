package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAllowCallInTimeBox(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timeBoxedLimiter := NewTimeBoxedLimiter(refTime, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := timeBoxedLimiter.Allow(refTime, "client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallInTimeBox(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timeBoxedLimiter := NewTimeBoxedLimiter(refTime, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow(refTime, "client_1")
	allowed := timeBoxedLimiter.Allow(refTime, "client_1")

	// then
	assert.False(t, allowed)
}

func TestResetBucketByTimeBox(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timeBoxedLimiter := NewTimeBoxedLimiter(refTime, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow(refTime, "client_1")

	nextCallTime := refTime.Add(time.Minute)
	allowed := timeBoxedLimiter.Allow(nextCallTime, "client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallAfterResetByTimeBox(t *testing.T) {
	// given
	limiter := NewLimiter(fakeRepo)

	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timeBoxedLimiter := NewTimeBoxedLimiter(refTime, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow(refTime, "client_1")

	nextCallTime := refTime.Add(time.Minute)
	_ = timeBoxedLimiter.Allow(nextCallTime, "client_1")
	allowed := timeBoxedLimiter.Allow(nextCallTime, "client_1")

	// then
	assert.False(t, allowed)
}
