package limiter

import "time"

type TimeBoxedLimiter struct {
	refTime time.Time
	limiter *RateLimiter
	period  time.Duration
	limits  map[string]int
}

func (l *TimeBoxedLimiter) SetMaxCallsForClient(clientName string, max int) {
	l.limiter.SetMaxCallsForClient(clientName, max)
	l.limits[clientName] = max
}

func (l *TimeBoxedLimiter) Allow(t time.Time, clientName string) bool {
	elapseTime := t.Sub(l.refTime)
	if elapseTime >= l.period {
		limit := l.limits[clientName]
		l.limiter.SetMaxCallsForClient(clientName, limit)
	}
	l.refTime = t
	return l.limiter.Allow(clientName)
}

func NewTimeBoxedLimiter(refTime time.Time, limiter *RateLimiter, period time.Duration) *TimeBoxedLimiter {
	return &TimeBoxedLimiter{
		refTime: refTime,
		limiter: limiter,
		period:  period,
		limits:  make(map[string]int),
	}
}
