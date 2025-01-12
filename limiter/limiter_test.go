package limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllowCall(t *testing.T) {
	// given
	limiter := NewLimiter()
	limiter.SetMaxCallsForClient("client_1", 1)

	// when
	allowed := limiter.Allow("client_1")

	// then
	assert.True(t, allowed)
}

func TestRefuseCall(t *testing.T) {
	// given
	limiter := NewLimiter()
	limiter.SetMaxCallsForClient("client_1", 0)

	// when
	allowed := limiter.Allow("client_1")

	// then
	assert.False(t, allowed)
}
