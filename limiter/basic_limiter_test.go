package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var repo = &InMemoryLimitRepo{}

func TestAllowCall(t *testing.T) {
	// given
	limiter := NewCounterLimiter(repo)
	limiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := limiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCall(t *testing.T) {
	// given
	limiter := NewCounterLimiter(repo)
	limiter.SetMaxCallsForClient("client_1", 0)

	// when
	allowed := limiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}

func TestDontAllowCallAnymore(t *testing.T) {
	// given
	limiter := NewCounterLimiter(repo)
	limiter.SetMaxCallsForClient("client_1", 1)
	_ = limiter.Allow("client_1")

	// when
	allowed := limiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}
