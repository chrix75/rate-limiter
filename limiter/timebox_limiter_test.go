package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAllowCallInTimeBox(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewConstantTimer(refTime)
	limiter := NewLimiter(fakeRepo)

	timeBoxedLimiter := NewTimeBoxedLimiter(timer, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := timeBoxedLimiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallInTimeBox(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewConstantTimer(refTime)
	limiter := NewLimiter(fakeRepo)

	timeBoxedLimiter := NewTimeBoxedLimiter(timer, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow("client_1")
	allowed := timeBoxedLimiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}

func TestResetBucketByTimeBox(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewDynamicTimer(refTime)

	limiter := NewLimiter(fakeRepo)

	timeBoxedLimiter := NewTimeBoxedLimiter(timer, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow("client_1")

	nextCallTime := refTime.Add(time.Minute)
	timer.T = nextCallTime
	allowed := timeBoxedLimiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCallAfterResetByTimeBox(t *testing.T) {
	// given
	refTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	timer := NewDynamicTimer(refTime)

	limiter := NewLimiter(fakeRepo)

	timeBoxedLimiter := NewTimeBoxedLimiter(timer, limiter, time.Minute)
	timeBoxedLimiter.SetMaxCallsForClient("client_1", 1)

	// when
	_ = timeBoxedLimiter.Allow("client_1")

	nextCallTime := refTime.Add(time.Minute)
	timer.T = nextCallTime
	_ = timeBoxedLimiter.Allow("client_1")
	allowed := timeBoxedLimiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}
