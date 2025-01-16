package limiter

type LimitRepository interface {
	AddClient(clientName string, maxCalls int)
	DecAndGet(clientName string) int
}

type RateLimiter interface {
	SetMaxCallsForClient(clientName string, max int)
	Allow(clientName string) bool
}

type CounterLimiter struct {
	limits LimitRepository
}

func (l *CounterLimiter) SetMaxCallsForClient(clientName string, max int) {
	if max < 0 {
		panic("max cannot be negative")
	}

	if clientName == "" {
		panic("clientName cannot be empty")
	}

	l.limits.AddClient(clientName, max)
}

func (l *CounterLimiter) Allow(clientName string) bool {
	return l.limits.DecAndGet(clientName) > 0
}

func NewLimiter(limitRepo LimitRepository) *CounterLimiter {
	return &CounterLimiter{
		limits: limitRepo,
	}
}
