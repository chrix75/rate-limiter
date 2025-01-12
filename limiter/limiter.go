package limiter

type RateLimiter struct {
	limits map[string]int
}

func (l *RateLimiter) SetMaxCallsForClient(clientName string, max int) {
	if max < 0 {
		panic("max cannot be negative")
	}

	if clientName == "" {
		panic("clientName cannot be empty")
	}

	l.limits[clientName] = max
}

func (l *RateLimiter) Allow(clientName string) bool {
	return l.limits[clientName] > 0
}

func NewLimiter() *RateLimiter {
	return &RateLimiter{
		limits: make(map[string]int),
	}
}
