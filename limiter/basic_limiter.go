package limiter

// RateLimiter defines an interface for managing rate limiting logic.
// SetMaxCallsForClient sets the maximum number of allowed calls for a specific client.
// Allow checks if a specified client is allowed to proceed based on the configured rate limits.
type RateLimiter interface {
	SetMaxCallsForClient(clientName string, max int)
	Allow(clientName string) bool
}

type CounterLimiter struct {
	limits LimitRepository
}

// SetMaxCallsForClient sets the maximum number of allowed calls for a specific client.
// Panics if the provided max is negative or the clientName is empty.
func (l *CounterLimiter) SetMaxCallsForClient(clientName string, max int) {
	if max < 0 {
		panic("max cannot be negative")
	}

	if clientName == "" {
		panic("clientName cannot be empty")
	}

	l.limits.AddClient(clientName, max)
}

// Allow determines if a client is permitted to perform an operation based on their remaining allowed calls.
// It decrements the client’s remaining call count and returns true if the remaining count is greater than zero.
func (l *CounterLimiter) Allow(clientName string) bool {
	return l.limits.DecAndGet(clientName) > 0
}

func NewCounterLimiter(limitRepo LimitRepository) *CounterLimiter {
	return &CounterLimiter{
		limits: limitRepo,
	}
}
