package limiter

type LimitRepository interface {
	AddClient(clientName string, maxCalls int)
	DecAndGet(clientName string) int
}

type RateLimiter struct {
	limits LimitRepository
}

func (l *RateLimiter) SetMaxCallsForClient(clientName string, max int) {
	if max < 0 {
		panic("max cannot be negative")
	}

	if clientName == "" {
		panic("clientName cannot be empty")
	}

	l.limits.AddClient(clientName, max)
}

func (l *RateLimiter) Allow(clientName string) bool {
	return l.limits.DecAndGet(clientName) > 0
}

func NewLimiter(limitRepo LimitRepository) *RateLimiter {
	return &RateLimiter{
		limits: limitRepo,
	}
}
